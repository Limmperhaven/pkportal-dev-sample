package stpg

import (
	"context"
	"errors"
	"fmt"
	"github.com/Limmperhaven/pkportal-be-v2/internal/config"
	"github.com/Limmperhaven/pkportal-be-v2/internal/storage"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"runtime/debug"
	"strconv"
	"time"
)

var errConnect = errors.New("config is empty or connect is not init")

type Storage struct {
	db *sqlx.DB
}

var instance *Storage

func InitConnect(cfg *config.Postgres) error {
	if cfg == nil {
		return errConnect
	}
	//if cfg.Host == "" {
	//	cfg.Host = "localhost"
	//}
	if cfg.Port == "" {
		cfg.Port = "5432"
	}
	if cfg.User == "" {
		cfg.User = "postgres"
	}
	if cfg.Password == "" {
		cfg.Password = "postgres"
	}
	if cfg.DbName == "" {
		cfg.DbName = "pk-portal"
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = 50
	}
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 50
	}

	port, _ := strconv.Atoi(cfg.Port)

	connConfig := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(port),
		Database: cfg.DbName,
		User:     cfg.User,
		Password: cfg.Password,
	}
	conPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: cfg.MaxOpenConns,
		AfterConnect:   nil,
		AcquireTimeout: 3 * time.Minute,
	})
	if err != nil {
		return err
	}

	sqlDb := stdlib.OpenDBFromPool(conPool)
	sqlxDb := sqlx.NewDb(sqlDb, "pgx")
	instance = &Storage{db: sqlxDb}

	return nil
}

func Gist() storage.PGer {
	if instance == nil {
		return nil
	}
	return instance
}

func (st *Storage) DBSX() *sqlx.DB {
	if instance == nil {
		return nil
	}
	return instance.db
}

func (st *Storage) QueryTx(ctx context.Context, f func(tx *sqlx.Tx) error) (err error) {
	if instance == nil {
		return errConnect
	}

	tx, err := instance.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	commit := false
	defer func() {
		if r := recover(); r != nil || !commit {
			if r != nil {
				err = fmt.Errorf("transaction panic: %s\n%s", r, string(debug.Stack()))
				_ = tx.Rollback()
			} else if e := tx.Rollback(); e != nil {
				err = e
			}
		} else if commit {
			if e := tx.Commit(); e != nil {
				err = e
			}
		}
	}()

	if err := f(tx); err != nil {
		return err
	}

	commit = true
	return nil
}
