package entitiy

func fetchShortHandAngle(currentHour int, currentMinute int) float64 {
	return 0.1
}

func fetchLongtHandAngle(currentMinute int) float64 {
	angleByOneMin := 6.0
	return float64(currentMinute) * angleByOneMin
}
