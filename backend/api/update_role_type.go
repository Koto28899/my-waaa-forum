package api

import (
	"backend/ent"
	"backend/ent/role"
	"backend/middleware"
	"backend/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateRoleTypeRequestParameters struct {
	DesRoleID int64  `json:"desRoleID" binding:"required;min=1"`
	Type      string `json:"type" binding:"required;oneof= user admin"`
}

func (server *Server) updateRoleType(ctx *gin.Context) {
	var arg updateRoleTypeRequestParameters
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	roleID, err := strconv.ParseInt(ctx.Param(middleware.AuthorizedRoleIDKey), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	role, err := server.client.Role.UpdateOneID(roleID).
		SetType(role.TypeAdmin).
		Save(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, role.Type)
}
