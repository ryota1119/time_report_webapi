package schema

import (
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type BudgetID entities.BudgetID

type BudgetAmount entities.BudgetAmount

type BudgetMemo entities.BudgetMemo

type Budget struct {
	ID             BudgetID       `gorm:"primaryKey" json:"id"`
	OrganizationID OrganizationID `gorm:"index;not null" json:"organizationId"`
	ProjectID      ProjectID      `gorm:"index;not null" json:"projectId"`
	Amount         BudgetAmount   `gorm:"type:bigint;not null" json:"amount"`
	Memo           BudgetMemo     `gorm:"type:text" json:"memo"`
	StartDate      time.Time      `gorm:"type:date" json:"startDate"`
	EndDate        time.Time      `gorm:"type:date" json:"endDate"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}
