package repository

import "github.com/antyomusa/go-rest-api/internal/entities"

func GetAllUsers() []entities.User{
	return []entities.User{
		{ID: 1, Name: "Anthony"},
		{ID: 2, Name: "John"},
	}
}