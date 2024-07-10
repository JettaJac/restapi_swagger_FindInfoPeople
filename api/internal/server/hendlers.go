package server

import (
	// "error"
	// "encoding/json"
	// "context"
	"errors"
	"fmt"
	"log/slog"
	swapi "main/generated"
	"main/internal/domain/models"
	"main/internal/service/people"
	"main/pkg/lib/logger"
	"net/http"
)

// !!! Нужно ли делать валидацию запроса?
// !!!Bpvtybnm Server на ServerAPI
var (
	ErrMethod = errors.New("ErrMethod")
)

func (s *Server) GetInfo(w http.ResponseWriter, r *http.Request, p swapi.GetInfoParams) {
	// if r.Method == http.MethodOptions {
	// 	// Обработка предварительного запроса OPTIONS
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }

	if r.Method == http.MethodGet {
		const op = "server.GetInfo"

		// log := *s.log.With( //возможно сделать надо s.log
		// 	slog.String("op", op),
		// )

		var req *models.User

		// var ctx context.Context // сделать правильно !!!

		req, err := s.info.GetInfo(p) // возможно надо ctx
		// req, err := s.storage.GetInfo(ctx, p)

		query := fmt.Sprintf("%d %d", p.PassportSerie, p.PassportNumber)
		s.log.Info(op, slog.Any("request", req))
		if errors.Is(err, people.ErrUserNotFound) { //!!! возможно сменить нейминк
			s.log.Error("info not found", slog.String("params: ", query))
			s.error(w, http.StatusNotFound, err)
			return
		}
		if err != nil {
			s.log.Error("failed to get command by pasport", sl.Err(err))
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.log.Info("got infa about people", slog.Int("id", int(req.ID))) //!!!

		s.respond(w, http.StatusOK, req) //!!! Должен выдавать Json с инфой
	} else {
		s.log.Error("incorrect request method, need a GET")
		s.error(w, http.StatusMethodNotAllowed, ErrMethod)
		return
	}
}

// handleHome returns the home page
func (s *Server) HandleHome() http.HandlerFunc { // не нужен
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// Обработка предварительного запроса OPTIONS
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodGet {
			const op = "server.GetInfo"

			log := *s.log.With(
				slog.String("op", op),
			)
			log.Info("handleHome")

			s.respond(w, 200, "Начнем!!!")
			w.Write([]byte("Hello World!"))
		} else {
			// var err error
			// err = errors.New("handleHome-bad") //!!!
			// s.log.Error("incorrect request method, need a POST") //!!! Нужно ли через сервер пробрасывать лог
			s.error(w, http.StatusMethodNotAllowed, ErrMethod)
			return
		}
	}
}
