package server

import (
	// "error"
	// "encoding/json"
	"errors"
	swapi "main/generated"
	"net/http"
)

func (s *Server) GetInfo(w http.ResponseWriter, r *http.Request, params swapi.GetInfoParams) {
	s.respond(w, http.StatusCreated, "_____________")
	// w.WriteHeader(http.StatusOK)
	// _ = json.NewEncoder(w).Encode("result")
}

// handleHome returns the home page
func (s *Server) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
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
