package database

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
)

var (
	db  *sql.DB
	ctx context.Context
)

func Connect(dsnUri string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsnUri)

	if err != nil {
		return nil, err
	}

	if db.Ping() != nil {
		return nil, err
	}

	return db, nil
}

func SelectQuery(query string, args ...string) (*sql.Rows, error) {
	transaction, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelLinearizable})
	if err != nil {
		return nil, err
	}

	result, execErr := transaction.Query(query, args)
	if execErr != nil {
		_ = transaction.Rollback()
		return nil, execErr
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
