package users

import (
	"encoding/json"
	"fmt"
	users2 "github.com/daniskazan/url-shortener/internal/core/users"
	"github.com/daniskazan/url-shortener/internal/repository/users"
	"net/http"
)

type UserHandler struct {
	Repo *users.UserRepository
}

func (h *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var body struct {
		Name  string `json:"name" validate:"required,omitempty"`
		Email string `json:"email" validate:"required,omitempty"`
	}

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user := users2.User{Name: body.Name, Email: body.Email}

	h.Repo.Save(user)

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write(res)
}

func (h *UserHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Update user"))
}

func (h *UserHandler) GetUserByID(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Get by ID"))
}

func (h *UserHandler) GetUserList(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Get User List"))
}

func (h *UserHandler) DeleteUserByID(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Delete User"))

}
