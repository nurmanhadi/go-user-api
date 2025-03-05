package model

type AuthRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}
type AuthResponseAccessToken struct {
	AccessToken string `json:"access_token"`
}
