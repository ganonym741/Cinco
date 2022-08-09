package utilities

import "time"

func StringToTime(date string) (time.Time, error) {
	result, err := time.Parse(LayoutFormat, date)
	if err != nil {
		return time.Time{}, err
	}

	return result, nil
}
