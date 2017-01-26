package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go was launched on %s\n", t)

	now := time.Now()
	fmt.Printf("Current time is %s\n", now)
	fmt.Println("Day: ", t.Day())
	fmt.Println("Month: ", t.Month())
	fmt.Println("WeekDay: ", t.Weekday())

	// lets get tomorrow time
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v ,%v %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	//above is bit too much typing
	// instead we can use format variable
	longformat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(longformat))

	shortformat := "2/1/2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortformat))

}
