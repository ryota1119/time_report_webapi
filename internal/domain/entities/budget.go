package entities

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// BudgetID ユーザーID
type BudgetID uint

// Uint は BudgetID を uint 型にキャストする
func (i BudgetID) Uint() uint {
	return uint(i)
}

// BudgetAmount 予算金額
type BudgetAmount int64

// Int は BudgetAmount を int64 型にキャストする
func (b BudgetAmount) Int() int64 {
	return int64(b)
}

// BudgetMemo 予算メモ
type BudgetMemo string

// String は BudgetMemo を string 型にキャストする
func (b BudgetMemo) String() string {
	return string(b)
}

// NewBudgetMemo は *string から entities.BudgetMemo を返す
func NewBudgetMemo(v *string) *BudgetMemo {
	if v == nil {
		return nil
	}
	c := BudgetMemo(*v)
	return &c
}

// BudgetStartDate は予算期間開始
type BudgetStartDate time.Time

// String は BudgetStartDate をフォーマットに合わせて string 型にキャストする
func (b BudgetStartDate) String() string {
	return time.Time(b).Format("2006-01")
}

// Value は BudgetStartDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (b BudgetStartDate) Value() driver.Value {
	return time.Time(b)
}

// NewBudgetStartDate は string から entities.BudgetStartDate を返す
func NewBudgetStartDate(v string) (*BudgetStartDate, error) {
	t, err := time.Parse("2006-01", v)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse budget start date: %v", err))
	}
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	startDate := BudgetStartDate(startOfMonth)
	return &startDate, nil
}

// BudgetEndDate は予算期間終了
type BudgetEndDate time.Time

// String は BudgetEndDate をフォーマットに合わせて string 型にキャストする
func (b BudgetEndDate) String() string {
	return time.Time(b).Format("2006-01")
}

// Value は BudgetEndDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (b BudgetEndDate) Value() driver.Value {
	return time.Time(b)
}

// NewBudgetEndDate は string から entities.BudgetEndDate を返す
func NewBudgetEndDate(v string) (*BudgetEndDate, error) {
	t, err := time.Parse("2006-01", v)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse budget end date: %v", err))
	}
	firstOfNextMonth := t.AddDate(0, 1, 0)
	endOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	endDate := BudgetEndDate(endOfMonth)
	return &endDate, nil
}

// BudgetPeriod 予算期間
type BudgetPeriod struct {
	Start BudgetStartDate
	End   BudgetEndDate
}

// NewBudgetPeriod は string 型から entities.BudgetPeriod を返す
func NewBudgetPeriod(start, end string) (*BudgetPeriod, error) {
	sDate, err := NewBudgetStartDate(start)
	if err != nil {
		return nil, err
	}
	eDate, err := NewBudgetEndDate(end)
	if err != nil {
		return nil, err
	}
	if time.Time(*sDate).After(time.Time(*eDate)) {
		return nil, errors.New("開始日は終了日より前でなければなりません")
	}
	return &BudgetPeriod{Start: *sDate, End: *eDate}, nil
}

// Budget 予算情報
type Budget struct {
	ID           BudgetID
	ProjectID    ProjectID
	BudgetAmount BudgetAmount
	BudgetMemo   *BudgetMemo
	BudgetPeriod BudgetPeriod
}

// NewBudget は指定された予算情報から entities.Budget を返す
func NewBudget(projectID uint, budgetAmount int64, budgetMemo *string, startDate, endDate string) (*Budget, error) {
	budgetPeriod, err := NewBudgetPeriod(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &Budget{
		ProjectID:    ProjectID(projectID),
		BudgetAmount: BudgetAmount(budgetAmount),
		BudgetMemo:   NewBudgetMemo(budgetMemo),
		BudgetPeriod: *budgetPeriod,
	}, nil
}
