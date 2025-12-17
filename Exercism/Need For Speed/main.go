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
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(trackDistance int) bool {
	if car.battery/car.batteryDrain*car.speed >= trackDistance {
		return true
	}
	return false
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Printf("%#v\n", car)
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())
	trackDistance := 100
	track := NewTrack(trackDistance)
	fmt.Printf("%#v\n", track)
	car.Drive()
	fmt.Printf("%#v\n", car)
	fmt.Printf("%#v\n", car.CanFinish(trackDistance))
	// => true
}
