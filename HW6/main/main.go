/*Create a “Public transport” interface that has the methods 
“Receive passengers” and “Disembark passengers”, and implement 
it for the types “Bus”, “Train”, “Airplane”.Create the “Route” 
type, which contains a list of vehicles that are necessary to 
follow a given route. The “Route” type should have the methods 
“Add a vehicle to the route” and “Show a list of vehicles on 
the route”. Now your traveler (“Passenger”) has to go through 
this route and display his or her journey on the screen. Store 
files of different groups of objects in different packages. */

package main

import (
	"main/route"
	"main/transport"
)

func main() {
	bus := transport.NewBus("Bus 101")
	train := transport.NewTrain("Train A")
	plane := transport.NewPlane("Plane X")

	myRoute := &route.Route{}
	myRoute.AddVehicle(bus)
	myRoute.AddVehicle(train)
	myRoute.AddVehicle(plane)

	myRoute.ShowRoute()

	passenger := "John Doe"
	myRoute.Travel(passenger)
}