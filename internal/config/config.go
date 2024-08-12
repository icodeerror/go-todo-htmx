package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// var Conn *pgxpool.Conn

// type database struct {
// 	User     string
// 	Password string
// 	Host     string
// 	Port     string
// 	Name     string
// }

// func init() {
// 	Conn = openDB()
// }

func OpenDB() (*pgxpool.Pool, *pgxpool.Conn, error) {

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
		return nil, nil, err
	}

	// defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Failed to acquire connection")
		return nil, nil, err
	}

	// defer conn.Release()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Println("Ping", err)
		log.Fatal(err)
	}

	query := `
						CREATE TABLE IF NOT EXISTS todos(
						id SERIAL PRIMARY KEY,
						description TEXT NOT NULL,
						completed BOOLEAN NOT NULL
						)
					`

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	return pool, conn, nil

}
