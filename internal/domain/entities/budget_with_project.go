package entities

// BudgetWithProject プロジェクト情報を含んだ予算情報
type BudgetWithProject struct {
	BudgetID         BudgetID
	ProjectName      ProjectName
	ProjectUnitPrice *ProjectUnitPrice
	BudgetAmount     BudgetAmount
	BudgetMemo       *BudgetMemo
	BudgetPeriod     BudgetPeriod
}
