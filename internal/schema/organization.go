package schema

import (
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type OrganizationID entities.OrganizationID

type OrganizationName entities.OrganizationName

type OrganizationCode entities.OrganizationCode

type Organization struct {
	ID               OrganizationID   `gorm:"primaryKey" json:"id"`
	OrganizationName OrganizationName `gorm:"not null" json:"organizationName"`
	OrganizationCode OrganizationCode `gorm:"unique;not null" json:"organizationCode"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	Users            []User           `gorm:"constraint:OnDelete:CASCADE" json:"users"`
	Customers        []Customer       `gorm:"constraint:OnDelete:CASCADE" json:"customers"`
	Projects         []Project        `gorm:"constraint:OnDelete:CASCADE" json:"projects"`
	Budgets          []Budget         `gorm:"constraint:OnDelete:CASCADE" json:"budgets"`
	Timer            []Timer          `gorm:"constraint:OnDelete:CASCADE" json:"timers"`
}
