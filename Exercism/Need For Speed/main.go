package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	return Car{
		battery:      car.battery - car.batteryDrain,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
		distance:     car.speed + car.distance,
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Printf("%#v\n", car)
	distance := 800
	track := NewTrack(distance)
	fmt.Printf("%#v\n", track)
	car = Drive(car)
	fmt.Printf("%#v\n", car)
}
