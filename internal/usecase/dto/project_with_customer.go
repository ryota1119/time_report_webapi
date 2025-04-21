package dto

import "github.com/ryota1119/time_resport_webapi/internal/domain/entities"

// ProjectWithCustomer entities.Project × entities.Customer の構造体
type ProjectWithCustomer struct {
	Project  *entities.Project
	Customer *entities.Customer
}
