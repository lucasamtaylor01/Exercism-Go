package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
    return car
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{
        distance: distance,
    }
    return track
}


// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if (car.battery - car.batteryDrain >= 0){
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }

    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numberCarCanDrive := car.battery / car.batteryDrain
    maximumDistance := car.speed*numberCarCanDrive
    
    if (maximumDistance >= track.distance){
        return true
    }else{
    	return false
    }
    
}

