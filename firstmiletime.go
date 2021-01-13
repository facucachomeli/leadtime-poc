package leadtime

import "time"

//FirstMileTime representa el tiempo desde que se colecta un envio hasta que llega al facility
type FirstMileTime struct {
	Speed  time.Duration
	Cutoff time.Time
}

//GetName devulve el nombre del componente
func (fmt FirstMileTime) GetName() string {
	return "FirstMileTime"
}

//GetTimeFrom devuelve la estimacion del componente
func (fmt FirstMileTime) GetTimeFrom(start time.Time, dimensions Dimensions) (time.Time, time.Time) {
	return time.Now(), time.Now()
}
