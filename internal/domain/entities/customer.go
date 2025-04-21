package entities

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// CustomerID 顧客ID
type CustomerID uint

// Uint は CustomerID を uint 型にキャストする
func (i CustomerID) Uint() uint {
	return uint(i)
}

// CustomerName 顧客名
type CustomerName string

// String は CustomerName を string 型にキャストする
func (n CustomerName) String() string {
	return string(n)
}

// CustomerUnitPrice 顧客単価
type CustomerUnitPrice int64

// Int64 は CustomerUnitPrice を int64 型にキャストする
func (u CustomerUnitPrice) Int64() int64 {
	return int64(u)
}

// NewCustomerUnitPrice は 指定された値から *entities.CustomerUnitPrice を返す
func NewCustomerUnitPrice(v *int64) *CustomerUnitPrice {
	if v == nil {
		return nil
	}
	c := CustomerUnitPrice(*v)
	return &c
}

// CustomerStartDate 顧客取引開始日
type CustomerStartDate time.Time

// String は CustomerStartDate をフォーマットに合わせて string 型にキャストする
func (d *CustomerStartDate) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format("2006-01-02")
}

// StringOrNil は CustomerStartDate をフォーマットに合わせて string 型にキャストする
func (d *CustomerStartDate) StringOrNil() *string {
	if d == nil {
		return nil
	}
	c := time.Time(*d).Format("2006-01-02")
	return &c
}

// Value は BudgetStartDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (d *CustomerStartDate) Value() driver.Value {
	if d == nil {
		return nil
	}
	return time.Time(*d)
}

// Equal は引数で受け取った entities.CustomerStartDate と比較する
func (d *CustomerStartDate) Equal(other CustomerStartDate) bool {
	if d == nil {
		return false
	}
	return time.Time(*d).Equal(time.Time(other))
}

// NewCustomerStartDate は string から entities.CustomerStartDate を返す
func NewCustomerStartDate(v *string) (*CustomerStartDate, error) {
	if v == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-2", *v)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse customer start date: %v", err))
	}
	startDate := CustomerStartDate(t)
	return &startDate, nil
}

// CustomerEndDate 顧客取引終了日
type CustomerEndDate time.Time

// String は CustomerEndDate をフォーマットに合わせて string 型にキャストする
func (d *CustomerEndDate) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format("2006-01-02")
}

// StringOrNil は CustomerStartDate をフォーマットに合わせて string 型にキャストする
func (d *CustomerEndDate) StringOrNil() *string {
	if d == nil {
		return nil
	}
	c := time.Time(*d).Format("2006-01-02")
	return &c
}

// Value は CustomerEndDate を Execメソッドに渡せるよう driver.Value 型にして返す
func (d *CustomerEndDate) Value() driver.Value {
	if d == nil {
		return nil
	}
	return time.Time(*d)
}

// Equal は引数で受け取った entities.CustomerStartDate と比較する
func (d *CustomerEndDate) Equal(other CustomerEndDate) bool {
	if d == nil {
		return false
	}
	return time.Time(*d).Equal(time.Time(other))
}

// NewCustomerEndDate は string から entities.CustomerEndDate を返す
func NewCustomerEndDate(v *string) (*CustomerEndDate, error) {
	if v == nil {
		return nil, nil
	}
	t, err := time.Parse("2006-01-2", *v)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse customer end date: %v", err))
	}
	endDate := CustomerEndDate(t)
	return &endDate, nil
}

// CustomerPeriod 顧客取引期間
type CustomerPeriod struct {
	Start *CustomerStartDate
	End   *CustomerEndDate
}

// NewCustomerPeriod は string 型から entities.BudgetPeriod を返す
func NewCustomerPeriod(start, end *string) (*CustomerPeriod, error) {
	var sDate *CustomerStartDate
	var eDate *CustomerEndDate
	var err error

	if start != nil {
		sDate, err = NewCustomerStartDate(start)
		if err != nil {
			return nil, err
		}
	}

	if end != nil {
		eDate, err = NewCustomerEndDate(end)
		if err != nil {
			return nil, err
		}
	}

	if sDate != nil && eDate != nil && time.Time(*sDate).After(time.Time(*eDate)) {
		return nil, errors.New("開始日は終了日より前でなければなりません")
	}

	return &CustomerPeriod{Start: sDate, End: eDate}, nil
}

// Customer 顧客情報
type Customer struct {
	ID        CustomerID
	Name      CustomerName
	UnitPrice *CustomerUnitPrice
	Period    CustomerPeriod
}

// NewCustomer は指定された顧客情報から entities.Customer を返す
func NewCustomer(name string, unitPrice *int64, startDate, endDate *string) (*Customer, error) {
	customerPeriod, err := NewCustomerPeriod(startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &Customer{
		Name:      CustomerName(name),
		UnitPrice: NewCustomerUnitPrice(unitPrice),
		Period:    *customerPeriod,
	}, nil
}
