package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// 7/13/2020 20:32:00
	d, _ := time.Parse("1/2/2006 15:04:05", date)
	return d
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	d, _ := time.Parse("January 2, 2006 15:04:05", date)

	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	d, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return d.Hour() >= 12 && d.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	//6/6/2005 10:30:00
	d, _ := time.Parse("1/2/2006 15:04:05", date)
	return d.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
