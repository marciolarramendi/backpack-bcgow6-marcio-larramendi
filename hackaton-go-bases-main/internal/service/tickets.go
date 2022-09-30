package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	ticketFound := Ticket{}
	ticketBool := false
	for i := 0; i < len(b.Tickets); i++ {
		if i+1 == id {
			ticketFound = b.Tickets[i]
			ticketBool = true
			break
			//fmt.Println(b.Tickets[i])
		}

		//fmt.Printf("%d - %v", i+1, b.Tickets[i])
	}
	if !ticketBool {
		return ticketFound, errors.New("ID not found")
	}
	//fmt.Println(ticketBool)
	return ticketFound, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
