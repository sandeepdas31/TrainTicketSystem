package models

const (
	FromLocation = "london"
	ToLocation   = "france"

	EmailRegex  = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	TicketPrice = "$20"

	EmailString = " Email :"

	SectionA = "A"
	SectionB = "B"

	CancellationResult = "The ticket has been called and refund of " + TicketPrice + " will be provided"
)
