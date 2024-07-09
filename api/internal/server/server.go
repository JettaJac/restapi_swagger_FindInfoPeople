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

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.Server.Handler.ServeHTTP(w, r)
// }

// configureRouter сonfigures server routing for commands.
func (s *Server) ConfigureRouter(router *http.ServeMux) {
	// router := http.NewServeMux() // !!!возможно перенести сюда
	// s.Router.HandleFunc("/info3", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	router.HandleFunc("/info", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Создание и заполнение структуры параметров
		s.log.Info("ConfigureRouter_GetInfo" /*, slog.String("id: ", "ded")*/) // vr may be
		params, err := paramsInfo( /*w,*/ r)
		fmt.Println("___________/info______")
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

		s.GetInfo(w, r, params)
		s.log.Info("ConfigureRouter_GetInfo" /*, slog.String("id: ", "ded")*/) // vr may be

		// router.ServeHTTP(w, r)
		// m := []byte("swagger.MarshalJSON()")
		// w.Write(m)
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
