# Application feature
This application books ticket from London to France and generates a receipt. 
The number of ticket that can be booked in a single go is 1. 
There are two sections sin the train A,B and has 2 tickets in each section which can be modified

# Languages used
Golang

# framework used
GRPC

# Stating the application
To run the code use the command  - go run main.go

The API's can be fired using two ways,
1. Using a GRPC client CLI such as grpcui. Install it. 
to run the server, use command - grpcui -plaintext localhost:8089, and the UI link will be generated. 
In the server all the API will be coming and we can hit from there 

2. Using applications such as Insomnia by importing the GRPC proto file and then hitting the API's

# API's present
1. Book ticket API
    Inputs:
        {
            JourneyDetails :  JourneyDetails{
                From : "London",
                To : "France
            }
            FirstName : "Sandeep
            LastName : "Das"
            Email : "ssandeepadas007@gmail.com"
        }

2. Cancel Ticket 
    Inputs:
        {
            TicketID: "sfdm54" // Should be a valid ticket ID
        }

3. CheckTicketReceipt
     Inputs:
        {
            TicketID: "sfdm54" // Should be a valid ticket ID
        }

4. CheckTicketAllotment - Checks the ticket alloted to the user in each section
    Inputs:
        {
            SectionID : "A"
        }

5. ModifySeat
    Inputs:
        {
             TicketID: "sfdm54" // Should be a valid ticket ID
        }

6. GetAllTicketDetails // Only for Admin    - No Inputs required
