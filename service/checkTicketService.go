package service

import (
	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func CheckTicketReceipt(TicketID string) (*TrainTicket.BookingResponse, error) {
	data, found := ticketDetails[TicketID]
	if found {
		return data, nil
	}
	return nil, models.InvalidReceiptID
}

func GetAllTicketDetailsService() ([]*TrainTicket.BookingResponse, error) {

	if len(ticketDetails) == 0 {
		return nil, models.TicketUnbooked
	}
	var allTickets []*TrainTicket.BookingResponse
	for _, value := range ticketDetails {
		allTickets = append(allTickets, value)
	}
	return allTickets, nil
}
