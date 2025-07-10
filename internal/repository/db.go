package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	_ = godotenv.Load(".env")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, pwd, host, port, dbname)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco com sucesso!")
	return db, nil
}
