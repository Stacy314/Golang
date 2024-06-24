package transport

import "fmt"

type Plane struct {
	Name string
}

func (p *Plane) BoardPassengers() {
	fmt.Println("Passengers are boarding the plane:", p.Name)
}

func (p *Plane) UnboardPassengers() {
	fmt.Println("Passengers are unboarding the plane:", p.Name)
}