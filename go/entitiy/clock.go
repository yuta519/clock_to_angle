package entitiy

func fetchShortHandAngle(currentHour int, currentMinute int) float64 {
	angleByOneHour := 30.0
	angleByOneMin := 0.5
	return float64(currentHour)*angleByOneHour + float64(currentMinute)*angleByOneMin
}

func fetchLongtHandAngle(currentMinute int) float64 {
	angleByOneMin := 6.0
	return float64(currentMinute) * angleByOneMin
}

func FetchClockToAngle(currentHour int, currentMinute int) float64 {
	shortHandAngle := fetchShortHandAngle(currentHour, currentMinute)
	longHandAngle := fetchLongtHandAngle(currentMinute)
	if shortHandAngle > longHandAngle {
		return shortHandAngle - longHandAngle
	}
	return longHandAngle - shortHandAngle
}
