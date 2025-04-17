package presenter

import (
	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type BudgetResponse struct {
	ID         uint    `json:"id"`
	ProjectID  uint    `json:"projectID"`
	Amount     int64   `json:"amount"`
	Memo       *string `json:"memo"`
	StartMonth string  `json:"startMonth"`
	EndMonth   string  `json:"endMonth"`
}

type BudgetCreateResponse BudgetResponse

func NewBudgetCreateResponse(budget *entities.Budget) BudgetCreateResponse {
	var memo *string
	if budget.BudgetMemo != nil {
		m := budget.BudgetMemo.String()
		memo = &m
	}
	return BudgetCreateResponse{
		ID:         budget.ID.Uint(),
		ProjectID:  budget.ProjectID.Uint(),
		Amount:     budget.BudgetAmount.Int(),
		Memo:       memo,
		StartMonth: budget.BudgetPeriod.Start.String(),
		EndMonth:   budget.BudgetPeriod.End.String(),
	}
}

type BudgetListResponse []BudgetResponse

func NewBudgetListResponse(budgets []entities.Budget) BudgetListResponse {
	var output BudgetListResponse
	for _, budget := range budgets {
		var memo *string
		if budget.BudgetMemo != nil {
			m := budget.BudgetMemo.String()
			memo = &m
		}
		output = append(output, BudgetResponse{
			ID:         budget.ID.Uint(),
			ProjectID:  budget.ProjectID.Uint(),
			Amount:     budget.BudgetAmount.Int(),
			Memo:       memo,
			StartMonth: budget.BudgetPeriod.Start.String(),
			EndMonth:   budget.BudgetPeriod.End.String(),
		})
	}
	return output
}

type BudgetGetResponse struct {
	ID               uint    `json:"id"`
	ProjectName      string  `json:"projectName"`
	ProjectUnitPrice *int64  `json:"projectUnitPrice"`
	Amount           int64   `json:"amount"`
	Memo             *string `json:"memo"`
	StartMonth       string  `json:"startMonth"`
	EndMonth         string  `json:"endMonth"`
}

func NewBudgetGetResponse(budgetWithProject *entities.BudgetWithProject) BudgetGetResponse {
	var memo *string
	if budgetWithProject.BudgetMemo != nil {
		m := budgetWithProject.BudgetMemo.String()
		memo = &m
	}
	var unitPrice *int64
	if budgetWithProject.ProjectUnitPrice != nil {
		p := budgetWithProject.ProjectUnitPrice.Int64()
		unitPrice = &p
	}
	return BudgetGetResponse{
		ID:               budgetWithProject.BudgetID.Uint(),
		ProjectName:      budgetWithProject.ProjectName.String(),
		ProjectUnitPrice: unitPrice,
		Amount:           budgetWithProject.BudgetAmount.Int(),
		Memo:             memo,
		StartMonth:       budgetWithProject.BudgetPeriod.Start.String(),
		EndMonth:         budgetWithProject.BudgetPeriod.End.String(),
	}
}

type BudgetUpdateResponse BudgetResponse

func NewBudgetUpdateResponse(budget *entities.Budget) BudgetUpdateResponse {
	var memo *string
	if budget.BudgetMemo != nil {
		m := budget.BudgetMemo.String()
		memo = &m
	}
	return BudgetUpdateResponse{
		ID:         budget.ID.Uint(),
		ProjectID:  budget.ProjectID.Uint(),
		Amount:     budget.BudgetAmount.Int(),
		Memo:       memo,
		StartMonth: budget.BudgetPeriod.Start.String(),
		EndMonth:   budget.BudgetPeriod.End.String(),
	}
}
