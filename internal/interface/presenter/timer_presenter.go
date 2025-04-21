package presenter

import (
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

type TimerResponse struct {
	ID        uint       `json:"id"`
	ProjectID uint       `json:"projectID"`
	Title     string     `json:"title"`
	Memo      *string    `json:"memo,omitempty"`
	StartAt   time.Time  `json:"startAt"`
	EndAt     *time.Time `json:"endAt,omitempty"`
}

type StartTimerResponse TimerResponse

func NewStartTimerResponse(timer *entities.Timer) StartTimerResponse {
	var memo *string
	if timer.Memo != nil {
		c := timer.Memo.String()
		memo = &c
	}

	var endAt *time.Time
	if timer.EndAt != nil {
		c := timer.EndAt.Time()
		endAt = &c
	}
	return StartTimerResponse{
		ID:        timer.ID.Uint(),
		ProjectID: timer.ProjectID.Uint(),
		Title:     timer.Title.String(),
		Memo:      memo,
		StartAt:   timer.StartAt.Time(),
		EndAt:     endAt,
	}
}

type StopTimerResponse struct {
	ID        uint      `json:"id"`
	ProjectID uint      `json:"projectID"`
	Title     string    `json:"title"`
	Memo      *string   `json:"memo,omitempty"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
}

func NewStopTimerResponse(timer *entities.Timer) StopTimerResponse {
	var memo *string
	if timer.Memo != nil {
		c := timer.Memo.String()
		memo = &c
	}

	return StopTimerResponse{
		ID:        timer.ID.Uint(),
		ProjectID: timer.ProjectID.Uint(),
		Title:     timer.Title.String(),
		Memo:      memo,
		StartAt:   timer.StartAt.Time(),
		EndAt:     timer.EndAt.Time(),
	}
}
