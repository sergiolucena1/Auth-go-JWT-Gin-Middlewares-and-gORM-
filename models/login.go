package models

// criando nosso login

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}