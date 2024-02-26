package handlers

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/trainTicketBooking/service"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
	"github.com/trainTicketBooking/utils"
)

type TrainTicketServer struct {
	TrainTicket.UnimplementedTrainTicketServer
}

func (s *TrainTicketServer) BookTicket(ctx context.Context, req *TrainTicket.BookingRequest) (*TrainTicket.BookingResponse, error) {
	logr := logrus.WithContext(ctx)
	err := utils.ValidateReq(req)
	if err != nil {
		return nil, err
	}
	err = utils.Validate(req)
	if err != nil {
		logr.Error("Invalid request")
		return nil, err
	}
	ticketBooking, err := service.BookTicketService(ctx, req)
	if err != nil {
		return nil, err
	}

	return ticketBooking, nil
}

func (s *TrainTicketServer) CancelTicket(ctx context.Context, req *TrainTicket.TicketDetails) (*TrainTicket.CancellationResponse, error) {
	logr := logrus.WithContext(ctx)
	logr.Println("request ", req)
	err := utils.ValidateReq(req)
	if err != nil {
		return nil, err
	}
	cancellationResult, err := service.CancelTicketService(ctx, req.TicketID)
	if err != nil {
		return nil, err
	}
	return &TrainTicket.CancellationResponse{CancellationResult: cancellationResult.CancellationResult}, nil
}

func (s *TrainTicketServer) CheckTicketReceipt(ctx context.Context, req *TrainTicket.TicketDetails) (*TrainTicket.BookingResponse, error) {
	logr := logrus.WithContext(ctx)
	logr.Println("request ", req)
	err := utils.ValidateReq(req)
	if err != nil {
		return nil, err
	}
	receiptDetails, err := service.CheckTicketReceipt(req.TicketID)
	if err != nil {
		return nil, err
	}
	return receiptDetails, nil
}

func (s *TrainTicketServer) CheckTicketAllotment(ctx context.Context, req *TrainTicket.SectionDetails) (*TrainTicket.SeatAllotment, error) {
	logr := logrus.WithContext(ctx)
	logr.Println("request ", req)
	err := utils.ValidateReq(req)
	if err != nil {
		return nil, err
	}
	seatDetails, err := service.CheckTicketSection(req.SectionID)
	if err != nil {
		return nil, err
	}
	return &TrainTicket.SeatAllotment{
		SeatAllotmentArray: seatDetails,
	}, nil
}

func (s *TrainTicketServer) ModifySeat(ctx context.Context, req *TrainTicket.TicketDetails) (*TrainTicket.BookingResponse, error) {
	err := utils.ValidateReq(req)
	if err != nil {
		return nil, err
	}
	resp, err := service.ModifySeatService(ctx, req.TicketID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TrainTicketServer) GetAllTicketDetails(context.Context, *TrainTicket.NoParam) (*TrainTicket.AllBookings, error) {
	resp, err := service.GetAllTicketDetailsService()
	if err != nil {
		return nil, err
	}
	return &TrainTicket.AllBookings{
		BookingResponse: resp,
	}, nil
}
