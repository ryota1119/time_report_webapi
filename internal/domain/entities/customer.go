package entities

import (
	"time"
)

// CustomerID 顧客ID
type CustomerID uint

type CustomerName string

type CustomerUnitPrice int64

type CustomerStartDate time.Time

type CustomerEndDate time.Time

// Customer 顧客情報
type Customer struct {
	ID        CustomerID
	Name      CustomerName
	UnitPrice *CustomerUnitPrice
	StartDate *time.Time
	EndDate   *time.Time
}

func NewCustomer(name string, unitPrice *int64, startDate *time.Time, endDate *time.Time) *Customer {
	return &Customer{
		Name:      CustomerName(name),
		UnitPrice: CustomerUnitPriceOrNil(unitPrice),
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func CustomerUnitPriceOrNil(v *int64) *CustomerUnitPrice {
	if v == nil {
		return nil
	}
	c := CustomerUnitPrice(*v)
	return &c
}

func (c CustomerName) String() string {
	return string(c)
}

func (c CustomerUnitPrice) Int64() *int64 {
	return (*int64)(&c)
}
