package test

import (
	"go-user-api/config"
	"testing"
)

func TestNewMysql(t *testing.T) {
	viper := config.NewViper()
	db := config.NewMysql(viper)
	defer db.Close()
}
