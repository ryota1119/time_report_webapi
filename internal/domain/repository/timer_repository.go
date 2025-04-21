package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// TimerRepository TimeRepositoryのインターフェースを定義
type TimerRepository interface {
	// Create はタイマー情報を作成する
	Create(ctx context.Context, tx *sql.Tx, timer *entities.Timer) (*entities.TimerID, error)
	// Find はタイマー情報を取得する
	Find(ctx context.Context, tx *sql.Tx, timerID *entities.TimerID) (*entities.Timer, error)
	// Update はタイマー情報を更新する
	Update(ctx context.Context, tx *sql.Tx, timer *entities.Timer) (*entities.TimerID, error)
	// ExistsRunningTimer はユーザーの稼働中のタイマーがあるかどうかチェックする
	ExistsRunningTimer(ctx context.Context, tx *sql.Tx) (bool, error)
}
