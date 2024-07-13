package db

import (
	"context"
	"database/sql"

	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/config"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/logger"
)

type DB interface {
	InitDB(dbConfig *config.DBConfig) (*sql.DB, error)
}

func NewDB() (DB, error) {
	return &PostgresDB{}, nil
}

type DBManeger struct {
	pool *sql.DB
}

func NewDBManeger(db DB, dbConfig *config.DBConfig) (*DBManeger, error) {
	pool, err := db.InitDB(dbConfig)
	if err != nil {
		logger.Error("DBの初期化に失敗しました", "err", err)
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
	logger.Info(
		"DBのトランザクションを開始します",
		"MaxOpenConnections",
		pool.Stats().MaxOpenConnections,
		"OpenConnections",
		pool.Stats().OpenConnections,
		"InUse",
		pool.Stats().InUse,
		"Idle",
		pool.Stats().Idle,
	)
	transaction, err := pool.BeginTx(
		ctx,
		&txOptions,
	)
	if err != nil {
		logger.Error("DBのトランザクションの開始に失敗しました", "err", err)
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
