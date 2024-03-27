package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Login     string    `json:"login" db:"login" validate:"required,min=3,max=20"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"password"  db:"password" validate:"required"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}
