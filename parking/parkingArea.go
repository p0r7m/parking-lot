package parking

import (
	"ud/payment"
	vehicle2 "ud/vehicle"
)

type ParkingArea struct {
	Vehicle       vehicle2.Vehicle
	StartDuration int64
	EndDuration   int64
	InUse         bool
	CostType      string
}

func RefreshParkingArea() *ParkingArea {
	return &ParkingArea{CostType: "hourly", Vehicle: vehicle2.Vehicle{}}
}

func (pa *ParkingArea) CalculateCost(hourlyCard payment.HourlyRateCard) int64 {
	return payment.CalculateCost(pa.CostType, pa.StartDuration, pa.EndDuration, hourlyCard)
}
