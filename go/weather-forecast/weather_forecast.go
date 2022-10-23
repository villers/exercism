// Package weather represent the weather.
package weather

// CurrentCondition represent a string condition.
var CurrentCondition string

// CurrentLocation represent a string of current location.
var CurrentLocation string

// Forecast represent a function to return a sentence.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
