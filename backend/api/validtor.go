package api

import (
	"backend/utils"

	"github.com/go-playground/validator/v10"
)

var validEmail validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.EmailIsValid(item)
	}
	return false
}

var validRoleName validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.RoleNameIsValid(item)
	}
	return false
}

var validPostTitle validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.PostTitleIsValid(item)
	}
	return false
}

var validPostBody validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.PostBodyIsValid(item)
	}
	return false
}

var validPwd validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.PwdIsValid(item)
	}
	return false
}

var validSectionName validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.SectionNameIsValid(item)
	}
	return false
}

var validStatment validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if item, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.StatementIsValid(item)
	}
	return false
}
