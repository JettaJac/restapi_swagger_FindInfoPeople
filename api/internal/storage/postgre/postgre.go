package postgre

import (
	// "context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	swapi "main/generated"
	"main/internal/domain/models"
	"main/internal/storage"
	"strings"
)

type Storage struct {
	db *sql.DB
}

type Postgres struct {
	Host     string `env:"host" env-default:"localhost"`
	Port     int    `env:"port" `
	User     string `env:"user" env-default:"postgres"`
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
	// err = migrations(p.DBName, ".")
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }

	return &Storage{db: db}, nil
}

// migrations create migrations
func migrations(host, path string) error {

	const op = "Migrations"
	qualy := fmt.Sprintf("file://%s/migrations", path)

	m, err := migrate.New(qualy, host)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = m.Up()

	if errors.Is(err, errors.New("no change")) && err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// CloseDB close database
// func (storage *Storage) CloseDB() { //!!!может и не надо
// 	storage.db.Close()
// }

// !!! сделать запросы горутиной, возмжно, как у Кати
func (s *Storage) GetInfo( /*ctx context.Context, */ p swapi.GetInfoParams) (*models.User, error) { //User заменить на Person
	const op = "storage.posgre.User"

	var user = &models.User{}
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("B:%s:  %w", op, err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT id, surname, name, patronymic, address FROM users WHERE passportSerie = $1 AND passportNumber = $2")

	err = tx.QueryRow(query, p.PassportSerie, p.PassportNumber).Scan(&user.ID, &user.Surname, &user.Name, &user.Patronymic, &user.Address)
	// fmt.Println(user, err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("N:%s: %w", op, storage.ErrUserNotFound)
		}
		return nil, fmt.Errorf("O:%s:  %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("С:%s:  %w", op, err)
	}

	return user, nil
}

// func appendarg(queryArgs []interface{}, param interface{}, *query string ){
// if param != nil {
//         queryArgs = append(queryArgs, param)
//         query += "$1, "
//     } else {
//         query += "NULL, "
//     }
// }

func (s *Storage) GetList( /*ctx context.Context, */ p swapi.GetListParams) ([]models.User, error) {
	const op = "storage.posgre.User"
	fmt.Println("___________listPeople 96 ______", *p.Name, *p.PassportSerie)
	fmt.Println(*p.PassportSerie, "IIIIIIIIIIIII_1", *p.Name, *p.PassportNumber, "__", *p.Page, *p.Limit)
	fields, values := actualRequest(p)
	fmt.Println(fields)
	fmt.Println(values)
	query := fmt.Sprintf("SELECT * FROM persons WHERE ")
	args := make([]interface{}, len(values))
	i := 0
	for _, field := range fields {
		fmt.Println(field)
		query += fmt.Sprintf("%v=? AND ", field)
		args[i] = values[i]
		i++
	}

	// query += " " + strings.Join(fields, " AND ") + " LIMIT ? OFFSET ?"
	// rows, err := db.Query(query, name, passportSerie, age, id, limit, (page-1)
	// Удаляем последний "AND "
	query = query[:len(query)-5]
	fmt.Println(query)

	rows, err := s.db.Query(query, args...)
	fmt.Println("{{{{{{{{{{{{{{{{{{{{", query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.User
	for rows.Next() {
		var p models.User
		err := rows.Scan(&p.Surname, &p.Name, &p.Patronymic, &p.Address, &p.PassportSerie, &p.PassportNumber)
		if err != nil {
			// log.Fatal(err)
			return nil, err
		}
		persons = append(persons, p)
	}

	// query := "SELECT id, surname, name, patronymic, address, passport_serie, passport_number FROM filter_users($1, $2, $3, $4, $5, $6, $7) LIMIT $8 OFFSET $9"
	// var queryArgs []interface{}

	// // if p.Id != nil {
	// // 	queryArgs = append(queryArgs, p.Id)
	// // 	query += "$1, "
	// // } else {
	// // 	query += "NULL, "
	// // }
	// if p.Surname != nil {
	// 	queryArgs = append(queryArgs, p.Surname)

	// } else {
	// 	queryArgs = append(queryArgs, "")
	// }
	// if p.Name != nil {
	// 	queryArgs = append(queryArgs, p.Name)

	// } else {
	// 	queryArgs = append(queryArgs, "")
	// }

	// if p.Patronymic != nil {
	// 	queryArgs = append(queryArgs, p.Patronymic)

	// } else {
	// 	queryArgs = append(queryArgs, "")
	// }
	// if p.Address != nil {
	// 	queryArgs = append(queryArgs, p.Address)
	// } else {
	// 	queryArgs = append(queryArgs, "")
	// }

	// if p.PassportSerie != nil {
	// 	queryArgs = append(queryArgs, p.PassportSerie)

	// } else {
	// 	queryArgs = append(queryArgs, 0)
	// }

	// if p.PassportNumber != nil {
	// 	queryArgs = append(queryArgs, p.PassportNumber)
	// } else {
	// 	queryArgs = append(queryArgs, "")
	// }

	// query += "?, ?, ?, ?, ?, ?, ?) LIMIT ? OFFSET ?"
	// queryArgs = append(queryArgs, *p.PassportSerie, *p.PassportNumber, *p.Limit, (*p.Page-1)**p.Limit)

	// rows, err := s.db.Query(query, queryArgs...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// fmt.Println("___________listPeople 97 ______", p)
	// // for rows.Next() {
	// // 	var user models.User
	// // 	if err := rows.Scan(&user.ID, &user.Surname, &user.Name, &user.Patronymic, &user.Address /*, &user.PassportSerie, &user.PassportNumber*/); err != nil {
	// // 		return nil, err
	// // 	}
	// // 	users = append(users, user)
	// // }

	// // if err := rows.Err(); err != nil {
	// // 	return nil, err
	// // }

	return persons, nil

}
func actualRequest(p swapi.GetListParams) ([]interface{}, []string) {
	var fields []interface{}
	var values []string
	fmt.Println(*p.PassportSerie, "IIIIIIIIIIIII", *p.Name, *p.PassportNumber)
	if p.Surname != nil {
		fields = append(fields, *p.Surname)
		values = append(values, "surname")
	}
	if p.Name != nil {
		fields = append(fields, *p.Name)
		values = append(values, "name")
	}
	if p.Patronymic != nil {
		fields = append(fields, *p.Patronymic)
		values = append(values, "patronic")
	}
	if p.Address != nil {
		fields = append(fields, *p.Address)
		values = append(values, "address")
	}
	if *p.PassportSerie != 0 {
		fields = append(fields, *p.PassportSerie)
		values = append(values, "passportserie")
	}
	if *p.PassportNumber != 0 {
		fields = append(fields, *p.PassportNumber)
		values = append(values, "passportnumber")
	}
	if *p.Limit != 0 {
		fields = append(fields, *p.Limit)
		values = append(values, "limit")
	}
	if *p.Page != 0 {
		fields = append(fields, *p.Page)
		values = append(values, "page")
	}

	return fields, values
}

// func FilterPersons2(fields []interface{}) ([]*models.User, error) {
// 	query := fmt.Sprintf("SELECT * FROM persons WHERE ")
// 	args := make([]interface{}, len(fields))

// 	i := 0
// 	for _, field := range fields {
// 		query += fmt.Sprintf("%s=? AND ", field)
// 		args[i] = values[i]
// 		i++
// 	}

// 	var persons []*models.User
// 	return persons, nil
// }

func FilterPersons(fields []string, values []interface{}) ([]*models.User, error) {
	// query := fmt.Sprintf("SELECT * FROM persons WHERE ")
	// args := make([]interface{}, len(values))
	// i := 0
	// for _, field := range fields {
	// 	query += fmt.Sprintf("%s=? AND ", field)
	// 	args[i] = values[i]
	// 	i++
	// }
	// // Удаляем последний "AND "
	// query = query[:len(query)-5]

	// rows, err := s.db.Query(query, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	var persons []*models.User
	// for rows.Next() {
	// 	var p models.User
	// 	err := rows.Scan(&p.Surname, &p.Name, &p.Patronymic, &p.Address, &p.PassportSerie, &p.PassportNumber)
	// 	if err != nil {
	// 		// log.Fatal(err)
	// 		return nil, err
	// 	}
	// 	persons = append(persons, &p)
	// }

	return persons, nil
}
