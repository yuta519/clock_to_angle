package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yuta519/clock_to_angle/go/entitiy"
)

func main() {
	currentTime := time.Now()

	fmt.Printf(
		"A current time is `%s:%s`\n",
		strconv.Itoa(currentTime.Hour()),
		strconv.Itoa(currentTime.Minute()),
	)

	fmt.Printf(
		"An angle between hands is %s°\n",
		strconv.FormatFloat(entitiy.FetchClockToAngle(currentTime.Hour(), currentTime.Minute()), 'f', 1, 64),
	)

}
