package usecases

import (
	"errors"

	"github.com/antyomusa/go-rest-api/internal/entities"
	"github.com/antyomusa/go-rest-api/internal/repository"
	"github.com/antyomusa/go-rest-api/security"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !security.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid password")
	}

	token, _ := security.GenerateToken(email)
	refreshToken, _ := security.GenerateRefreshToken(email)

	s.Repo.SaveRefreshToken(user.ID, refreshToken)

	return token, refreshToken, nil
}

func (s *AuthService) Register(name, email, password string) (string, string, error) {

	// cek email sudah ada
	exists, err := s.Repo.ExistsByEmail(email)
	if err != nil {
		return "", "", err
	}
	if exists {
		return "", "", errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return "", "", err
	}

	// set default role
	user := &entities.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "USER",
	}

	// simpan ke DB
	err = s.Repo.CreateUser(user)
	if err != nil {
		return "", "", err
	}

	// generate token (optional, biar auto login)
	token, _ := security.GenerateToken(user.Email)
	refreshToken, _ := security.GenerateRefreshToken(user.Email)

	// simpan refresh token
	s.Repo.SaveRefreshToken(user.ID, refreshToken)

	return token, refreshToken, nil
}