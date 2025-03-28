package postgres

import (
	"fmt"
	"time"

	"github.com/umed-hotamov/realty-service/internal/app/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	maxConns        = 50
	maxIdleConns    = 25
	maxConnLifeTime = 10 * time.Minute
	maxIdleConnTime = 5 * time.Minute
)

func NewPostgresDB(cfg *config.Config, logger *zap.Logger) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.User,
		cfg.Postgres.Password)

	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		logger.Fatal("failed to connect to postgres: %s", zap.String("dsn", dsn))
	}

	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxConnLifeTime)
	db.SetConnMaxIdleTime(maxIdleConnTime)

	return db
}
