package server

import (
	"encoding/json"
	// middleware "github.com/oapi-codegen/nethttp-middleware"
	// "log/slog"
	// "github.com/go-openapi/spec"
	swapi "main/generated"
	// api3 "github.com/oapi-codegen/oapi-codegen"
	"main/internal/config"
	// "main/internal/storage"
	// "context"
	// "main/internal/domain/models"
	// "main/internal/service/people"
	"net/http"
	"time"
)

var (
// PathSave   = "/command/save"   // handleSaveRunCommand
// PathFind   = "/command/find"   //HandleGetOneCommand
// PathList   = "/commands/all"   // HandleGetListCommands
// PathDelete = "/command/delete" // HandleDeleteCommand
)

// Server struct
type Server struct {
	Router *http.ServeMux
	Addr   string
	// authService  server.Auth
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// type PeopleProvider interface {
// 	GetInfo(ctx context.Context, PassportSerie, PassportNumber int64) (models.User, error)
// 	// ProvideAllPosts(ctx context.Context, page int64) ([]models.Post, error)
// }

var _ swapi.ServerInterface = (*Server)(nil)

type serverAPI struct {
	// ssov1.UnimplementedAuthServer
	// auth Auth
}

// func (p *people.People) GetInfo() {

// }

// ServeHTTP routes HTTP requests using router // !!!Может и не надо
// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.Router.ServeHTTP(w, r)
// }

// NewServer create a new server
func NewServer(config *config.Config /*, swagger *api3.openapi3.T*/) *Server {
	s := &Server{
		Router:       http.NewServeMux(),
		Addr:         config.Address,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	// _ = swagger
	// Настройка промежуточного ПО для валидации запросов
	// s.Handler = middleware.OapiRequestValidator(swagger)(router) //!!! если адо настроить для своего сервера
	//!!! если адо настроить для своего сервера
	// h = middleware.OapiRequestValidator(swagger)(s.Router)
	// s.configureRouter()
	return s
}

// configureRouter сonfigures server routing for commands.
// func (s *Server) configureRouter() {
// 	// s.router.HandleFunc("/", s.handleHome())
// 	// s.router.HandleFunc(PathSave, s.handleSaveRunCommand(*s.log))
// 	// s.router.HandleFunc(PathFind, s.handleGetOneCommand(*s.log))
// 	// s.router.HandleFunc(PathList, s.handleGetListCommands(*s.log))
// 	// s.router.HandleFunc(PathDelete, s.handleDeleteCommand(*s.log))
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
