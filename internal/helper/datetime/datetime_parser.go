package datetime

import (
	"errors"
	"time"
)

// ParseDate は "YYYY-MM-DD" の形式の日付を time.Time に変換
func ParseDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil {
		return nil, nil
	}

	parsedTime, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

// ParseStartEndDate は "YYYY-MM-DD" の形式の日付を time.Time に変換
func ParseStartEndDate(startDateStr, endDateStr *string) (*time.Time, *time.Time, error) {
	var startDate *time.Time
	var endDate *time.Time

	if startDateStr != nil {
		parsedStartDate, err := time.Parse("2006-01-02", *startDateStr)
		if err != nil {
			return nil, nil, errors.New("invalid startDate format, expected YYYY-MM-DD")
		}
		startDate = &parsedStartDate
	}

	if endDateStr != nil {
		parsedEndDate, err := time.Parse("2006-01-02", *endDateStr)
		if err != nil {
			return nil, nil, errors.New("invalid endDate format, expected YYYY-MM-DD")
		}
		endDate = &parsedEndDate
	}

	return startDate, endDate, nil
}
