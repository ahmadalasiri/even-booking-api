package models

import (
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
