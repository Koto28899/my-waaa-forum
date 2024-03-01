package middleware

import (
	"backend/ent"
	"backend/ent/role"
	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminAuth(config utils.Config, client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prevValue, exists := ctx.Get(AuthorizedRoleIDKey)
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
			return
		}

		roleID, ok := prevValue.(int64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
			return
		}

		_, err := client.Role.Query().
			Where(
				role.IDEQ(roleID),
				role.Or(
					role.DeletedAtGT(time.Now()),
					role.DeletedAtIsNil(),
				),
				role.TypeEQ(role.TypeAdmin),
			).
			Only(context.Background())
		if err != nil {
			ctx.Set(AuthorizedRoleIDKey, 0)
			if ent.IsNotFound(err) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(ErrUnauthorizedAction))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}

		ctx.Next()
	}
}
