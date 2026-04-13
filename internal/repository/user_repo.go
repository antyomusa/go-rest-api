package repository

import (
	"database/sql"
"github.com/antyomusa/go-rest-api/internal/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetAllUsers() ([]entities.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE email=$1", email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID)
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)`

	err := r.DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func (r *UserRepository) SaveRefreshToken(userID int, token string) error {
	_, err := r.DB.Exec("UPDATE users SET refresh_token=$1 WHERE id=$2", token, userID)
	return err
}
