package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Eslam-Amin/ecommerce/types"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	// get JSON Payload
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, resBody types.ResponseBody) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(types.ResponseBody{
		Success: resBody.Success,
		Message: resBody.Message,
		Data:    resBody.Data,
		Error:   resBody.Error,
	})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, types.ResponseBody{
		Success: false,
		Error:   err.Error(),
	})
}

var Validate = validator.New()

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
