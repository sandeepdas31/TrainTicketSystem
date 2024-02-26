package main

import (
	"fmt"
	"log"
	"net"

	"github.com/trainTicketBooking/handlers"
	TrainTicket "github.com/trainTicketBooking/trainTicketProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// var from, to, Name, Email string
	// fmt.Println("Insert ticket details 1. From 2. To 3. Name 4. Email")
	// fmt.Scan(&from)
	// fmt.Scan(&to)
	// fmt.Scan(&Name)
	// fmt.Scan(&Email)

	// fmt.Println(from, to, Name, Email)
	z := byte(99)
	fmt.Println(string(z))
	// var x = make([][]string, 2)
	// for j := 0; j < 2; j++ {
	// 	for i := 0; i < 2; i++ {
	// 		x[j.string][i.string] = i

	// 	}
	// }

	//// Berth Logic
	// berth := 97
	// seat := 1
	// for i := 0; i < 20; i++ {
	// 	if berth > 108 {
	// 		fmt.Println("All seats are full")
	// 		return
	// 	}
	// 	seatAssigned := string(berth) + strconv.Itoa(seat)
	// 	seat++
	// 	if seat > 2 {
	// 		seat = 1
	// 		berth++
	// 	}

	// 	fmt.Println(seatAssigned)
	// }

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal("cannot create listener, ", err)
	}
	//seatCancelled := []string{}
	serviceRegister := grpc.NewServer()
	TrainTicket.RegisterTrainTicketServer(serviceRegister, &handlers.TrainTicketServer{})
	reflection.Register(serviceRegister)
	err = serviceRegister.Serve(lis)
	if err != nil {
		log.Fatal("cannot create listener, ", err)
	}
}
