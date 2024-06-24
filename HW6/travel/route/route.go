package route

import (
	"fmt"
	"travel/transport"
)

type Route struct {
	Vehicles []transport.PublicTransport
}

func (r *Route) AddVehicle(vehicle transport.PublicTransport) {
	r.Vehicles = append(r.Vehicles, vehicle)
}

func (r *Route) ShowVehicles() {
	for _, vehicle := range r.Vehicles {
		switch v := vehicle.(type) {
		case *transport.Bus:
			fmt.Println("Bus:", v.Name)
		case *transport.Train:
			fmt.Println("Train:", v.Name)
		case *transport.Plane:
			fmt.Println("Plane:", v.Name)
		}
	}
}

func (r *Route) Travel() {
	for _, vehicle := range r.Vehicles {
		vehicle.BoardPassengers()
		vehicle.UnboardPassengers()
	}
}