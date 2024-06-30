package postgre

import (
	"context"
	"database/sql"
	// "errors"
	"fmt"
	// "github.com/lib/pq"
	// "sso/internal/domain/models"
	// "sso/internal/storage"
)

type Storage struct {
	db *sql.DB
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"postgres"`
	Password string `yaml:"postgres"`
	DBName   string `yaml:"db"`
	SSLMode  string `yaml:"ssl_mode"`
}

func New(storagePath string /*p Postgres*/) (*Storage, error) {
	const op = "storage.postgre.New"
	// conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
	// 	p.Host, p.Port, p.User, p.DBName, p.Password, p.SSLMode) // !!!Переделать по своему

	db, err := sql.Open("postgres", storagePath)
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
func (storage *Storage) CloseDB() { //!!!может и не надо
	storage.db.Close()
}

// GetInfo(w http.ResponseWriter, r *http.Request, params GetInfoParams)
func (s *Storage) GetInfo(ctx context.Context, passportSerie int, passportNumber int) (surname, name, patronymic, address string) {

}
