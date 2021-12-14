package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

func ConnectSQL() (*DB, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading file .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5)

	dbConn.SQL = db
	return dbConn, nil

}