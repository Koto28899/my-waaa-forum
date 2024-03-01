package middleware

import (
	"backend/ent"
	"backend/ent/role"
	"backend/ent/session"
	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SessionAuth(config utils.Config, client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		cookie, err := ctx.Cookie(config.CookieName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
			return
		}
		session, err := client.Session.
			Query().
			Where(session.ID(uuid.MustParse(cookie)),
				session.ClientIP(clientIP),
				session.ExpiresAtGT(time.Now()),
				session.IsBlocked(false),
			).
			Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		_, err = client.Role.Query().
			Where(
				role.ID(session.RoleID),
				role.Or(role.DeletedAtIsNil(), role.DeletedAtGT(time.Now())),
			).
			Only(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		if time.Now().Add(30 * 24 * time.Hour).After(session.ExpiresAt) {
			_, err := client.Session.
				UpdateOneID(session.ID).
				SetExpiresAt(time.Now().Add(config.SessionDuration)).
				Save(context.Background())
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
				return
			}
		}

		ctx.Set(AuthorizedRoleIDKey, session.RoleID)
		ctx.Next()
	}
}
