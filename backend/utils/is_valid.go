package utils

import "strings"

func EmailIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	if !strings.Contains(trimed, "@") {
		return false
	}
	tmp := len(trimed)
	return 1 < tmp && tmp < 321+1
}

func PwdIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return 0 < tmp && tmp < 255+1
}

func RoleNameIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return 0 < tmp && tmp < 21+1
}

func SectionNameIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return 0 < tmp && tmp < 21+1
}

func PostTitleIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return 6 < tmp && tmp < 127+1
}

func PostBodyIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return 0 < tmp && tmp < 1023+1
}

func StatementIsValid(item string) bool {
	trimed := strings.TrimSpace(item)
	tmp := len(trimed)
	return tmp < 127+1
}

func LimitIsValid(item int) bool {
	return 0 < item && item < 20
}

func OffsetIsValid(item int) bool {
	return 0 <= item && item < 20
}
