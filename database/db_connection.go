package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"

	// blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//DbConn stores the connexion to the database
var (
	DbConn *sql.DB
)

// Config for DB connection
type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
	DbConn     *sql.DB
}

// Connect connection to database
func Connect() error {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("%+v", err)
	}
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + "/" + cfg.
		DbName + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	var dbErr error
	for i := 1; i <= 8; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			if i < 8 {
				log.Printf("db connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}

	if dbErr != nil {
		return errors.New("can't connect to database after 3 attempts")
	}

	DbConn = db

	return nil
}
