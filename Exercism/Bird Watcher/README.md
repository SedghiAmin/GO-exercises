# Bird Watcher (Exercism – Go)

This repository contains a Go solution for the **Bird Watcher** exercise from the **Exercism Go track**. The exercise focuses on working with slices, loops, and basic data manipulation in Go.

---

## 🧩 Problem Overview

The program analyzes a log of bird sightings, where each element in a slice represents the number of birds seen on a particular day.

Using this data, the program can:

* Calculate the total number of birds observed
* Calculate the number of birds seen in a specific week
* Fix an incorrectly recorded bird count log

---

## 📊 Data Model

The bird count log is represented as a slice of integers:

```go
[]int{2, 5, 0, 7, 4, 1, 3, ...}
```

Each value corresponds to the number of birds observed on a given day.

---

## ⚙️ Functions

### `TotalBirdCount`

```go
func TotalBirdCount(birdsPerDay []int) int
```

Calculates the total number of birds observed over all recorded days.

**Logic:**

* Iterates through the slice
* Accumulates the sum of all elements

---

### `BirdsInWeek`

```go
func BirdsInWeek(birdsPerDay []int, week int) int
```

Calculates the number of birds observed in a specific week.

**Details:**

* Each week consists of 7 days
* The `week` parameter is **1-based** (week 1 starts at index 0)
* Handles cases where the data does not contain a full week

---

### `FixBirdCountLog`

```go
func FixBirdCountLog(birdsPerDay []int) []int
```

Fixes the bird count log by correcting entries recorded every other day.

**Logic:**

* Increments the bird count by 1 for every even-indexed day
* Modifies the original slice (shared underlying array)

---

## ▶️ Example Usage

```go
birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

fmt.Println(TotalBirdCount(birdsPerDay))
fmt.Println(BirdsInWeek(birdsPerDay, 2))
fmt.Printf("%#v", FixBirdCountLog(birdsPerDay))
```

---

## 🧪 Learning Goals

This exercise helps reinforce:

* Working with slices in Go
* Looping with indices
* Boundary and range handling
* Understanding slice mutability and shared memory

---

## 📌 Notes

* `FixBirdCountLog` mutates the underlying slice data
* The implementation prioritizes clarity and correctness
* Designed to match Exercism’s idiomatic Go expectations

---

Happy coding! 🐦📈
