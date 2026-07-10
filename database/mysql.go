package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"user-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg config.Config, log *slog.Logger) (*sql.DB, error) {
	host := cfg.DBHost
	port := cfg.DBPort
	user := cfg.DBUser
	password := cfg.DBPassword
	dbName := cfg.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	log.Info("Connected to SQl", "host", cfg.DBHost, "database", cfg.DBName)
	if err != nil {
		log.Error("failed to Connect to sql", "host", cfg.DBHost, "database", cfg.DBName, "error", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
