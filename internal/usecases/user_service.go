package usecases

import (
	"github.com/antyomusa/go-rest-api/internal/entities"
	"github.com/antyomusa/go-rest-api/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetUsers() ([]entities.User, error) {
	return s.Repo.GetAllUsers()
}