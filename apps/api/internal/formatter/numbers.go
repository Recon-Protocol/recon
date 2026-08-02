package formatter

import "math"

func Round(value float64, digits int) float64 {

	pow := math.Pow(10, float64(digits))

	return math.Round(value*pow) / pow
}
