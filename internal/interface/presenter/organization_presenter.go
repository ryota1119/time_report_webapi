package presenter

import (
	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type OrganizationResponse struct {
	OrganizationName entities.OrganizationName `json:"organization_name" example:"My Organization"`
	OrganizationCode entities.OrganizationCode `json:"organization_code" example:"my_organization_code"`
}

type OrganizationRegisterResponse OrganizationResponse

func NewOrganizationRegisterResponse(organization *entities.Organization) OrganizationRegisterResponse {
	return OrganizationRegisterResponse{
		OrganizationName: organization.Name,
		OrganizationCode: organization.Code,
	}
}
