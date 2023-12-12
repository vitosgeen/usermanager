package datastore

import (
	"fmt"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/config"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	SQL *sqlx.DB
}

const (
	maxOpenDbConn   = 5
	maxIdleDbConn   = 5
	maxDbLifeTime   = 5 * time.Minute
	connMaxIdleTime = 20
	driverName      = "pgx"
)

func NewDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=%s",
		cfg.Postgres.PostgresHost,
		cfg.Postgres.PostgresPort,
		cfg.Postgres.PostgresUser,
		cfg.Postgres.PostgresDBName,
		cfg.Postgres.PostgresPass,
		cfg.Postgres.PostgresTimezone,
	)

	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		return nil, apperrors.SqlOpenError.AppendMessage(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	err = pingDB(db)
	if err != nil {
		return nil, apperrors.PingDBError.AppendMessage(err)
	}

	return &DB{
		SQL: db,
	}, nil
}

func pingDB(db *sqlx.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}
