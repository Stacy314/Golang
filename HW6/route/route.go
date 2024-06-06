package route

import (
	"fmt"
	"main/transport"
)

type Route struct {
	vehicles []transport.PublicTransport
}

func (r *Route) AddVehicle(vehicle transport.PublicTransport) {
	r.vehicles = append(r.vehicles, vehicle)
}

func (r *Route) ShowRoute() {
	fmt.Println("Route includes the following vehicles:")
	for _, vehicle := range r.vehicles {
		fmt.Printf("%T\n", vehicle)
	}
}

func (r *Route) Travel(passenger string) {
	for _, vehicle := range r.vehicles {
		vehicle.AcceptPassengers(passenger)
		vehicle.DropPassengers(passenger)
	}
}