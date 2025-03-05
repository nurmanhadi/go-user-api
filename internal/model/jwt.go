package model

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaim struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}
