package presenter

import (
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

type CustomerResponse struct {
	ID        entities.CustomerID
	Name      entities.CustomerName
	UnitPrice *entities.CustomerUnitPrice
	StartDate *string
	EndDate   *string
}

type CustomerCreateResponse CustomerResponse

func NewCustomerCreateResponse(customer *entities.Customer) CustomerCreateResponse {
	return CustomerCreateResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		UnitPrice: customer.UnitPrice,
		StartDate: customer.Period.Start.StringOrNil(),
		EndDate:   customer.Period.End.StringOrNil(),
	}
}

type CustomerListResponse []CustomerResponse

func NewCustomerListResponse(customers []entities.Customer) []CustomerResponse {
	var output CustomerListResponse
	for _, customer := range customers {
		output = append(output, CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			UnitPrice: customer.UnitPrice,
			StartDate: customer.Period.Start.StringOrNil(),
			EndDate:   customer.Period.End.StringOrNil(),
		})
	}
	return output
}

type CustomerGetResponse CustomerResponse

func NewCustomerGetResponse(customer *entities.Customer) CustomerGetResponse {
	return CustomerGetResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		UnitPrice: customer.UnitPrice,
		StartDate: customer.Period.Start.StringOrNil(),
		EndDate:   customer.Period.End.StringOrNil(),
	}
}

type CustomerUpdateResponse CustomerResponse

func NewCustomerUpdateResponse(customer *entities.Customer) CustomerUpdateResponse {
	return CustomerUpdateResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		UnitPrice: customer.UnitPrice,
		StartDate: customer.Period.Start.StringOrNil(),
		EndDate:   customer.Period.End.StringOrNil(),
	}
}
