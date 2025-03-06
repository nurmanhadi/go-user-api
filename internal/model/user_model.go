package model

import "time"

type UserResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserChangeNameRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}
