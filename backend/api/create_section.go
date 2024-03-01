package api

import (
	"backend/ent"
	"backend/ent/role"
	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createSectionRequiredParameters struct {
	SectionName string `json:"sectionName" binding:"required,validSectionName"`
	Statement   string `json:"statement" binding:"validStatement"`
	ManagerID   int64  `json:"managerID" binding:"required,min=1"`
}

func (server *Server) createSection(ctx *gin.Context) {
	var arg createSectionRequiredParameters
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	tx, err := server.client.Tx(context.Background())
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	_, err = tx.Role.Query().
		Where(
			role.Or(
				role.DeletedAtGT(time.Now()),
				role.DeletedAtIsNil(),
			),
			role.TypeEQ(role.TypeAdmin),
			role.ID(arg.ManagerID),
		).
		ForShare().
		Only(context.Background())
	if err != nil {
		tx.Rollback()
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(ErrInvalidRoleIDGiven))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	section, err := server.client.Section.Create().
		SetSectionName(arg.SectionName).
		SetStatement(arg.Statement).
		SetManagerID(arg.ManagerID).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		if ent.IsConstraintError(err) {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(ErrDuplicateSectionName))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	if tx.Commit() != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, section)
}
