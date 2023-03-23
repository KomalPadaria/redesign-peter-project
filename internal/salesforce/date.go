package salesforce

import (
	"fmt"
	"time"
)

func ParseStartDateEndDate(startDate, endDate string) (time.Time, time.Time, error) {
	startD, err := parseDate(startDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parsing SBQQStartDateC %v", err)
	}

	endD, err := parseDate(endDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("parsing SBQQEndDateC %v", err)
	}

	return startD, endD, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
