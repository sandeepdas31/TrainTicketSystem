package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trainTicketBooking/models"
	trainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func TestValidate(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		request  *trainTicket.BookingRequest
		expected error
	}{
		{
			name: "ValidRequest",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "France",
				},
			},
			expected: nil,
		},
		{
			name: "MissingFrom",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "", // Missing From
					To:   "France",
				},
			},
			expected: models.MissingFrom,
		},
		{
			name: "MissingTo",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "", // Missing To
				},
			},
			expected: models.MissingTo,
		},
		{
			name: "InvalidFromLocation",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "CityC", // Invalid From Location
					To:   "France",
				},
			},
			expected: models.InvalidFromLocation,
		},
		{
			name: "InvalidToLocation",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "CityC", // Invalid To Location
				},
			},
			expected: models.InvalidToLocation,
		},
		{
			name: "MissingFirstName",
			request: &trainTicket.BookingRequest{
				FirstName: "", // Missing FirstName
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "France",
				},
			},
			expected: models.MissingFirstName,
		},
		{
			name: "MissingLastName",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "", // Missing LastName
				Email:     "john.doe@example.com",
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "France",
				},
			},
			expected: models.MissingLastName,
		},
		{
			name: "MissingEmail",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "", // Missing Email
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "France",
				},
			},
			expected: models.MissingEmailName,
		},
		{
			name: "InvalidEmail",
			request: &trainTicket.BookingRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "invalidemail", // Invalid Email
				JourneyDetails: &trainTicket.JourneyDetails{
					From: "London",
					To:   "France",
				},
			},
			expected: models.InvalidEmail,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.request)
			assert.Equal(t, tc.expected, err)
		})
	}
}

func TestValidateReq(t *testing.T) {
	testCases := []struct {
		name string
		req  map[string]interface{}
		err  error
	}{
		{
			name: "ValidRequest",
			req: map[string]interface{}{
				"key": "value",
			},
			err: nil,
		},
		{
			name: "EmptyRequest",
			req:  map[string]interface{}{},
			err:  fmt.Errorf("invalid request"),
		},
		{
			name: "InvalidJSON",

			req: map[string]interface{}{"": make(chan int)}, // Invalid JSON
			err: fmt.Errorf("grpc: error while marshaling: json: unsupported type: chan int"),
		},
	}

	for _, tc := range testCases {
		err := ValidateReq(tc.req)
		assert.Equal(t, err, tc.err)
	}
}
