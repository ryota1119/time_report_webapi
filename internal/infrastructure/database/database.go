package database

import (
	"database/sql"
	"fmt"
	"github.com/ryota1119/time_resport/internal/infrastructure/logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// NewDB　はデータベースの接続を初期化する
func NewDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	newDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := newDB.Ping(); err != nil {
		newDB.Close()
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	db = newDB
	logger.Logger.Info("Connected to database")
	return nil
}

// GetDB はgorm.DBを返す
func GetDB() *sql.DB {
	if db == nil {
		logger.Logger.Info("Database is not initialized. Call NewDB() first.")
	}
	return db
}

// CloseDB はデータベース接続を閉じる
func CloseDB() {
	if db != nil {
		db.Close()
		logger.Logger.Info("Closing database connection")
	}
}
