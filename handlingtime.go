package leadtime

import "time"

//HandlingTime representa el tiempo desde que el seller recibe la compra hasta que lo despacha
type HandlingTime struct {
	Speed  time.Duration
	Cutoff time.Time
}

//GetName devulve el nombre del componente
func (ht HandlingTime) GetName() string {
	return "HandlingTime"
}

//GetTimeFrom devuelve la estimacion del componente
func (ht HandlingTime) GetTimeFrom(start time.Time, dimensions Dimensions) (time.Time, time.Time) {
	//handlingtime lo toma de la velocidad para el dia hora en la api de HT
	//Cutoff deberia obtenerse recorriendo la proyeccion desde el start time hasta la hora donde la velocidad cambia
	return start.Add(ht.Speed), ht.Cutoff
}
