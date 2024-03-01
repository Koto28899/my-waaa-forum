package middleware

import "errors"

var (
	ErrCookieMissed                  = errors.New("cookies missed")
	ErrUnauthorizedAction            = errors.New("unauthorized action")
	ErrNoSuchSession                 = errors.New("there is not such a session")
	ErrNewUserAgentNeedAuthorization = errors.New("new user agent needs authorization")
	ErrNewClientIPNeedAuthorization  = errors.New("new client ip needs authorization")
	ErrBlocked                       = errors.New("session blocked")
)
