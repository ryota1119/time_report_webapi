package presenter

import (
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/usecase/dto"
)

type ProjectResponse struct {
	ID         entities.ProjectID         `json:"id"`
	CustomerID entities.CustomerID        `json:"customerID"`
	Name       entities.ProjectName       `json:"name"`
	UnitPrice  *entities.ProjectUnitPrice `json:"unitPrice"`
	StartDate  *string                    `json:"startDate"`
	EndDate    *string                    `json:"endDate"`
}

type ProjectCreateResponse ProjectResponse

func NewProjectCreateResponse(project *entities.Project) ProjectCreateResponse {
	return ProjectCreateResponse{
		ID:         project.ID,
		CustomerID: project.CustomerID,
		Name:       project.Name,
		UnitPrice:  project.UnitPrice,
		StartDate:  project.Period.Start.StringOrNil(),
		EndDate:    project.Period.End.StringOrNil(),
	}
}

type ProjectListResponse struct {
	ID           entities.ProjectID         `json:"id"`
	CustomerName entities.CustomerName      `json:"customerName"`
	Name         entities.ProjectName       `json:"name"`
	UnitPrice    *entities.ProjectUnitPrice `json:"unitPrice"`
	StartDate    *string                    `json:"startDate"`
	EndDate      *string                    `json:"endDate"`
}

func NewProjectListResponse(projects []*dto.ProjectWithCustomer) []ProjectListResponse {
	var output []ProjectListResponse
	for _, project := range projects {
		output = append(output, ProjectListResponse{
			ID:           project.Project.ID,
			CustomerName: project.Customer.Name,
			Name:         project.Project.Name,
			UnitPrice:    project.Project.UnitPrice,
			StartDate:    project.Project.Period.Start.StringOrNil(),
			EndDate:      project.Project.Period.End.StringOrNil(),
		})
	}
	return output
}

type ProjectGetResponse struct {
	ID           entities.ProjectID         `json:"id"`
	CustomerName entities.CustomerName      `json:"customerName"`
	Name         entities.ProjectName       `json:"name"`
	UnitPrice    *entities.ProjectUnitPrice `json:"unitPrice"`
	StartDate    *string                    `json:"startDate"`
	EndDate      *string                    `json:"endDate"`
}

func NewProjectGetResponse(project *dto.ProjectWithCustomer) ProjectGetResponse {
	return ProjectGetResponse{
		ID:           project.Project.ID,
		CustomerName: project.Customer.Name,
		Name:         project.Project.Name,
		UnitPrice:    project.Project.UnitPrice,
		StartDate:    project.Project.Period.Start.StringOrNil(),
		EndDate:      project.Project.Period.End.StringOrNil(),
	}
}

type ProjectUpdateResponse ProjectResponse

func NewProjectUpdateResponse(project *entities.Project) ProjectUpdateResponse {
	return ProjectUpdateResponse{
		ID:         project.ID,
		CustomerID: project.CustomerID,
		Name:       project.Name,
		UnitPrice:  project.UnitPrice,
		StartDate:  project.Period.Start.StringOrNil(),
		EndDate:    project.Period.End.StringOrNil(),
	}
}
