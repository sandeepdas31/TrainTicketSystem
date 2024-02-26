package service

import (
	"context"

	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func ModifySeatService(ctx context.Context, TicketID string) (*TrainTicket.BookingResponse, error) {
	ticketDetails, err := CheckTicketReceipt(TicketID)
	if err != nil {
		return nil, err
	}
	req := TrainTicket.BookingRequest{JourneyDetails: &TrainTicket.JourneyDetails{
		From: ticketDetails.JourneyDetails.From,
		To:   ticketDetails.JourneyDetails.To,
	},
		FirstName: string(ticketDetails.User[0]),
		LastName:  string(ticketDetails.User[1]),
		Email:     string(ticketDetails.User[3])}

	resp, err := BookTicketService(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.SeatNo == ticketDetails.SeatNo {
		return nil, models.SeatUnavailable
	}
	// cancel the old ticket
	_, err = CancelTicketService(ctx, TicketID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
