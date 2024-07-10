package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	swapi "main/generated"
	"main/internal/config"
	"main/internal/domain/models"
	"main/internal/service/people"
	"net/http"
	"strconv"
)

var (
// GetInfo  = "/info"   // GetInfo
)

type Storage interface {
	GetInfo(ctx context.Context, p swapi.GetInfoParams) (*models.User, error) // ctx -!!! точно надо
	GetList(ctx context.Context, p swapi.GetListParams) ([]models.User, error)
}

// Server struct
type Server struct {
	Server *http.Server
	// Router *http.ServeMux
	// Addr   string
	// // authService  server.Auth
	// ReadTimeout  time.Duration
	// WriteTimeout time.Duration
	// IdleTimeout  time.Duration
	Router *http.ServeMux
	log    *slog.Logger
	info   people.Info
}

var _ swapi.ServerInterface = (*Server)(nil)

// type serverAPI struct {
// 	// ssov1.UnimplementedAuthServer
// 	// auth Auth
// }

// NewServer create a new server
func NewServer(config *config.Config, log *slog.Logger, info people.Info) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         config.HTTPServer.Address,
			ReadTimeout:  config.HTTPServer.Timeout,
			WriteTimeout: config.HTTPServer.Timeout,
			IdleTimeout:  config.HTTPServer.IdleTimeout,
			// Handler:      http.NewServeMux(),
		},
		// Router: http.NewServeMux(),
		log:  log,
		info: info,
		//!!! Нужно ли через сервер пробрасывать лог
	}
	// s.ConfigureRouter2()
	return s
}
func checkInt(param string) (num int, err error) {
	op := fmt.Sprintf("handler.GetInfo.paramsInfo.%s", param)
	if param != "" {
		num, err = strconv.Atoi(param)
		if err != nil {
			// s.log.Error("incorrect ID entered", slog.String("id: ", param))

			return num, fmt.Errorf("%s: %s", op, err)
		}
		// } else { // !!! Наверное нужно, так как не должен быть параметр пустым
		// 	s.log.Error("incorrect ID entered", slog.String("id: ", param))
		// 	// s.error(w, http.StatusBadRequest, storage.ErrEmptyRequest)
		// 	return params, fmt.Errorf("%s: %s", op, err)
	}
	fmt.Println("___________paramsList33______", num)
	return num, nil

}

func paramsInfo( /*w http.ResponseWriter,*/ r *http.Request) (p swapi.GetInfoParams, err error) {
	const op = "handler.GetInfo.paramsInfo"

	passportSerie := r.URL.Query().Get("passportSerie")
	p.PassportSerie, err = checkInt(passportSerie)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		return p, fmt.Errorf("%s: %s", op, err)
	}

	passportNumber := r.URL.Query().Get("passportNumber")
	p.PassportNumber, err = checkInt(passportNumber)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		return p, fmt.Errorf("%s: %s", op, err)
	}

	return p, nil
}

func checkparamsInt(r *http.Request, param string) (num int, err error) {
	// fmt.Println("___________paramsList3______", r)
	data := r.URL.Query().Get(param)
	// fmt.Println("___________paramsList4______", data)
	if data != "" {
		// data = "jok"
		// fmt.Println("___________paramsList5______", p)
		num, err = checkInt(data)

		// fmt.Println("___________paramsList6______", p)
		if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
			return 0, err
		}

	}
	fmt.Println("___________paramsList2______")

	return num, nil
}

func paramsList( /*w http.ResponseWriter,*/ r *http.Request) (swapi.GetListParams, error) {
	const op = "handler.GetInfo.paramsInfo"
	var p swapi.GetListParams

	// data := r.URL.Query().Get(passportSerie)

	// var err error
	surname := r.URL.Query().Get("surname")
	if surname != "" {
		p.Surname = &surname
	}

	name := r.URL.Query().Get("name")
	if name != "" {
		p.Name = &name
	}

	patronymic := r.URL.Query().Get("patronymic")
	if surname != "" {
		p.Patronymic = &patronymic
	}

	address := r.URL.Query().Get("address")
	if surname != "" {
		p.Address = &address
	}

	passportSerie, err := checkparamsInt(r, "passportSerie")
	// fmt.Println("___________paramsList22______", p)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		// fmt.Println("___________paramsList202______", err)
		return p, fmt.Errorf("%s: %s", op, err)
	}
	p.PassportSerie = &passportSerie

	passportNumber, err := checkparamsInt(r, "passportNumber")
	fmt.Println("___________paramsList23______", p)

	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		return p, fmt.Errorf("%s: %s", op, err)
	}
	p.PassportNumber = &passportNumber

	// validation
	fmt.Println("___________paramsList25______")

	page, err := checkparamsInt(r, "page")

	// fmt.Println("___________paramsList22______", p)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		// fmt.Println("___________paramsList202______", err)
		return p, fmt.Errorf("%s: %s", op, err)
	}
	if page == 0 {
		page = 1
	}
	p.Page = &page

	limit, err := checkparamsInt(r, "limit")

	// fmt.Println("___________paramsList22______", p)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		// fmt.Println("___________paramsList202______", err)
		return p, fmt.Errorf("%s: %s", op, err)
	}
	if limit == 0 {
		limit = 10
	}
	p.Limit = &limit

	fmt.Println("FFFFFFFFFF ", *p.Name, *p.PassportSerie)
	return p, nil
}

// !!!возможно не надо проверять параметры у нас же сваггер проверяет
// configureRouter сonfigures server routing for commands.
func (s *Server) ConfigureRouter(router *http.ServeMux) {
	// router := http.NewServeMux() // !!!возможно перенести сюда
	// s.Router.HandleFunc("/info3", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	router.HandleFunc("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Создание и заполнение структуры параметров
		s.log.Info("ConfigureRouter_GetInfo" /*, slog.String("id: ", "ded")*/) // vr may be
		params, err := paramsInfo(r)
		fmt.Println("___________/info______")
		if err != nil { //!!! вообще-то не должен выходить.. возможно убрать
			//усли пустой запрос ничего не делаем, если неккотректный ввод выводи ошибку в лог
			s.log.Error("incorrect params entered" /*, slog.String("id: ", "ded")*/) // !!!
			s.error(w, http.StatusBadRequest, err)
			return
		}
		_ = params

		s.GetInfo(w, r, params)
		s.log.Info("ConfigureRouter_f_GetInfo" /*, slog.String("id: ", "ded")*/) // vr may be

	}))

	router.HandleFunc("/list", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Создание и заполнение структуры параметров
		s.log.Info("ConfigureRouter_GetList" /*, slog.String("id: ", "ded")*/) // vr may be
		params, err := paramsList(r)
		fmt.Println("___________listPeople______", params)
		if err != nil {
			//усли пустой запрос ничего не делаем, если неккотректный ввод выводи ошибку в лог
			s.log.Error("incorrect params entered" /*, slog.String("id: ", "ded")*/) // !!!
			s.error(w, http.StatusBadRequest, err)
			return
		}
		_ = params

		// params := swapi.GetInfoParams{
		// 	PassportSerie:  passportSerie,
		// 	PassportNumber: passportNumber,
		// }
		// s.ServeHTTP(w, r)
		// fmt.Println("___________listPeople 88 ______", params)
		s.GetList(w, r, params)
		s.log.Info("ConfigureRouter_f_GetList" /*, slog.String("id: ", "ded")*/) // vr may be
	}))

	router.Handle("/swagger", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.Info("ConfigureRouter_/swagger" /*, slog.String("id: ", "ded")*/)
		w.Header().Set("Content-Type", "application/json")
		m := []byte("swagger.MarshalJSON()")
		w.Write(m)
	}))

	//  RO .Handle("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

}

// error generates a error to the client
func (s *Server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}

// response generates a response to the client
func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
