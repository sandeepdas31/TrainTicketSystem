package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trainTicketBooking/models"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
)

func TestBookTicketService(t *testing.T) {
	// Mock request data
	req := &TrainTicket.BookingRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		JourneyDetails: &TrainTicket.JourneyDetails{
			From: "CityA",
			To:   "CityB",
		},
	}
	tests := []struct {
		name          string
		totalSeat     int
		expectedError error
	}{
		{
			name:          "Success",
			totalSeat:     1,
			expectedError: nil,
		},
		{
			name:          "SeatUnavailable",
			totalSeat:     seatPerSection * 3,
			expectedError: models.SeatUnavailable,
		},
		// Add more test cases for other scenarios
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// Call the function
			totalSeatAllocated = tc.totalSeat
			resp, err := BookTicketService(context.Background(), req)

			// Assertions
			if tc.expectedError != nil {
				assert.Nil(t, resp)
				//assert.Equal(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestCheckAvailability(t *testing.T) {

	testCases := []struct {
		name           string
		TotalSeat      int
		seatPersection int
		got            bool
	}{
		{
			name:           "Available",
			TotalSeat:      2,
			seatPersection: 4,
			got:            true,
		},
		{
			name:           "Not Available",
			TotalSeat:      4,
			seatPersection: 2,
			got:            false,
		},
	}
	// Success case: total seats allocated is less than twice the seat per section

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			totalSeatAllocated = tc.TotalSeat
			seatPerSection = tc.seatPersection
			val := CheckAvailability()
			assert.Equal(t, val, tc.got)
		})
	}
}
