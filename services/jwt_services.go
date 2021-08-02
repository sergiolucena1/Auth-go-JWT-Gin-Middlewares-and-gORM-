package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// gerar o JWT

type jwtService struct {
	secretKey string
	issure string
}
func NewJWTService() *jwtService{
	return &jwtService{
		secretKey: "secret-key",
		issure: "blogposts-api",
	}
}

//Claim : campos do JWT
type  Claim struct {
	Sum uint `json: "sun"`
	jwt.StandardClaims
}

//função de gerar o token
func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour *2).Unix(),// quando o token vai expirar(2horas, finalizando com unix()
			Issuer: s.issure, //  emissor
			IssuedAt: time.Now().Unix(), // tempo que foi emitido
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))// assinando a string
	if err != nil{
		return "", err
	}
	return t, nil
}

// Método de validação do token
func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { //  analisar informacoes do token
		//se o nosso token é valido ou n
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil //se der erro ele retorna false
					// se der nil ele retorna true(token valido)

}
