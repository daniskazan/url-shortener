package users

import (
	"github.com/daniskazan/url-shortener/internal/repository/users"
	"net/http"
)

type UserHandler struct {
	repo *users.UserRepository
}

func (h *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Create user"))
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
