package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ TimerStopUsecase = (*timerStopUsecase)(nil)

// TimerStopUsecase TimerUsecaseのインターフェースを定義
type TimerStopUsecase interface {
	// Stop は予算を新規作成する
	Stop(ctx context.Context, input TimerStopUsecaseInput) (*entities.Timer, error)
}

// timerStopUsecase ユースケース
type timerStopUsecase struct {
	db        *sql.DB
	timerRepo repository.TimerRepository
}

func NewTimerStopUsecase(
	db *sql.DB,
	timerRepo repository.TimerRepository,
) TimerStopUsecase {
	return &timerStopUsecase{
		db:        db,
		timerRepo: timerRepo,
	}
}

// TimerStopUsecaseInput TimerStopUsecase Stopメソッド用input
type TimerStopUsecaseInput struct {
	TimerID uint
}

// Stop は予算を新規作成する
func (a *timerStopUsecase) Stop(ctx context.Context, input TimerStopUsecaseInput) (*entities.Timer, error) {
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

	timerID := entities.TimerID(input.TimerID)
	timer, err := a.timerRepo.Find(ctx, tx, &timerID)
	if err != nil {
		return nil, err
	}

	if timer.EndAt != nil {
		return nil, entities.ErrTimerAlreadyStopped
	}

	now := time.Now()
	timer.EndAt = entities.TimerEndAtOrNil(&now)
	_, err = a.timerRepo.Update(ctx, tx, timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}
