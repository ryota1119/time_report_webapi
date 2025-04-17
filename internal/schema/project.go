package schema

import (
	"database/sql"
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"gorm.io/gorm"
)

type ProjectID entities.ProjectID

type ProjectName entities.ProjectName

type ProjectUnitPrice sql.NullInt64

type Project struct {
	ID             ProjectID        `gorm:"primaryKey" json:"id"`
	OrganizationID OrganizationID   `gorm:"index;not null" json:"organizationId"`
	CustomerID     CustomerID       `gorm:"index;not null" json:"customerId"`
	Name           ProjectName      `gorm:"not null" json:"name"`
	UnitPrice      ProjectUnitPrice `gorm:"type:bigint" json:"unitPrice"`
	StartDate      time.Time        `gorm:"type:date" json:"startDate"`
	EndDate        time.Time        `gorm:"type:date" json:"endDate"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt   `gorm:"index" json:"deletedAt"`
	Timer          []Timer          `gorm:"constraint:OnDelete:CASCADE" json:"timers"`
	Budgets        []Budget         `gorm:"constraint:OnDelete:CASCADE" json:"budgets"`
}
