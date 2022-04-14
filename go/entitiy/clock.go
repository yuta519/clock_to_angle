package entitiy

func FetchShortHandAngle(currentHour int, currentMinute int) float64 {
	return 0.1
}

func FetchLongtHandAngle(currentMinute int) float64 {
	angleByOneMin := 6.0
	return float64(currentMinute) * angleByOneMin
}

func FetchClockToAngle(currentHour int, currentMinute int) float64 {
	return 0.0
}
