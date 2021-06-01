package vehicle

type IVehicle interface {
	Park(int64)
	UnPark() IVehicle
}
