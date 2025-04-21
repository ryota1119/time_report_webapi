package entities

import (
	"database/sql/driver"
	"time"
)

// ProjectID プロジェクトID
type ProjectID uint

// Uint は ProjectID を uint 型にキャストする
func (i ProjectID) Uint() uint {
	return uint(i)
}

// ProjectName プロジェクト名
type ProjectName string

// String は ProjectName を string 型にキャストする
func (p ProjectName) String() string {
	return string(p)
}

// ProjectUnitPrice プロジェクト単価
type ProjectUnitPrice int64

// Int64 は ProjectUnitPrice を int64 型にキャストする
func (p ProjectUnitPrice) Int64() int64 {
	return int64(p)
}

// NewProjectUnitPrice は *int64 を受け取り *ProjectUnitPrice を返す
func NewProjectUnitPrice(v *int64) *ProjectUnitPrice {
	if v == nil {
		return nil
	}
	c := ProjectUnitPrice(*v)
	return &c
}

// ProjectStartDate プロジェクト開始日
type ProjectStartDate time.Time

// String は ProjectStartDate をフォーマットに合わせて string 型にキャストする
func (d *ProjectStartDate) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format("2006-01-02")
}

// StringOrNil は ProjectStartDate をフォーマットに合わせて string 型にキャストする
func (d *ProjectStartDate) StringOrNil() *string {
	if d == nil {
		return nil
	}
	c := time.Time(*d).Format("2006-01-02")
	return &c
}

// Value は ProjectStartDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (d *ProjectStartDate) Value() driver.Value {
	if d == nil {
		return nil
	}
	return time.Time(*d)
}

// Equal は引数で受け取った entities.ProjectStartDate と比較する
func (d *ProjectStartDate) Equal(other ProjectStartDate) bool {
	if d == nil {
		return false
	}
	return time.Time(*d).Equal(time.Time(other))
}

// NewProjectStartDate は string から ProjectStartDate を返す
func NewProjectStartDate(v *string) (*ProjectStartDate, error) {
	if v == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-2", *v)
	if err != nil {
		return nil, ErrCouldNotParseProjectStartDate
	}
	c := ProjectStartDate(t)
	return &c, nil
}

// ProjectEndDate プロジェクト終了日
type ProjectEndDate time.Time

// String は ProjectEndDate をフォーマットに合わせて string 型にキャストする
func (d *ProjectEndDate) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format("2006-01-02")
}

// StringOrNil は ProjectEndDate をフォーマットに合わせて string 型にキャストする
func (d *ProjectEndDate) StringOrNil() *string {
	if d == nil {
		return nil
	}
	c := time.Time(*d).Format("2006-01-02")
	return &c
}

// Value は ProjectEndDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (d *ProjectEndDate) Value() driver.Value {
	if d == nil {
		return nil
	}
	return time.Time(*d)
}

// Equal は引数で受け取った entities.ProjectEndDate と比較する
func (d *ProjectEndDate) Equal(other ProjectEndDate) bool {
	if d == nil {
		return false
	}
	return time.Time(*d).Equal(time.Time(other))
}

// NewProjectEndDate は string から ProjectEndDate を返す
func NewProjectEndDate(v *string) (*ProjectEndDate, error) {
	if v == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-2", *v)
	if err != nil {
		return nil, ErrCouldNotParseProjectEndDate
	}
	c := ProjectEndDate(t)
	return &c, nil
}

// ProjectPeriod はプロジェクト期間
type ProjectPeriod struct {
	Start *ProjectStartDate
	End   *ProjectEndDate
}

func NewProjectPeriod(start, end *string) (*ProjectPeriod, error) {
	var sDate *ProjectStartDate
	var eDate *ProjectEndDate
	var err error

	if start != nil {
		sDate, err = NewProjectStartDate(start)
		if err != nil {
			return nil, err
		}
	}

	if end != nil {
		eDate, err = NewProjectEndDate(end)
		if err != nil {
			return nil, err
		}
	}

	if sDate != nil && eDate != nil && time.Time(*sDate).After(time.Time(*eDate)) {
		return nil, ErrStartDateMustBeBefore
	}

	return &ProjectPeriod{Start: sDate, End: eDate}, nil
}

// Project 顧客情報
type Project struct {
	ID         ProjectID
	CustomerID CustomerID
	Name       ProjectName
	UnitPrice  *ProjectUnitPrice
	Period     ProjectPeriod
}

func NewProject(customerID uint, name string, unitPrice *int64, startDate, endDate *string) (*Project, error) {
	projectPeriod, err := NewProjectPeriod(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &Project{
		CustomerID: CustomerID(customerID),
		Name:       ProjectName(name),
		UnitPrice:  NewProjectUnitPrice(unitPrice),
		Period:     *projectPeriod,
	}, nil
}
