package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const jwtSecret = "g6VT7Z8mY4EIoSD2XgPmjHWKPcPR1nKj"

func CreateToken(username string, email string) (string, error) {
	claims := newRedesignClaims(username, email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(signedToken string) (*RedesignClaims, error) {
	if signedToken == "" {
		return nil, ErrInvalidToken
	}
	claims := &RedesignClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtSecret), nil
	}

	tkn, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, ErrSignatureInvalid
		}
		return nil, errors.Wrap(ErrBadRequest, err.Error())
	}
	if !tkn.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func GenerateWebhookToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = 0
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "Webhook token signing Error", err
	}

	return tokenString, nil
}

func VerifyWebhookToken(token string) error {
	if token == "" {
		return ErrInvalidToken
	}
	claims := &RedesignClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtSecret), nil
	}

	tkn, err := jwt.ParseWithClaims(
		token,
		claims,
		keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return ErrSignatureInvalid
		}
		return errors.Wrap(ErrBadRequest, err.Error())
	}
	if !tkn.Valid {
		return ErrInvalidToken
	}

	return nil
}
