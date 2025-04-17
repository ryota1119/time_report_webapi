package entities

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBudgetID_Uint(t *testing.T) {
	id := BudgetID(1)
	if id.Uint() != 1 {
		t.Errorf("expected id uint to be 1, got %v", id)
	}
}

func TestBudgetAmount_Int(t *testing.T) {
	amount := BudgetAmount(1)
	if amount.Int() != 1 {
		t.Errorf("expected amount to be 1, got %v", amount)
	}
}

func TestBudgetMemo_String(t *testing.T) {
	memo := BudgetMemo("test")
	if memo.String() != "test" {
		t.Errorf("expected memo to be test, got %v", memo)
	}
}

func TestBudgetStartDate_String(t *testing.T) {
	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date := BudgetStartDate(tm)

	want := "2024-03"
	got := date.String()

	assertEqual(t, want, got)
}

func TestBudgetStartDate_Value(t *testing.T) {
	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date := BudgetStartDate(tm)

	want := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	got := date.Value()

	assertEqual(t, want, got)
}

func TestNewBudgetStartDate(t *testing.T) {
	str := "2024-03"

	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	want := BudgetStartDate(tm)

	got, err := NewBudgetStartDate(str)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, want, *got)
}

func TestBudgetEndDate_String(t *testing.T) {
	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date := BudgetEndDate(tm)

	want := "2024-03"
	got := date.String()

	assertEqual(t, want, got)
}

func TestBudgetEndDate_Value(t *testing.T) {
	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	date := BudgetEndDate(tm)

	want := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	got := date.Value()

	assertEqual(t, want, got)
}

func TestNewBudgetEndDate(t *testing.T) {
	str := "2024-03"

	tm := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	firstOfNextMonth := tm.AddDate(0, 1, 0)
	endOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	want := BudgetEndDate(endOfMonth)

	got, err := NewBudgetEndDate(str)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, want, *got)
}

func TestNewBudgetPeriod(t *testing.T) {
	start := "2024-03"
	end := "2024-04"

	startDate, err := NewBudgetStartDate(start)
	if err != nil {
		t.Fatal(err)
	}
	endDate, err := NewBudgetEndDate(end)
	if err != nil {
		t.Fatal(err)
	}

	want := BudgetPeriod{
		Start: *startDate,
		End:   *endDate,
	}

	got, err := NewBudgetPeriod(start, end)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, *got, cmpopts.IgnoreUnexported(BudgetStartDate{}, BudgetEndDate{})) {
		t.Errorf("want %+v, got %+v", want, got)
	}

	if want.Start.String() != got.Start.String() {
		t.Errorf("want %+v, got %+v", want, got)
	}

	if want.End.String() != got.End.String() {
		t.Errorf("want %+v, got %+v", want, got)
	}
}

func TestNewBudgetPeriod_InvalidRange(t *testing.T) {
	start := "2024-05"
	end := "2024-04"

	_, err := NewBudgetPeriod(start, end)
	if err == nil {
		t.Error("expected error when start date is after end date, got nil")
	}
}

func TestNewBudget(t *testing.T) {
	projectID := uint(1)
	amount := int64(1000)
	memo := "test"
	startDate := "2024-03"
	endDate := "2024-04"

	budgetPeriod, err := NewBudgetPeriod(startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}

	want := Budget{
		ProjectID:    ProjectID(projectID),
		BudgetAmount: BudgetAmount(amount),
		BudgetMemo:   NewBudgetMemo(&memo),
		BudgetPeriod: *budgetPeriod,
	}

	got, err := NewBudget(projectID, amount, &memo, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, *got, cmpopts.IgnoreUnexported(BudgetStartDate{}, BudgetEndDate{})) {
		t.Errorf("want %+v, got %+v", want, got)
	}

	if want.BudgetPeriod.Start.String() != got.BudgetPeriod.Start.String() {
		t.Errorf("want %+v, got %+v", want, got)
	}

	if want.BudgetPeriod.End.String() != got.BudgetPeriod.End.String() {
		t.Errorf("want %+v, got %+v", want, got)
	}
}

func assertEqual(t *testing.T, want, got any) {
	t.Helper()
	if want != got {
		t.Errorf("assertEqual failed:\n  want = %v\n  got  = %v", want, got)
	}
}
