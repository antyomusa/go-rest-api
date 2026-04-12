package usecases

import (
	"github.com/antyomusa/go-rest-api/internal/entities"
	"github.com/antyomusa/go-rest-api/internal/repository"
)

func GetUsers() []entities.User {
	users := repository.GetAllUsers()

	return users;
}