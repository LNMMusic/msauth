package handler

import (
	"errors"
	"net/http"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/internal/user/storage"
	"github.com/LNMMusic/msauth/pkg/web/request"
	"github.com/LNMMusic/msauth/pkg/web/response"
	"github.com/LNMMusic/msauth/pkg/web/validator"
	"github.com/LNMMusic/optional"
)

// NewHandlersRegister returns a new HandlersRegister struct
func NewHandlersRegister(stWrite storage.StorageWrite) *HandlersRegister {
	return &HandlersRegister{
		stWrite: stWrite,
	}
}

// HandlersRegister is the struct that contains the dependencies for the register handlers
type HandlersRegister struct {
	// stWrite is the write interface for the user store
	stWrite storage.StorageWrite
}

// SignUp is the handler for the sign up route
type UserSignUp struct {
	// Username is the username of the user
	Username 	optional.Option[string] `json:"username"`
	// Password is the password of the user
	Password 	optional.Option[string] `json:"password"`
	// Email is the email of the user
	Email 		optional.Option[string] `json:"email"`
}
func (h *HandlersRegister) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body: validate
		err := validator.RequiredJSON(r.Body, "username", "password", "email")
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "missing required fields")
			return
		}
		// - body: decode
		var userSignUp UserSignUp
		err = request.JSON(r, &userSignUp)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid json")
			return
		}

		// process
		// - deserialize
		u := user.User{
			Username: userSignUp.Username,
			Password: userSignUp.Password,
			Email: userSignUp.Email,
		}
		// - store
		err = h.stWrite.Create(&u)
		if err != nil {
			switch {
			case errors.Is(err, storage.ErrStorageExists):
				response.JSON(w, http.StatusConflict, "user already exists")
			case errors.Is(err, storage.ErrStorageInvalid):
				response.JSON(w, http.StatusUnprocessableEntity, "invalid user")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "user created",
			"data": map[string]any{
				"username": u.Username,
				"email": u.Email,
			},
		})
	}
}