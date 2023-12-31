package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carsPerHour :=  float64(productionRate) * successRate/100
    return carsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
    return carsPerMinute
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsTenGroup := carsCount / 10
    carsIndividual := carsCount % 10
    cost :=  uint(carsTenGroup*95000 + carsIndividual*10000)
    
    return cost
}
