package transport

import "fmt"

type Bus struct {
	Name string
}

func (b *Bus) BoardPassengers() {
	fmt.Println("Passengers are boarding the bus:", b.Name)
}

func (b *Bus) UnboardPassengers() {
	fmt.Println("Passengers are unboarding the bus:", b.Name)
}