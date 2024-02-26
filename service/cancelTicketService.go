package service

import (
	"context"

	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func CancelTicketService(ctx context.Context, TicketID string) (*TrainTicket.CancellationResponse, error) {

	ticket, err := CheckTicketReceipt(TicketID)
	if err != nil {
		return nil, err
	}
	// removing details from the store
	cancelledSeatList = append(cancelledSeatList, ticket.SeatNo)
	// delete from section details
	sectionDetails[string(ticket.SeatNo[0])] = removeTicket(ticket.SeatNo)
	// delete from ticket details
	delete(ticketDetails, TicketID)
	// incrasing the available seats
	totalSeatAllocated -= 1
	return &TrainTicket.CancellationResponse{
		CancellationResult: models.CancellationResult,
	}, nil
}

func removeTicket(seatNo string) []*TrainTicket.SeatAllotmentArray {
	val := sectionDetails[string(seatNo[0])]
	for i := 0; i < len(val); i++ {
		if val[i].SeatNo == seatNo {
			val = append(val[:i], val[i+1:]...)
		}
	}
	return val
}
