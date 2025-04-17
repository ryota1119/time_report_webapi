package entities

import (
	"time"
)

// ProjectID ユーザーID
type ProjectID uint

func (i ProjectID) Uint() uint {
	return uint(i)
}

type ProjectName string

type ProjectUnitPrice int64

// Project 顧客情報
type Project struct {
	ID         ProjectID
	CustomerID CustomerID
	Name       ProjectName
	UnitPrice  *ProjectUnitPrice
	StartDate  *time.Time
	EndDate    *time.Time
}

func NewProject(customerID uint, name string, unitPrice *int64, startDate *time.Time, endDate *time.Time) *Project {
	return &Project{
		CustomerID: CustomerID(customerID),
		Name:       ProjectName(name),
		UnitPrice:  ProjectUnitPriceOrNil(unitPrice),
		StartDate:  startDate,
		EndDate:    endDate,
	}
}

func ProjectUnitPriceOrNil(v *int64) *ProjectUnitPrice {
	if v == nil {
		return nil
	}
	c := ProjectUnitPrice(*v)
	return &c
}

func (p ProjectName) String() string {
	return string(p)
}

func (p ProjectUnitPrice) Int64() int64 {
	return int64(p)
}
