package api

import (
	"backend/ent"
	"backend/ent/post"
	"backend/utils"
	"context"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
)

func (server *Server) getRolePosts(ctx *gin.Context) {
	roleID, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	posts, err := server.client.Post.Query().
		Where(
			post.FromRoleID(roleID),
			post.DeletedAtNotNil(),
		).
		Offset(offset).
		Limit(limit).
		Order(
			post.ByCreatedAt(
				sql.OrderDesc(),
			),
		).
		All(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.Status(http.StatusOK)
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, posts)
}
