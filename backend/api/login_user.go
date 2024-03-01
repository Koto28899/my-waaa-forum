package api

import (
	"backend/ent"
	"backend/ent/role"
	"backend/ent/session"
	"log"

	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loginUserRequestParams struct {
	Email string `json:"email" binding:"required,validEmail"`
	Pwd   string `json:"pwd" binding:"required,validPwd"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	clientIP, userAgent := ctx.ClientIP(), ctx.Request.UserAgent()

	var arg loginUserRequestParams
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	role, err := server.client.Role.
		Query().
		Where(role.Email(arg.Email)).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUserLoginFailed))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	if utils.CheckPwd(arg.Pwd, role.HashedPwd) != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUserLoginFailed))
		return
	}

	session, err := server.client.Session.
		Query().
		Where(session.ClientIP(clientIP), session.UserAgent(userAgent), session.RoleID(role.ID)).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			session, err = server.client.Session.
				Create().
				SetRoleID(role.ID).
				SetExpiresAt(time.Now().Add(server.config.SessionDuration)).
				SetUserAgent(userAgent).
				SetClientIP(clientIP).
				SetIsBlocked(false).
				Save(context.Background())
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
				log.Fatal("1:", err.Error())
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			log.Fatal("2:", err.Error())
			return
		}
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		server.config.CookieName,
		session.ID.String(),
		int(server.config.SessionDuration),
		server.config.CookiePath,
		server.config.CookieDomain,
		true,
		true)

	ctx.JSON(http.StatusOK, RoleBasisParams{
		RoleID:   role.ID,
		RoleType: string(role.Type),
	})
}
