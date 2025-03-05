package token

import (
	"go-user-api/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateAccessToken(id string, viper *viper.Viper) (string, error) {
	claims := model.JwtCustomClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(viper.GetInt("jwt.exp")))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(viper.GetString("jwt.key")))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func VerifyToken(tokenString string, config *viper.Viper) (*model.JwtCustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*model.JwtCustomClaim)
	claimType := &model.JwtCustomClaim{
		Id: claims.Id,
	}
	return claimType, nil
}
