package utilities

import "time"

func StringToTime(date string) (time.Time, error) {
	result, err := time.Parse(LayoutFormat, date)
	if err != nil {
		return time.Time{}, err
	}

	return result, nil
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
