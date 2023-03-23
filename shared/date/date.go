package date

import (
	"time"

	"github.com/pkg/errors"
)

func QuarterOf(month int) int {
	q := [12]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	return q[month-1]
}

func StartAndEndMonthOfQuarter(qtr int) (time.Month, time.Month, error) {
	switch qtr {
	case 1:
		return time.January, time.March, nil
	case 2:
		return time.April, time.June, nil
	case 3:
		return time.July, time.September, nil
	case 4:
		return time.October, time.December, nil
	default:
		return 0, 0, errors.New("invalid quarter")
	}
}
func NextQuarterStartDateAndEndDateByTime(t time.Time) (time.Time, time.Time, error) {
	qtr := QuarterOf(int(t.Month()))
	year := t.Year()
	if qtr == 4 {
		qtr = 1
		year += 1
	}

	startMonth, _, err := StartAndEndMonthOfQuarter(qtr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	t = time.Date(year, startMonth, 1, 0, 0, 0, 0, t.Location())

	return QuarterStartDateAndEndDateByTime(t)
}

func QuarterStartDateAndEndDateByTime(t time.Time) (time.Time, time.Time, error) {
	qtr := QuarterOf(int(t.Month()))
	startMonth, endMonth, err := StartAndEndMonthOfQuarter(qtr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	startDate := time.Date(t.Year(), startMonth, 1, 0, 0, 0, 0, t.Location())

	endDate := time.Date(t.Year(), endMonth+1, 0, 0, 0, 0, 0, t.Location())

	return startDate, endDate, nil
}

// Check whether "given" datetime is between "start" and "end" datetime
func InTimeSpan(start, end, given time.Time) bool {
	return given.After(start) && given.Before(end)
}
