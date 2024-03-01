package api

import (
	"backend/ent"
	"backend/ent/section"
	"backend/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) getSection(ctx *gin.Context) {
	sectionName := ctx.Query(requestParamSectionName)
	if !utils.SectionNameIsValid(sectionName) {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.InvalidRequestParamErr(requestParamSectionName)))
		return
	}

	section, err := server.client.Section.
		Query().
		Where(
			section.Or(section.DeletedAtIsNil(), section.DeletedAtGT(time.Now())),
			section.SectionName(sectionName),
		).
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.InvalidRequestParamErr(requestParamSectionName)))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sectionName": section.SectionName,
		"statement":   *section.Statement,
		"managerID":   section.ManagerID,
	})
}

func (server *Server) suggestSection(ctx *gin.Context) {
	sectionName := ctx.Query(requestParamSectionName)
	if !utils.SectionNameIsValid(sectionName) {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.InvalidRequestParamErr(sectionName)))
		return
	}

	limit, err := strconv.Atoi(ctx.Query(requestParamLimit))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	if !utils.LimitIsValid(limit) {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.InvalidRequestParamErr(requestParamLimit)))
		return
	}

	offset, err := strconv.Atoi(ctx.Query(requestParamOffset))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	if !utils.OffsetIsValid(offset) {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.InvalidRequestParamErr(requestParamOffset)))
		return
	}

	var suggestSection []struct {
		ID          int    `json:"id"`
		SectionName string `json:"section_name"`
	}
	err = server.client.Section.
		Query().
		Where(
			section.Or(
				section.SectionNameHasPrefix(sectionName),
				section.SectionNameContains(sectionName),
			),
			section.Or(
				section.DeletedAtGT(time.Now()),
				section.DeletedAtIsNil(),
			),
		).
		Limit(limit).
		Offset(offset).
		Order(
			ent.Desc(section.FieldUpdatedAt),
		).
		Select(
			section.FieldID,
			section.FieldSectionName,
		).
		Scan(context.Background(), &suggestSection)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, suggestSection)
}
