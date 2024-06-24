package passenger

import "travel/route"

type Passenger struct {
	Name string
}

func (p *Passenger) Travel(route *route.Route) {
	route.Travel()
}