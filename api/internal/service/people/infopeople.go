package people

//переименовать возможно с добавлением сервис

// !!!

import (
	// "errors"
	// "context"
	"errors"
	"fmt"
	"log/slog"
	swapi "main/generated"
	"main/internal/domain/models"
	"main/internal/storage"
	"main/pkg/lib/logger"
)

// var (
// 	ErrUserExists         = errors.New("user does not exists")
// 	ErrInvalidCredentials = errors.New("invalid login or password")
// 	ErrConnectionTime     = errors.New("cannot connect to database")
// )

type Info struct {
	log     *slog.Logger
	storage PeopleProvider //!!! Возможно изменить на provider
	// UserSaver    UserSaver
	// UserProvider UserProvider
	// TokenTTl     time.Duration
}

// PeopleProvider - интерфейс для получения информации о людях
type PeopleProvider interface {
	GetInfo( /*ctx context.Context,*/ p swapi.GetInfoParams /* PassportSerie, PassportNumber int64*/) (*models.User, error)
	// ProvideAllPosts(ctx context.Context, page int64) ([]models.Post, error)
	GetList( /*ctx context.Context, */ p swapi.GetListParams) ([]models.User, error)
}

var ( //Откорректировать
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppID       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

func New(log *slog.Logger, storage PeopleProvider /*, userSaver UserSaver, userProvider UserProvider, tokenTTL time.Duration*/) *Info {
	return &Info{
		log:     log,
		storage: storage,
		// UserSaver:    userSaver,
		// UserProvider: userProvider,
		// TokenTTl:     tokenTTL,
	}
}

func (i *Info) GetInfo( /*ctx context.Context, */ p swapi.GetInfoParams) (*models.User, error) {
	var user *models.User

	const op = "info.GetInfo"
	fmt.Println(p.PassportSerie, "IIIIIIIIIIIII", p.PassportNumber)
	user, err := i.storage.GetInfo( /*ctx,*/ p)

	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			i.log.Warn("user not found", sl.Err(err))

			return nil, fmt.Errorf("%s:  %v", op, ErrInvalidCredentials)
		}

		i.log.Error("failed to get user", sl.Err(err))
		return nil, fmt.Errorf("%s:  %v", op, err)
	}

	//проверка если надо вводных данных

	return user, nil
}

func (i *Info) GetList( /*ctx context.Context, */ p swapi.GetListParams) ([]models.User, error) {
	// var user *models.User
	// users := make([]models.User, 0)
	fmt.Println("___________listPeople 93 ______", p)
	const op = "info.GetInfo"
	fmt.Println(*p.PassportSerie, "Info_IIIIIIIIIIIII", *p.Name, *p.PassportNumber)
	users, err := i.storage.GetList( /*ctx,*/ p)
	fmt.Println("___________listPeople 94 ______", p)
	if err != nil {
		i.log.Error("failed to get user", sl.Err(err))
		return nil, fmt.Errorf("%s:  %v", op, err)
	}

	// //проверка если надо вводных данных, но скорее всего на уровне хендлера

	return users, nil
}
