package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"Gin-Api/config"

	_ "github.com/lib/pq"

	"github.com/pkg/errors"
)

func OpenDB(cfg *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "open sql")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "connection is not established")
	}

	log.Println("Connected to DB")

	return db, nil
}
