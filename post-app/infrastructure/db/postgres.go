package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/config"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PostgresDB struct {
}

func (p *PostgresDB) InitDB(dbConfig *config.DBConfig) (*sql.DB, error) {

	sslMode := "disable"
	if dbConfig.SSLMode {
		sslMode = "enable"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s connect_timeout=%s",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
			dbConfig.Host,
			strconv.Itoa(int(dbConfig.Port)),
			sslMode,
			strconv.Itoa(int(dbConfig.Timeout)),
		),
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
	db.SetMaxOpenConns(int(dbConfig.MaxConnections))                                   // 同時に使用できる最大コネクション数
	db.SetMaxIdleConns(int(dbConfig.MinConnections))                                   // コネクションプール内に保持する最大アイドルコネクション数
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnectionMaxLifetime) * time.Second) // コネクションを利用可能な最長時間
	db.SetConnMaxIdleTime(time.Duration(dbConfig.ConnectionMaxIdletime) * time.Second) // コネクションがアイドル状態で保持される時間

	boil.SetDB(db)
	return db, nil
}
