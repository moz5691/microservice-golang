package domain

import (
	"fmt"
	"github.com/moz5691/microservices/mvc/utils"
	"net/http"
)

var (
	users = map[uint64]*User{
		123: &User{Id: 123, FirstName: "John", LastName: "Denver", Email: "jd@mail.com"},
	}
)

func GetUser(userId uint64) (*User, *utils.AppliationErrors) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppliationErrors{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
