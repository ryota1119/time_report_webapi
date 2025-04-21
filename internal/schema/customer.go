package schema

import (
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"gorm.io/gorm"
)

type CustomerID entities.CustomerID

type CustomerName entities.CustomerName

type CustomerUnitPrice entities.CustomerUnitPrice

type Customer struct {
	ID             CustomerID        `gorm:"primaryKey" json:"id"`
	OrganizationID OrganizationID    `gorm:"index;not null" json:"organizationId"`
	Name           CustomerName      `gorm:"not null" json:"name"`
	UnitPrice      CustomerUnitPrice `gorm:"type:bigint" json:"unitPrice"`
	StartDate      time.Time         `gorm:"type:date" json:"startDate"`
	EndDate        time.Time         `gorm:"type:date" json:"endDate"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt    `gorm:"index" json:"deletedAt"`
	Projects       []Project         `gorm:"constraint:OnDelete:CASCADE" json:"projects"`
}
