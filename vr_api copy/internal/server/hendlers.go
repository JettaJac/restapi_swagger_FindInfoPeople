package server

import (
	// "error"
	"log"
	swapi "main/generated"
	"net/http"
)

// func (s *Server) setupHandlers() {
// 	// Обработчик для запроса GetInfo
// 	// s.handlers["GetInfo"] = func(ctx *gin.Context) {
// 	// 	// Получение ID пользователя из запроса
// 	// 	userID := ctx.Param("userID")

// 	// 	// Логика получения информации о пользователе
// 	// 	userInfo, err := s.getUserInfo(userID)
// 	// 	if err != nil {
// 	// 		ctx.AbortWithStatus(http.StatusInternalServerError)
// 	// 		return
// 	// 	}

// 	// 	// Возврат информации о пользователе в формате, определенном в Swagger
// 	// 	ctx.JSON(http.StatusOK, userInfo)
// 	// }

// 	// // Обработчик для запроса SaveInfo
// 	// s.handlers["SaveInfo"] = func(ctx *gin.Context) {
// 	// 	// Получение данных пользователя из тела запроса
// 	// 	var newUserInfo UserInfo
// 	// 	if err := ctx.BindJSON(&newUserInfo); err != nil {
// 	// 		ctx.AbortWithStatus(http.StatusBadRequest)
// 	// 		return
// 	// 	}

// 	// 	// Логика сохранения информации о новом пользователе
// 	// 	err := s.saveUserInfo(newUserInfo)
// 	// 	if err != nil {
// 	// 		ctx.AbortWithStatus(http.StatusInternalServerError)
// 	// 		return
// 	// 	}

// 	// 	// Возврат подтверждения успешного сохранения
// 	// 	ctx.JSON(http.StatusOK, gin.H{
// 	// 		"message": "User information saved successfully",
// 	// 	})
// 	// }
// }

func (s *Server) GetInfo(w http.ResponseWriter, r *http.Request, params swapi.GetInfoParams) {
	// s.log.Info("command added", slog.Int("id", id))
	log.Println(params)
	var err error
	s.error(w, http.StatusCreated, err)
	s.respond(w, http.StatusCreated, "req")
}
