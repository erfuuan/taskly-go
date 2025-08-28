package db

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        os.Getenv("DATABASE_USER"),
        os.Getenv("DATABASE_PASSWORD"),
        os.Getenv("DATABASE_HOST"),
        os.Getenv("DATABASE_PORT"),
        os.Getenv("DATABASE_NAME"),
    )

    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        return err
    }

    DB = pool
    return nil
}

func Close() {
    if DB != nil {
        DB.Close()
    }
}
