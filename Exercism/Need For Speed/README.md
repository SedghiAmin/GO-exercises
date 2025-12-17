# Remote Controlled Car (Exercism – Go)

This repository contains a simple Go implementation for one of the **Exercism Go track exercises**. The goal of the exercise is to model a remote controlled car, simulate its movement, track battery usage, and determine whether it can finish a given track.

---

## 🧩 Problem Overview

The program models two main concepts:

* **Car**: A remote controlled car with limited battery capacity
* **Track**: A track with a fixed distance

The car can:

* Drive forward while consuming battery
* Report the total distance driven
* Report the remaining battery percentage
* Check whether it can finish a given track with its current battery

---

## 🏗️ Data Structures

### `Car`

```go
type Car struct {
    battery      int
    batteryDrain int
    speed        int
    distance     int
}
```

| Field          | Description                        |
| -------------- | ---------------------------------- |
| `battery`      | Current battery level (percentage) |
| `batteryDrain` | Battery consumed per drive         |
| `speed`        | Distance traveled per drive        |
| `distance`     | Total distance driven              |

---

### `Track`

```go
type Track struct {
    distance int
}
```

| Field      | Description        |
| ---------- | ------------------ |
| `distance` | Total track length |

---

## ⚙️ Functions & Methods

### `NewCar`

```go
func NewCar(speed, batteryDrain int) Car
```

Creates a new car with:

* 100% battery
* Given speed
* Given battery drain

---

### `NewTrack`

```go
func NewTrack(distance int) Track
```

Creates a new track with the specified distance.

---

### `Drive`

```go
func (car *Car) Drive()
```

Drives the car **once**:

* If there is enough battery, it:

    * Reduces the battery
    * Increases the driven distance
* If the battery is insufficient, the car does nothing

---

### `CanFinish`

```go
func (car Car) CanFinish(trackDistance int) bool
```

Checks whether the car can finish a track **without actually driving it**.

The logic is based on:

* Number of possible drives: `battery / batteryDrain`
* Total reachable distance: `possibleDrives * speed`

Returns:

* `true` if the car can finish the track
* `false` otherwise

---

### `DisplayDistance`

```go
func (car Car) DisplayDistance() string
```

Returns a formatted string representing the total distance driven.

Example:

```
Driven 40 meters
```

---

### `DisplayBattery`

```go
func (car Car) DisplayBattery() string
```

Returns a formatted string representing the remaining battery percentage.

Example:

```
Battery at 80%
```

---

## ▶️ Example Usage

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

fmt.Println(car.DisplayDistance())
fmt.Println(car.DisplayBattery())

trackDistance := 100
track := NewTrack(trackDistance)

car.Drive()
fmt.Println(car.CanFinish(track.distance))
```

---

## 🧪 Learning Goals

This exercise helps practice:

* Structs and methods in Go
* Pointer receivers vs value receivers
* Basic arithmetic reasoning
* Clean and readable API design

---

## 📌 Notes

* This implementation is intentionally simple and educational.
* The `CanFinish` method does **not** mutate the car state.
* Designed to align with Exercism's idiomatic Go style.

---

Happy coding! 🚗💨
