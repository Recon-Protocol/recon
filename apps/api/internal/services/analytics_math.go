package services

import (
	"math"
	"time"
)

func ToPHS(hashrateGH float64) float64 {
	return math.Round((hashrateGH/1000.0)*100) / 100
}

func PercentMined(circulating, max uint64) float64 {
	if max == 0 {
		return 0
	}

	value := (float64(circulating) / float64(max)) * 100
	return math.Round(value*100) / 100
}

func RemainingSupply(circulating, max uint64) uint64 {
	if max <= circulating {
		return 0
	}

	return max - circulating
}

func DaysUntil(date time.Time) float64 {
	hours := time.Until(date).Hours()

	if hours < 0 {
		return 0
	}

	return math.Round((hours/24)*10) / 10
}
