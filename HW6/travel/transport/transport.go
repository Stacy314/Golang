package transport

type PublicTransport interface {
	BoardPassengers()
	UnboardPassengers()
}