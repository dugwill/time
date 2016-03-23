// Copyright © 2016 Douglas A. Will.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// time shows examples of working with the time package in go

package main

import (
	"fmt"
	"math"
	"time"
)

// the following statements set up time constants
const ThirtyDays time.Duration = 720 * time.Hour // 720 hours = 30 days
const OneDay time.Duration = 24 * time.Hour      // 24 hours = 1 day
const OneMinute time.Duration = 1 * time.Minute  // 1 minute
const OneHour time.Duration = 1 * time.Hour      // 1 hour

//var MST *time.Location = &mstLoc

func main() {

	//Working with time

	t0 := time.Now()
	myBirthday := time.Date(1967, 1, 13, 13, 45, 00, 00, time.UTC)

	fmt.Printf("\nCurrent time is: %s\n\n", t0)

	// the formatting in time must reference Jan 2 15:04 2006 MST (01-02 3:04:05 pm 2006 -7GMT)
	// Various time formats
	fmt.Println(t0.Format("Mon 3:04:05 PM MST"))
	fmt.Println(t0.Format("3:04:05 PM"))
	fmt.Println(t0.Format("2006"))
	fmt.Println(t0.Format("MST"))
	fmt.Println(t0.Format("Jan 2 15:04:05 2006 MST"))
	fmt.Println(t0.Format("Mon"))
	fmt.Println(t0.Format("01-02 2006"))

	// Caluclate some times in the past
	monthAgo := t0.Add(-ThirtyDays)
	dayAgo := t0.Add(-OneDay)
	hourAgo := t0.Add(-OneHour)
	minuteAgo := t0.Add(-OneMinute)

	// Print the past times
	fmt.Println(monthAgo.Format("Thirty days ago was: Jan 2 15:04:05 2006 MST"))
	fmt.Println(dayAgo.Format("One day ago was: Jan 2 15:04:05 2006 MST"))
	fmt.Println(hourAgo.Format("One hour ago was: Jan 2 15:04:05 2006 MST"))
	fmt.Println(minuteAgo.Format("One minute ago was: Jan 2 15:04:05 2006 MST"))

	//My Birthday Calcuations
	fmt.Println("...")
	fmt.Println(myBirthday.Format("My Birthday is: Jan 2 15:04:05 2006 MST"))
	fmt.Println(myBirthday.Format("My Birthday was on a : Mon"))

	myAge := time.Since(myBirthday)

	fmt.Printf("My Age in hours %f\n", math.Abs(myAge.Hours()))

	// Sleep one second
	fmt.Println("Sleeping one second.")
	time.Sleep(time.Second)

	// Calculate and print the time it took this to run
	t1 := time.Now()
	fmt.Printf("\nThis took %v to run.\n", t1.Sub(t0))

}
