// Package weather provides tools to get the weather.
package weather

// CurrentCondition represents a certain.
var CurrentCondition string

// CurrentLocation represents a certain.
var CurrentLocation string

// Forecast returns an integer.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
