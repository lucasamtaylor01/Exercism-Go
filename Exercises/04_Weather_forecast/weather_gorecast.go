// Package weather has functions that are responsible for describing weather conditions.
package weather

// CurrentCondition stores information about the local meteorological conditions.
var CurrentCondition string 

// CurrentLocation stores information about geographical region.
var CurrentLocation string 


// Forecast updates and returns the weather climatic condition of a given geographical region.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
