package schema

import (
	"database/sql"
	"time"
)

type TimerID uint
type TimerTitle string
type TimerMemo string

type Timer struct {
	ID             TimerID        `gorm:"primaryKey" json:"id"`
	OrganizationID OrganizationID `gorm:"index;not null" json:"organizationId"`
	UserID         UserID         `gorm:"index;not null" json:"userId"`
	ProjectID      ProjectID      `gorm:"index" json:"projectId"`
	Title          TimerTitle     `json:"title"`
	Memo           TimerMemo      `json:"memo"`
	StartAt        time.Time      `gorm:"index;not null" json:"startAt"`
	EndAt          sql.NullTime   `gorm:"index" json:"endAt"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}
