package schema

import (
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"gorm.io/gorm"
)

type UserID entities.UserID

type UserName entities.UserName

type UserEmail entities.UserEmail

type Password entities.HashedPassword

type Role entities.Role

type User struct {
	ID             UserID         `gorm:"primaryKey" json:"id"`
	OrganizationID OrganizationID `gorm:"index;not null" json:"organizationId"`
	Name           UserName       `gorm:"size:255;not null" json:"name"`
	Email          UserEmail      `gorm:"size:255;unique;not null" json:"email"`
	Role           Role           `gorm:"type: enum('admin', 'member');not null" json:"role"`
	Password       Password       `gorm:"size:255;not null" json:"password"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Timer          []Timer        `gorm:"constraint:OnDelete:CASCADE" json:"times"`
}
