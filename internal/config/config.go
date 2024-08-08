package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Conn *pgxpool.Conn

// type database struct {
// 	User     string
// 	Password string
// 	Host     string
// 	Port     string
// 	Name     string
// }

func init() {
	Conn = openDB()
}

func openDB() *pgxpool.Conn {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	pool, err := pgxpool.New(context.Background(), dsn)
	// conn, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Failed to acquire connection")
	}

	defer conn.Release()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// conn.Close()

	return conn

}
