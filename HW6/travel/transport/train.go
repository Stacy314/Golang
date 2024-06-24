package transport

import "fmt"

type Train struct {
	Name string
}

func (t *Train) BoardPassengers() {
	fmt.Println("Passengers are boarding the train:", t.Name)
}

func (t *Train) UnboardPassengers() {
	fmt.Println("Passengers are unboarding the train:", t.Name)
}