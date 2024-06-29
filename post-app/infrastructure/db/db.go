package db

import (
	"context"
	"database/sql"
)

type DB interface {
	InitDB() (*sql.DB, error)
}

func NewDB() (DB, error) {
	return &PostgresDB{}, nil
}

type DBManeger struct {
	pool *sql.DB
}

func NewDBManeger(db DB) (*DBManeger, error) {
	pool, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	return &DBManeger{
		pool: pool,
	}, nil
}

func (m *DBManeger) GetPool() *sql.DB {
	return m.pool
}

func (m *DBManeger) Close() error {
	if m.pool != nil {
		return m.pool.Close()
	}

	return nil
}

type DBTxManeger struct {
	tx *sql.Tx
}

func NewDBTxManeger(
	dbManeger *DBManeger,
	ctx context.Context,
	txOptions sql.TxOptions,
) (*DBTxManeger, error) {
	transaction, err := dbManeger.GetPool().BeginTx(
		ctx,
		&txOptions,
	)
	if err != nil {
		return nil, err
	}
	return &DBTxManeger{
		tx: transaction,
	}, nil
}

func NewDBTxManegerWithPool(
	pool *sql.DB,
	ctx context.Context,
	txOptions sql.TxOptions,
) (*DBTxManeger, error) {
	transaction, err := pool.BeginTx(
		ctx,
		&txOptions,
	)
	if err != nil {
		return nil, err
	}
	return &DBTxManeger{
		tx: transaction,
	}, nil
}

func (tm *DBTxManeger) GetTx() *sql.Tx {
	return tm.tx
}

func (tm *DBTxManeger) CommitTx() error {
	return tm.tx.Commit()
}

func (tm *DBTxManeger) RollbackTx() error {
	return tm.tx.Rollback()
}
