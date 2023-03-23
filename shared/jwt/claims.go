package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RedesignClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func newRedesignClaims(username string, email string) *RedesignClaims {
	claims := &RedesignClaims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "redesign",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	return claims
}
