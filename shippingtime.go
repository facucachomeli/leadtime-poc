package leadtime

import (
	"time"
)

//ShippingTime representa el tiempo desde que el correo obtiene el envio hasta que lo entrega
type ShippingTime struct {
	Speed int
}

//GetName devulve el nombre del componente
func (st ShippingTime) GetName() string {
	return "ShippingTime"
}

const zeroSpeed = 0
const serviceOperationalStartHour = 6
const deliveryHour = 12

//GetTimeFrom devuelve la estimacion del componente
func (st ShippingTime) GetTimeFrom(start time.Time, dimensions Dimensions) (time.Time, time.Time) {
	speed := st.getServiceSpeed(start)

	end := start.Add(time.Hour * time.Duration(24*speed))

	return time.Date(end.Year(), end.Month(), end.Day(), deliveryHour, 0, 0, 0, end.Location()), start
}

func (st ShippingTime) getServiceSpeed(start time.Time) int {
	if !sameDaySpeed(st.Speed) && beforeOperationalShippingTime(start.Hour()) {
		return st.Speed - 1
	}

	return st.Speed
}

func sameDaySpeed(speed int) bool {
	return speed != zeroSpeed
}

func beforeOperationalShippingTime(hour int) bool {
	return hour < serviceOperationalStartHour
}
