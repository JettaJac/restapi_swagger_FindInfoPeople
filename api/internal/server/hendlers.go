package server

import (
	// "error"
	// "encoding/json"
	"errors"
	"log/slog"
	swapi "main/generated"
	"net/http"
)

func (s *Server) GetInfo(w http.ResponseWriter, r *http.Request, params swapi.GetInfoParams) {
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

		log.Info("GetInfo_log", slog.Any("request", r))   // !!! не надо
		s.log.Info("GetInfo", slog.Any("request", "req")) //!!! не надо

		// result, err := s.storage.GetInfo()
		// 	if err != nil {
		// 		s.log.Error("failed to add command", sl.Err(err))
		// 		s.error(w, http.StatusInternalServerError, err)
		// 		return
		// 	}

		s.respond(w, http.StatusOK, params.PassportSerie)
	} else {
		var err error
		err = errors.New("handleHome-bad") //!!!
		// s.log.Error("incorrect request method, need a POST") //!!! Нужно ли через сервер пробрасывать лог
		s.error(w, http.StatusMethodNotAllowed, err)
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
			var err error
			err = errors.New("handleHome-bad") //!!!
			// s.log.Error("incorrect request method, need a POST") //!!! Нужно ли через сервер пробрасывать лог
			s.error(w, http.StatusMethodNotAllowed, err)
			return
		}
	}
}
