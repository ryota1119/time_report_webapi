package presenter

import (
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type ProjectResponse struct {
	ID         entities.ProjectID
	CustomerID entities.CustomerID
	Name       entities.ProjectName
	UnitPrice  *entities.ProjectUnitPrice
	StartDate  *time.Time
	EndDate    *time.Time
}

type ProjectCreateResponse ProjectResponse

func NewProjectCreateResponse(project *entities.Project) ProjectCreateResponse {
	return ProjectCreateResponse{
		ID:         project.ID,
		CustomerID: project.CustomerID,
		Name:       project.Name,
		UnitPrice:  project.UnitPrice,
		StartDate:  project.StartDate,
		EndDate:    project.EndDate,
	}
}

type ProjectListResponse []ProjectResponse

func NewProjectListResponse(projects []entities.Project) []ProjectResponse {
	var output ProjectListResponse
	for _, project := range projects {
		output = append(output, ProjectResponse{
			ID:         project.ID,
			CustomerID: project.CustomerID,
			Name:       project.Name,
			UnitPrice:  project.UnitPrice,
			StartDate:  project.StartDate,
			EndDate:    project.EndDate,
		})
	}
	return output
}

type ProjectGetResponse ProjectResponse

func NewProjectGetResponse(project *entities.Project) ProjectGetResponse {
	return ProjectGetResponse{
		ID:         project.ID,
		CustomerID: project.CustomerID,
		Name:       project.Name,
		UnitPrice:  project.UnitPrice,
		StartDate:  project.StartDate,
		EndDate:    project.EndDate,
	}
}

type ProjectUpdateResponse ProjectResponse

func NewProjectUpdateResponse(project *entities.Project) ProjectUpdateResponse {
	return ProjectUpdateResponse{
		ID:         project.ID,
		CustomerID: project.CustomerID,
		Name:       project.Name,
		UnitPrice:  project.UnitPrice,
		StartDate:  project.StartDate,
		EndDate:    project.EndDate,
	}
}
