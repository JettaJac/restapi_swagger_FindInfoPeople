package people

import (
	// "errors"
	"context"
	"log/slog"
	"main/internal/domain/models"
)

// var (
// 	ErrUserExists         = errors.New("user does not exists")
// 	ErrInvalidCredentials = errors.New("invalid login or password")
// 	ErrConnectionTime     = errors.New("cannot connect to database")
// )

type People struct {
	Log      *slog.Logger
	Provider PeopleProvider
	// UserSaver    UserSaver
	// UserProvider UserProvider
	// TokenTTl     time.Duration
}

// type PeopleSaver interface {
// 	SavePost(ctx context.Context, userID int64, title, content string, allowComments bool) (int64, time.Time, error)
// }

// PeopleProvider - интерфейс для получения информации о людях
type PeopleProvider interface {
	GetInfo(ctx context.Context, PassportSerie, PassportNumber int64) (models.User, error)
	// ProvideAllPosts(ctx context.Context, page int64) ([]models.Post, error)
}

func New(log *slog.Logger, peopleProvider PeopleProvider /*, userSaver UserSaver, userProvider UserProvider, tokenTTL time.Duration*/) *People {
	return &People{
		Log:      log,
		Provider: peopleProvider,
		// UserSaver:    userSaver,
		// UserProvider: userProvider,
		// TokenTTl:     tokenTTL,
	}
}

// func (p *People) GetInfo(ctx context.Context, PassportSerie, PassportNumber int64) (models.User, error) {
// 	return p.Provider.GetInfo(ctx, PassportSerie, PassportNumber)
// }
