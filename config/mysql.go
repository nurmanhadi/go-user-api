package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func NewMysql(viper *viper.Viper) *sql.DB {
	dsn := viper.GetString("mysql.dsn")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("cannot connect to mysql: %s", err)
	}
	db.SetMaxIdleConns(viper.GetInt("db.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("db.max_open_conns"))
	db.SetConnMaxIdleTime(time.Minute * time.Duration(viper.GetInt("db.conn_max_idletime")))
	db.SetConnMaxLifetime(time.Minute * time.Duration(viper.GetInt("db.conn_max_lifetime")))

	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping to database: %s", err)
	}
	return db
}
