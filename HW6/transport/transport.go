package transport

import "fmt"

type PublicTransport interface {
	AcceptPassengers(passenger string)
	DropPassengers(passenger string)
}

type Bus struct {
	name string
}

func NewBus(name string) *Bus {
	return &Bus{name: name}
}

func (b *Bus) AcceptPassengers(passenger string) {
	fmt.Printf("Passenger %s boards the bus %s\n", passenger, b.name)
}

func (b *Bus) DropPassengers(passenger string) {
	fmt.Printf("Passenger %s gets off the bus %s\n", passenger, b.name)
}

type Train struct {
	name string
}

func NewTrain(name string) *Train {
	return &Train{name: name}
}

func (t *Train) AcceptPassengers(passenger string) {
	fmt.Printf("Passenger %s boards the train %s\n", passenger, t.name)
}

func (t *Train) DropPassengers(passenger string) {
	fmt.Printf("Passenger %s gets off the train %s\n", passenger, t.name)
}

type Plane struct {
	name string
}

func NewPlane(name string) *Plane {
	return &Plane{name: name}
}

func (p *Plane) AcceptPassengers(passenger string) {
	fmt.Printf("Passenger %s boards the plane %s\n", passenger, p.name)
}

func (p *Plane) DropPassengers(passenger string) {
	fmt.Printf("Passenger %s gets off the plane %s\n", passenger, p.name)
}