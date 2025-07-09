package database

import (
	"awesomeProject/internal/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConfig.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
