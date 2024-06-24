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
	"travel/passenger"
	"travel/route"
	"travel/transport"
)

func main() {
	bus := &transport.Bus{Name: "City Bus"}
	train := &transport.Train{Name: "Express Train"}
	plane := &transport.Plane{Name: "Boeing 747"}

	travelRoute := &route.Route{}
	travelRoute.AddVehicle(bus)
	travelRoute.AddVehicle(train)
	travelRoute.AddVehicle(plane)

	traveler := &passenger.Passenger{Name: "John Doe"}
	traveler.Travel(travelRoute)
}