package commons

import (
	"fmt"
	"time"
)

func ConvertToDate(year, month, day int32) (*time.Time, error) {
	yearInt := int(year)
	monthInt := int(month)
	dayInt := int(day)

	if yearInt < 1 || yearInt > 9999 || monthInt < 1 || monthInt > 12 || dayInt < 1 || dayInt > 31 {
		return nil, fmt.Errorf("invalid birthdate")
	}

	time := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)

	return &time, nil
}
