package api

import (
	"backend/middleware"
	"backend/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createPostRequestParameters struct {
	SectionID int64  `json:"sectionID" binding:"required;"`
	Title     string `json:"title" binding:"required; validPostTitle"`
	Body      string `json:"body" binding:"required; validPostBody"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var arg createPostRequestParameters
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	roleID, err := strconv.ParseInt(ctx.Param(middleware.AuthorizedRoleIDKey), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tx, err := server.client.Tx(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	post, err := tx.Post.
		Create().
		SetFromRoleID(roleID).
		SetToSectionID(arg.SectionID).
		SetTitle(arg.Title).
		SetBody(arg.Body).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = tx.Role.UpdateOneID(roleID).
		AddPostsCount(1).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}
