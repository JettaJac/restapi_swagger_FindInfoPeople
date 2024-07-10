package server

import (
	// "errors"
	"fmt"
	// "log/slog"
	swapi "main/generated"
	"main/internal/domain/models"
	"main/pkg/lib/logger"
	"net/http"
)

// type Serverh struct {
// 	server *server.Server
// }

// !!!NTcns на лимит и тест на page
func (s *Server) GetList(w http.ResponseWriter, r *http.Request, p swapi.GetListParams) {

	// if r.Method == http.MethodOptions {
	// 	// Обработка предварительного запроса OPTIONS
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }

	if r.Method == http.MethodGet {
		const op = "server.GetList"

		// log := *s.log.With(
		// 	slog.String("op", op),
		// )
		fmt.Println("___________listPeople 90 ______", p)
		var req []models.User
		_ = req

		// if p.Page == nil {
		// 	*p.Page = 1
		// }

		// if p.Limit == nil {
		// 	*p.Limit = 10
		// }

		// var ctx context.Context // сделать правильно !!!
		// fmt.Println("___________listPeople 91 ______", p)
		fmt.Println(*p.PassportSerie, "service _IIIIIIIIIIIII", *p.Name, *p.PassportNumber)
		listPeoples, err := s.info.GetList(p) // возможно надо ctx
		// req, err := s.storage.GetInfo(ctx, p)

		if err != nil {
			s.log.Error("failed to get info about people", sl.Err(err))
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		// s.log.Info("infa about people uploaded", slog.Int("page", int(*p.Page))) //!!!
		s.respond(w, http.StatusOK, listPeoples) //!!! Должен выдавать Json с инфой
	} else {
		s.log.Error("incorrect request method, need a GET")
		s.error(w, http.StatusMethodNotAllowed, ErrMethod)
		return
	}
}
