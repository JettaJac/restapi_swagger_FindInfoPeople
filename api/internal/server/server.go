package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	swapi "main/generated"
	"main/internal/config"
	"net/http"
	"strconv"
)

var (
// GetInfo  = "/info"   // GetInfo
)

// Server struct
type Server struct {
	*http.Server
	// Router *http.ServeMux
	// Addr   string
	// // authService  server.Auth
	// ReadTimeout  time.Duration
	// WriteTimeout time.Duration
	// IdleTimeout  time.Duration
	log *slog.Logger
}

var _ swapi.ServerInterface = (*Server)(nil)

// type serverAPI struct {
// 	// ssov1.UnimplementedAuthServer
// 	// auth Auth
// }

// NewServer create a new server
func NewServer(config *config.Config, log *slog.Logger) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         config.HTTPServer.Address,
			ReadTimeout:  config.HTTPServer.Timeout,
			WriteTimeout: config.HTTPServer.Timeout,
			IdleTimeout:  config.HTTPServer.IdleTimeout,
			Handler:      http.NewServeMux(),
		},
		log: log,
		//!!! Нужно ли через сервер пробрасывать лог
	}

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
	return num, nil

}

func paramsInfo( /*w http.ResponseWriter,*/ r *http.Request) (params swapi.GetInfoParams, err error) {
	const op = "handler.GetInfo.paramsInfo"

	passportSerie := r.URL.Query().Get("passportSerie")
	params.PassportSerie, err = checkInt(passportSerie)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		return params, fmt.Errorf("%s: %s", op, err)
	}

	passportNumber := r.URL.Query().Get("passportNumber")
	params.PassportNumber, err = checkInt(passportNumber)
	if err != nil { // !!! обработать ошибку, чтоб падал запрос в случае не корректности
		return params, fmt.Errorf("%s: %s", op, err)
	}

	return params, nil
}

// configureRouter сonfigures server routing for commands.
func (s *Server) ConfigureRouter(router *http.ServeMux) {
	// router := http.NewServeMux() // !!!возможно перенести сюда
	router.Handle("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Создание и заполнение структуры параметров
		params, err := paramsInfo( /*w,*/ r)
		fmt.Println(err)
		if err != nil {
			//усли пустой запрос ничего не делаем, если неккотректный ввод выводи ошибку в лог
			s.log.Error("incorrect params entered" /*, slog.String("id: ", "ded")*/) // !!!
			s.error(w, http.StatusBadRequest, err)
			return
		}
		// params := swapi.GetInfoParams{
		// 	PassportSerie:  passportSerie,
		// 	PassportNumber: passportNumber,
		// }

		s.GetInfo(w, r, params)
	}))

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
