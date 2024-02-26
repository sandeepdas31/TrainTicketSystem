package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/thanhpk/randstr"
	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

var (
	//seatAvailable      = true
	seatsAllocatedInA  = 0
	seatsAllocatedInB  = 0
	totalSeatAllocated = 0
	currentSeat        = models.SectionA
	tempSeat           = models.SectionB
	seatBooking        = map[string]int{
		currentSeat: seatsAllocatedInA,
		tempSeat:    seatsAllocatedInB,
	}
	seatPerSection    = 2
	ticketDetails     = make(map[string]*TrainTicket.BookingResponse)
	sectionDetails    = make(map[string][]*TrainTicket.SeatAllotmentArray)
	cancelledSeatList = []string{}
)

func BookTicketService(ctx context.Context, req *TrainTicket.BookingRequest) (*TrainTicket.BookingResponse, error) {
	var seatNo string
	// Check ticket availability
	if !CheckAvailability() {
		return nil, models.SeatUnavailable
	}
	// Booking seat
	if seatBooking[currentSeat] < seatPerSection {
		seatNo = currentSeat + strconv.Itoa(seatBooking[currentSeat]+1)
		seatBooking[currentSeat] += 1
	} else {
		// assign seat fropm cancelled ticket list
		if len(cancelledSeatList) > 0 {
			seatNo = cancelledSeatList[0]
			cancelledSeatList = cancelledSeatList[1:]
		}
	}

	ticketID := randstr.Hex(5)
	user := req.FirstName + " " + req.LastName + models.EmailString + req.Email
	resp := TrainTicket.BookingResponse{
		JourneyDetails: &TrainTicket.JourneyDetails{
			From: req.JourneyDetails.From,
			To:   req.JourneyDetails.To,
		},
		User:      user,
		PricePaid: models.TicketPrice,
		TicketID:  ticketID,
		SeatNo:    seatNo,
	}
	// storing the ticket details
	ticketDetails[ticketID] = &resp
	sectionDetails[currentSeat] = append(sectionDetails[currentSeat], &TrainTicket.SeatAllotmentArray{SeatNo: seatNo, User: user})
	currentSeat, tempSeat = tempSeat, currentSeat
	fmt.Println("section Details ", sectionDetails)
	//Incrementing Available seats
	totalSeatAllocated++
	return &resp, nil

}

func CheckAvailability() bool {
	return totalSeatAllocated < (seatPerSection * 2)
}
