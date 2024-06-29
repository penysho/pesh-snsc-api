package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PostgresDB struct {
}

func (p *PostgresDB) InitDB() (*sql.DB, error) {

	db, err := sql.Open(
		"postgres",
		"user=postgres password=postgres dbname=postgres host=db port=5432 sslmode=disable connect_timeout=10",
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// https://pkg.go.dev/database/sql#DB.SetConnMaxIdleTime
	// https://please-sleep.cou929.nu/go-sql-db-connection-pool.html
	db.SetMaxOpenConns(10)                   // 同時に使用できる最大コネクション数
	db.SetMaxIdleConns(5)                    // コネクションプール内に保持する最大アイドルコネクション数
	db.SetConnMaxLifetime(300 * time.Second) // コネクションが利用可能な最長時間
	db.SetConnMaxIdleTime(300 * time.Second) // コネクションがアイドル状態で保持される時間

	boil.SetDB(db)
	return db, nil
}
