package service

import (
	"errors"
	"path/filepath"
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
	indexPosition, err := b.GetIndexPosition(t.Id)
	if err != nil {
		b.Tickets = append(b.Tickets, t)
		filepath.Write(b.Tickets)
		return t, nil
	}
	ticketFound := b.Tickets[indexPosition]
	return ticketFound, errors.New("Ya hay un ticket con ese ID")
}

func (b *bookings) Read(id int) (Ticket, error) {
	indexPosition, err := b.GetIndexPosition(id)
	if err != nil {
		return Ticket{}, err
	}
	ticketFound := b.Tickets[indexPosition]
	return ticketFound, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	indexPosition, err := b.GetIndexPosition(id)
	if err != nil {
		return Ticket{}, err
	}
	ticketToUpd := b.Tickets[indexPosition]
	ticketToUpd.Names = t.Names
	ticketToUpd.Email = t.Email
	ticketToUpd.Destination = t.Destination
	ticketToUpd.Date = t.Date
	ticketToUpd.Price = t.Price
	return ticketToUpd, nil
}

func (b *bookings) Delete(id int) (int, error) {
	indexPosition, err := b.GetIndexPosition(id)
	if err != nil {
		return 0, err
	}
	b.Tickets = append(b.Tickets[:indexPosition], b.Tickets[indexPosition+1:]...)
	return id, nil
}

func (b *bookings) GetIndexPosition(id int) (int, error) {
	ticketBool := false
	var indexPosition int
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			ticketBool = true
			indexPosition = i
			break
		}
	}
	if !ticketBool {
		return 0, errors.New("ID not found")
	}
	return indexPosition, nil
}
