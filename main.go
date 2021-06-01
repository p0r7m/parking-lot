package main

import (
	"fmt"
	"ud/parking"
)

func main() {

	parkingLot := parking.CreateParkingLot()
	parkingLot.AddCapacity(100)
	parkingLot.ParkVehicle("RZ-01-81736")
	parkingLot.UnparkVehicle("RZ-01-81736")
	amount := parkingLot.GetParkingAmount("RZ-01-81736")
	fmt.Print("Amount=")
	fmt.Println(amount)
	parkingLot.DisplayParkingHistory("RZ-01-81736")
}