package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// TimerRepository TimeRepositoryのインターフェースを定義
type TimerRepository interface {
	// Create はタイマー情報を作成する
	Create(ctx context.Context, tx *sql.Tx, timer *entities.Timer) (*entities.TimerID, error)
	// ExistsRunningTimer はユーザーの稼働中のタイマーがあるかどうかチェックする
	ExistsRunningTimer(ctx context.Context, tx *sql.Tx) (bool, error)
}
