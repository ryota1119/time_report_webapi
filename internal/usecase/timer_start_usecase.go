package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport/internal/domain/repository"
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

var _ TimerStartUsecase = (*timerStartUsecase)(nil)

// TimerStartUsecase TimerUsecaseのインターフェースを定義
type TimerStartUsecase interface {
	// Start は予算を新規作成する
	Start(ctx context.Context, input TimerStartUsecaseInput) (*entities.Timer, error)
}

// timerStartUsecase ユースケース
type timerStartUsecase struct {
	db        *sql.DB
	timerRepo repository.TimerRepository
}

func NewTimerStartUsecase(
	db *sql.DB,
	timerRepo repository.TimerRepository,
) TimerStartUsecase {
	return &timerStartUsecase{
		db:        db,
		timerRepo: timerRepo,
	}
}

// TimerStartUsecaseInput TimerStartUsecase Startメソッド用input
type TimerStartUsecaseInput struct {
	ProjectID uint
	Title     string
	Memo      *string
}

// Start は予算を新規作成する
func (a *timerStartUsecase) Start(ctx context.Context, input TimerStartUsecaseInput) (*entities.Timer, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	isRunning, err := a.timerRepo.ExistsRunningTimer(ctx, tx)
	if err != nil {
		return nil, err
	}
	if isRunning {
		return nil, entities.ErrTimerAlreadyRunning
	}

	now := time.Now()
	// 予算を作成する
	timerRecord := entities.NewTimer(input.ProjectID, input.Title, input.Memo, now, nil)
	storedTimerID, err := a.timerRepo.Create(ctx, tx, timerRecord)
	if err != nil {
		return nil, err
	}
	timerRecord.ID = *storedTimerID

	return timerRecord, nil
}
