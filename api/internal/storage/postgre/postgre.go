package postgre

import (
	// "context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	swapi "main/generated"
	"main/internal/domain/models"
	"main/internal/storage"
)

type Storage struct {
	db *sql.DB
}

type Postgres struct {
	Host     string `env:"host" env-default:"localhost__"`
	Port     int    `env:"port" `
	User     string `env:"user" env-default:"postgres_5"`
	Password string `env:"password"`
	DBName   string `env:"db_name"`
	SSLMode  string `env:"ssl_mode"`
}

func New(p Postgres) (*Storage, error) {
	const op = "storage.postgre.New"
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		p.Host, p.Port, p.User, p.DBName, p.Password, p.SSLMode) // !!!Переделать
	// conn := fmt.Sprintf("host=localhost port=%d user=postgres dbname=postgres password=postgres sslmode=disable", 5432)

	db, err := sql.Open("postgres" /*storagePath*/, conn)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s.Ping: %w", op, err)
	}

	// возможно здесь запустить миграции как в п.посгре

	return &Storage{db: db}, nil
}

// CloseDB close database
// func (storage *Storage) CloseDB() { //!!!может и не надо
// 	storage.db.Close()
// }

// !!! сделать запросы горутиной, возмжно, как у Кати
func (s *Storage) GetInfo( /*ctx context.Context, */ p swapi.GetInfoParams) (*models.User, error) {
	const op = "storage.posgre.User"

	var user = &models.User{}

	query := fmt.Sprintf("SELECT id, surname, name, patronymic, address FROM users WHERE passportSerie = $1 AND passportNumber = $2")
	err := s.db.QueryRow(query, p.PassportSerie, p.PassportNumber).Scan(&user.ID, &user.Surname, &user.Name, &user.Patronymic, &user.Address)
	fmt.Println(user, err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("N:%s: %w", op, storage.ErrUserNotFound)
		}
		return nil, fmt.Errorf("O:%s:  %w", op, err)
	}

	return user, nil
}
