package models

import (
	"errors"
	"event-booking-api/db"
	"event-booking-api/utils"
)

type User struct {
	ID        int
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	CreatedAt string
}

func (u User) Save() (User, error) {

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, created_at
	`
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return User{}, err
	}

	var user User
	err = db.DB.QueryRow(query, u.Email, hashedPassword).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u User) ValidateCredentials() error {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`
	var user User
	err := db.DB.QueryRow(query, u.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	isValidPassword := utils.ComparePasswords(user.Password, u.Password)
	if !isValidPassword {
		return errors.New("invalid credentials")
	}

	return nil
}

func (u User) FindByEmail() (User, error) {
	query := `
		SELECT id, email
		FROM users
		WHERE email = $1
	`
	var user User
	err := db.DB.QueryRow(query, u.Email).Scan(&user.ID, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
