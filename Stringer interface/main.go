package main

import "fmt"

type DistanceMeasureUnit int

const (
	Kilometer DistanceMeasureUnit = 0
	Mile      DistanceMeasureUnit = 1
)

type Distant struct {
	number float64
	unit   DistanceMeasureUnit
}

func (d DistanceMeasureUnit) String() string {
	units := []string{"km", "mi"}
	return units[d]
}

func (d Distant) String() string {
	return fmt.Sprintf("%v %s", d.number, d.unit)
}

func main() {
	d := Distant{79.6, Mile}
	fmt.Println(d)
}
