package database

import (
	"template/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQL(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", cfg.MySQLUri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
