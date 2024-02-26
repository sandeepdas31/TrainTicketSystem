package service

import (
	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func CheckTicketSection(Section string) ([]*TrainTicket.SeatAllotmentArray, error) {

	if Section == models.SectionA || Section == models.SectionB {
		return sectionDetails[Section], nil
	}
	return nil, models.InvalidSection
}
