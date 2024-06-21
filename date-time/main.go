package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)

	t = time.Now()

	fmt.Println(t)

	fmt.Println(t.In(time.UTC))

	loc, _ := time.LoadLocation("Asia/Jakarta")
	fmt.Println(t.In(loc))

	timeString := "2024-07-20 11:00:00"

	// convert string to time
	t, _ = time.Parse("2006-01-02 15:04:05", timeString)
	fmt.Println(t.Year(), t.Month(), int(t.Month()), t.Day())

	// convert time to string with format
	fmt.Println(t.Format("02 Jan 2006 15:04"))

	currentTime := time.Now()
	// after one hour
	oneHourLater := currentTime.Add(time.Hour)
	fmt.Println(oneHourLater, oneHourLater.Hour())

	dateNow := time.Date(2024, time.July, 20, 11, 0, 0, 0, time.Now().Location())
	fmt.Println(dateNow)
}
