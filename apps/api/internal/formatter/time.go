package formatter

import "time"

func DaysBetween(start, end time.Time) float64 {

	hours := end.Sub(start).Hours()

	if hours < 0 {
		return 0
	}

	return Round(hours/24, 1)
}
