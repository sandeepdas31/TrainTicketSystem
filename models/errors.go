package models

import "fmt"

var (
	MissingFrom = ErrorResponse{ErrorMessage: "The field \"From\" is not populated"}

	MissingTo = ErrorResponse{ErrorMessage: "The field \"To\" is not populated"}

	InvalidFromLocation = ErrorResponse{ErrorMessage: "The From Location you have mentioned is either invalid or we do not support this location"}

	InvalidToLocation = ErrorResponse{ErrorMessage: "The To Location you have mentioned is either invalid or we do not support this location"}

	MissingFirstName = ErrorResponse{ErrorMessage: "The field \"First Name\" is not populated"}

	MissingLastName = ErrorResponse{ErrorMessage: "The field \"Last Name\" is not populated"}

	MissingEmailName = ErrorResponse{ErrorMessage: "The field \"Email\" is not populated"}

	InvalidEmail = ErrorResponse{ErrorMessage: "Invalid email address"}

	SeatUnavailable = ErrorResponse{ErrorMessage: "Seat Unavailable"}

	InvalidReceiptID = ErrorResponse{ErrorMessage: "Ticket not found with the receipt ID"}

	InvalidSection = ErrorResponse{ErrorMessage: "Section Not found"}

	TicketUnbooked = ErrorResponse{ErrorMessage: "No Tickets are Booked"}
)

type ErrorResponse struct {
	ErrorMessage string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("The API returned an error: %s ", e.ErrorMessage)
}
