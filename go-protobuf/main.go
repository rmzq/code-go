package main

import (
	"fmt"
	pb "go-grpc/pb/greeting"
)

func main() {
	pbAddress := &pb.Address{Street: "123 Main Street", City: "New York"}
	fmt.Println(pbAddress)

	pbPerson := &pb.Person{Name: "John", Id: 123, Gender: pb.Gender_MALE, Address: pbAddress}
	fmt.Println(pbPerson)
}
