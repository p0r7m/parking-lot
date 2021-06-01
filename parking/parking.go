package parking

import (
	"fmt"
	"time"
	"ud/payment"
)

type ParkingLot struct {
	capacity int64
	area     []ParkingArea
	hourlyCard payment.HourlyRateCard
	history  map[string]ParkingHistory
}

func CreateParkingLot() *ParkingLot {
	return &ParkingLot{capacity: -1, history: make(map[string]ParkingHistory)}
}

func (pl *ParkingLot) AddCapacity(capacity int64) {
	pl.capacity = capacity
	pl.area = make([]ParkingArea, capacity)
}

func (pl *ParkingLot) ParkVehicle(vehicleNum string) {
	if pl.capacity == -1  || len(pl.area) == 0 {
		return
	}

	areaId := pl.getFreeArea()
	if areaId == -1 {
		return
	}
	pl.area[areaId].InUse = true
	pl.area[areaId].StartDuration = time.Now().Unix()
	pl.area[areaId].Vehicle.VehicleNum = vehicleNum

	fmt.Println("Park Vehicle")
	pl.showParkingStats()
}

func  (pl *ParkingLot) getFreeArea() int {
	for areaId, eachArea := range pl.area {
		if eachArea.InUse == false {
			return areaId
		}
	}
	return -1
}

func (pl *ParkingLot) UnparkVehicle(vehicleNum string) int64 {
	if pl.capacity == -1  || len(pl.area) == 0 {
		return -1
	}
	for areaId, eachArea := range pl.area {
		if eachArea.Vehicle.VehicleNum == vehicleNum {
			eachArea.EndDuration = time.Now().Unix()
			pl.area[areaId] = eachArea
		}
	}
	fmt.Println("Unpark Vehicle")
	pl.showParkingStats()
	return -1
}

func (pl *ParkingLot) GetParkingAmount(vehicleNum string) int64 {
	if pl.capacity == -1  || len(pl.area) == 0 {
		return -1
	}
	pa := pl.getParkingArea(vehicleNum)

	if pa.InUse == false {
		return -1
	}
	cost := pa.CalculateCost(pl.hourlyCard)

	ph := ParkingHistory{}
	ph.duration = pa.EndDuration - pa.StartDuration
	ph.rateCard = "hourly"
	ph.vehicleNum = vehicleNum
	ph.amount = cost
	pl.history[vehicleNum] = ph

	fmt.Println("====Parking Cost=====")
	pl.DisplayParkingHistory(vehicleNum)
	fmt.Println("----------------------")
	return cost
}

func (pl *ParkingLot) DisplayParkingHistory(vehicleNum string) {
	history, ok := pl.history[vehicleNum]

	if !ok {
		return
	}
	fmt.Println(history.vehicleNum, history.duration, history.amount)
}

func (pl *ParkingLot) getParkingArea(vehicleNum string) ParkingArea {
	for _, eachArea := range pl.area {
		if eachArea.Vehicle.VehicleNum == vehicleNum {
			return eachArea
		}
	}
	return ParkingArea{}
}


func (pl *ParkingLot) showParkingStats() {
	for _, eachArea := range pl.area {
		if eachArea.StartDuration != 0 {
			fmt.Println("____________________")
			fmt.Println(eachArea.Vehicle.VehicleNum, eachArea.InUse, eachArea.StartDuration, eachArea.EndDuration)
			fmt.Println("____________________")
		}
	}
}






//Vehicle
//id
//plate_number
//kind (Two, Four Wheeler)
//created_at
//updated_at
//
//ParkingLot
//id
//capacity
//created_at
//updated_at
//
//ParkingArea
//id
//lot_id
//start_duration
//end_duration
//status [free, in_use]
//created_at
//updated_at
//
//VehicleParking
//id
//parking_area_id
//rate_card_type
//
//RateCard
//id
//kind (daily, hourly)
//start_duration
//end_duration
//cost
//created_at
//updated_at
//
//ParkingPayment
//id
//vehicle_id
//rate_card_id
//cost
//created_at
//
//ParkingHistory
//id
//vehicle_id
//area_id
//parking_payment_id
//start_duration
//end_duration
