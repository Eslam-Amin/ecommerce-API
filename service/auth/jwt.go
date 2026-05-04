package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Eslam-Amin/ecommerce/config"
	"github.com/Eslam-Amin/ecommerce/types"
	"github.com/Eslam-Amin/ecommerce/utils"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserKey contextKey = "userId"

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

func WithJWTAut(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := getTokenFromRequest(r)

		token, err := validateToken(tokenStr)

		if err != nil {
			log.Printf("faile to validate token %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIdStr := claims["userId"].(string)

		userId, _ := strconv.Atoi(userIdStr)

		user, err := store.GetUserById(userId)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.ID)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}

}

func getTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Auhorization")
}

func validateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userId
}
