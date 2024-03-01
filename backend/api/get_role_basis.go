package api

import (
	"backend/ent"
	"backend/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) getRoleBasis(ctx *gin.Context) {
	roleID, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	role, err := server.client.Role.Get(context.Background(), roleID)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(ErrRoleDoesNotExist))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	statement := "这个人很懒，什么也没留下。"
	if role.Statement != nil {
		statement = *role.Statement
	}
	roleType := "普通成员"
	if role.Type == "admin" {
		roleType = "管理员"
	}

	createdAt := role.CreatedAt.Format("2006-01-02 15:04:05")

	ctx.JSON(http.StatusOK, gin.H{
		"id":            role.ID,
		"createdAt":     createdAt,
		"roleName":      role.RoleName,
		"type":          roleType,
		"statement":     statement,
		"postsCount":    role.PostsCount,
		"commentsCount": role.CommentsCount,
		"repliesCount":  role.RepliesCount,
		"likesCount":    role.LikesCount,
		"helpsCount":    role.HelpsCount,
		"fansCount":     role.FansCount,
	})
}
