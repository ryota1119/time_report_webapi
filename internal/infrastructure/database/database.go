package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/infrastructure/logger"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// NewDB　はデータベースの接続を初期化する
func NewDB() error {
	const (
		maxRetries = 10
		retryDelay = 3
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var newDB *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		newDB, err = sql.Open("mysql", dsn)
		if err != nil {
			logger.Warnf("[Retry %d/%d] ドライバ初期化エラー: %v", i, maxRetries, err)
			time.Sleep(time.Duration(retryDelay) * time.Second)
			continue
		}

		err = newDB.Ping()
		if err == nil {
			db = newDB
			logger.Logger.Info("✅ Connected to database")
			return nil
		}

		logger.Warnf("[Retry %d/%d] 接続失敗: %v", i, maxRetries, err)
		newDB.Close()
		time.Sleep(time.Duration(retryDelay) * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := newDB.Ping(); err != nil {
		newDB.Close()
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	return fmt.Errorf("DB接続に失敗しました（リトライ上限に到達）: %w", err)
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
