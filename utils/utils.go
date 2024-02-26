package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	mail "github.com/AfterShip/email-verifier"

	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func Validate(req *TrainTicket.BookingRequest) error {
	if req.JourneyDetails.From == "" {
		return models.MissingFrom
	}
	if req.JourneyDetails.To == "" {
		return models.MissingTo
	}
	if strings.ToLower(strings.TrimSpace(req.JourneyDetails.From)) != models.FromLocation {
		return models.InvalidFromLocation
	}
	if strings.ToLower(strings.TrimSpace(req.JourneyDetails.To)) != models.ToLocation {
		return models.InvalidToLocation
	}
	if req.FirstName == "" {
		return models.MissingFirstName
	}
	if req.LastName == "" {
		return models.MissingLastName
	}
	if req.Email == "" {
		return models.MissingEmailName
	}

	if !mail.IsAddressValid(req.Email) {
		return models.InvalidEmail
	}
	return nil
}

func ValidateReq(req interface{}) error {
	invalidReq, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("grpc: error while marshaling: %v", err.Error())
	}
	if string(invalidReq) == "{}" {
		return fmt.Errorf("invalid request")
	}
	return nil
}
