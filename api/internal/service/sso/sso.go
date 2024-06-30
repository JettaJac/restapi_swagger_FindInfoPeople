package sso

import (
	// "errors"
	"log/slog"
)

// var (
// 	ErrUserExists         = errors.New("user does not exists")
// 	ErrInvalidCredentials = errors.New("invalid login or password")
// 	ErrConnectionTime     = errors.New("cannot connect to database")
// )

type SSO struct {
	Log *slog.Logger
	// UserSaver    UserSaver
	// UserProvider UserProvider
	// TokenTTl     time.Duration
}

func New(log *slog.Logger /*, userSaver UserSaver, userProvider UserProvider, tokenTTL time.Duration*/) *SSO {
	return &SSO{
		Log: log,
		// UserSaver:    userSaver,
		// UserProvider: userProvider,
		// TokenTTl:     tokenTTL,
	}
}
