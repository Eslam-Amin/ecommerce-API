package auth

import (
	"time"

	"github.com/Eslam-Amin/ecommerce/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenStr, nil

}
