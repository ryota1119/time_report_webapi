package entities

import (
	"database/sql/driver"
	"time"
)

type TimerID uint

func (i TimerID) Uint() uint {
	return uint(i)
}

type TimerTitle string

func (t TimerTitle) String() string {
	return string(t)
}

type TimerMemo string

func (m TimerMemo) String() string {
	return string(m)
}

type TimerStartAt time.Time

func (t TimerStartAt) Time() time.Time {
	return time.Time(t)
}

func (t TimerStartAt) Value() driver.Value {
	return time.Time(t)
}

type TimerEndAt time.Time

func (t TimerEndAt) Time() time.Time {
	return time.Time(t)
}

func (t TimerEndAt) Value() driver.Value {
	return time.Time(t)
}

type Timer struct {
	ID        TimerID
	UserID    UserID
	ProjectID ProjectID
	Title     TimerTitle
	Memo      *TimerMemo
	StartAt   TimerStartAt
	EndAt     *TimerEndAt
}

func NewTimer(
	projectID uint,
	title string,
	memo *string,
	startAt time.Time,
	endAt *time.Time,
) *Timer {
	return &Timer{
		ProjectID: ProjectID(projectID),
		Title:     TimerTitle(title),
		Memo:      TimerMemoOrNil(memo),
		StartAt:   TimerStartAt(startAt),
		EndAt:     TimerEndAtOrNil(endAt),
	}
}

func TimerMemoOrNil(v *string) *TimerMemo {
	if v == nil {
		return nil
	}
	c := TimerMemo(*v)
	return &c
}

func TimerEndAtOrNil(v *time.Time) *TimerEndAt {
	if v == nil {
		return nil
	}
	c := TimerEndAt(*v)
	return &c
}
