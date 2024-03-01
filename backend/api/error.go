package api

import "errors"

var (
	ErrUserLoginFailed      = errors.New("用户不存在或密码错误")
	ErrInvalidRoleIDGiven   = errors.New("提供的管理员 ID 不存在或不是管理员")
	ErrDuplicateSectionName = errors.New("版块名已被使用")
	ErrRoleDoesNotExist     = errors.New("用户不存在")
)
