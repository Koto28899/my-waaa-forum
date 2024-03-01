package api

import (
	"backend/ent"
	"backend/ent/role"
	"backend/ent/session"
	"backend/middleware"
	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleBasisParams struct {
	RoleID   int64  `json:"roleID"`
	RoleType string `json:"roleType"`
}

func (server *Server) parseCookie(ctx *gin.Context) {
	// userAgent := ctx.Request.UserAgent()
	clientIP := ctx.ClientIP()
	cookie, err := ctx.Cookie(server.config.CookieName)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	session, err := server.client.Session.
		Query().
		Where(session.ID(uuid.MustParse(cookie)),
			session.ClientIP(clientIP),
			// session.UserAgent(userAgent),
			session.ExpiresAtGT(time.Now()),
			session.IsBlocked(false),
		).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusUnauthorized, middleware.ErrUnauthorizedAction)
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	_, err = server.client.Role.Query().
		Where(
			role.ID(session.RoleID),
			role.Or(role.DeletedAtIsNil(), role.DeletedAtGT(time.Now())),
		).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusUnauthorized, middleware.ErrUnauthorizedAction)
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	if time.Now().Add(30 * 24 * time.Hour).After(session.ExpiresAt) {
		_, err := server.client.Session.
			UpdateOneID(session.ID).
			SetExpiresAt(time.Now().Add(server.config.SessionDuration)).
			Save(context.Background())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
	}
	role, err := server.client.Role.Query().
		Where(role.DeletedAtIsNil(), role.ID(session.RoleID)).
		Only(context.Background())
	// _, err = server.client.Role.Query().
	// 	Where(role.DeletedAtIsNil(), role.ID(session.RoleID)).
	// 	Only(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, RoleBasisParams{
		RoleID:   role.ID,
		RoleType: role.Type.String(),
	})
}
