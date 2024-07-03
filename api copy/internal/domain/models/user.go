package models

type User struct {
	ID             int64  `json:"id"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address"`
	PassportSerie  int    `json:"passportSerie"`
	PassportNumber int    `json:"passportNumber"`
}
