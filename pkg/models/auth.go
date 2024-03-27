package models

type AuthLoginResponse struct {
	Token    string `json:"token"`
	FullName string `json:"full_name"`
}
