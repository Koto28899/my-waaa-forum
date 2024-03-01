package api

import (
	"backend/ent"
	"backend/ent/role"
	"backend/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type registerUserRequestParams struct {
	Email string `json:"email" binding:"required,validEmail"`
	Pwd   string `json:"pwd" binding:"required,validPwd"`
	Name  string `json:"name" binding:"required,validRoleName"`
}

func (server *Server) registerUser(ctx *gin.Context) {
	var arg registerUserRequestParams
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	hashedPwd, err := utils.HashedPwd(arg.Pwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tx, err := server.client.Tx(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	role, err := tx.Role.
		Create().
		SetEmail(arg.Email).
		SetHashedPwd(hashedPwd).
		SetRoleName(arg.Name).
		SetType(role.TypeUser).
		Save(context.Background())
	if err != nil {
		if ent.IsConstraintError(err) {
			tx.Rollback()
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(fmt.Errorf("email is already existing")))
			return
		}
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	session, err := tx.Session.
		Create().
		SetRoleID(role.ID).
		SetExpiresAt(time.Now().Add(server.config.SessionDuration)).
		SetUserAgent(ctx.Request.UserAgent()).
		SetClientIP(ctx.ClientIP()).
		SetIsBlocked(false).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
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
