package services

import (
	"github.com/moz5691/microservices/mvc/domain"
	"github.com/moz5691/microservices/mvc/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.AppliationErrors) {
	return domain.GetUser(userId)
}
