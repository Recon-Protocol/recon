package formatter

const SompiPerKAS = 100000000

func SompiToKAS(v uint64) float64 {
	return float64(v) / SompiPerKAS
}

func SompiToBillions(v uint64) float64 {
	return SompiToKAS(v) / 1_000_000_000
}
