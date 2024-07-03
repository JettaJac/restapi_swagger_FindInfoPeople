package server

import (
	"encoding/json"
	swapi "main/generated"
	"main/internal/config"
	"net/http"
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
}

var _ swapi.ServerInterface = (*Server)(nil)

// type serverAPI struct {
// 	// ssov1.UnimplementedAuthServer
// 	// auth Auth
// }

// NewServer create a new server
func NewServer(config *config.Config) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         config.Address,
			ReadTimeout:  config.HTTPServer.Timeout,
			WriteTimeout: config.HTTPServer.Timeout,
			IdleTimeout:  config.HTTPServer.IdleTimeout,
			Handler:      http.NewServeMux(),
		},
		//!!! Нужно ли через сервер пробрасывать лог
	}
	// s.configureRouter()

	return s
}

// configureRouter сonfigures server routing for commands.
// func (s *Server) configureRouter() {
// 	// s.Router.HandleFunc("/", s.handleHome())
// 	s.Router.HandleFunc("/info", s.handleHome())
// }

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
