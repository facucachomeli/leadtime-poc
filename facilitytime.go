package leadtime

import "time"

//FirstMileTime representa el tiempo desde que llega al facility hasta que sale del facility (despacho)
type FacilityTime struct {
	WeavingTime time.Duration
	Cutoffs     []time.Time
}

//GetName devulve el nombre del componente
func (ft FacilityTime) GetName() string {
	return "FacilityTime"
}

//GetTimeFrom devuelve la estimacion del componente
func (ft FacilityTime) GetTimeFrom(start time.Time, dimensions Dimensions) (time.Time, time.Time) {
	//Weavingtime estaria en princpio hardcodeado a 3 horas
	//cutoff deberia obtenerse de la api de cutoff por las dimensiones(fecha, facility, destino)

	start = start.Add(ft.WeavingTime)

	cutoff := ft.getCutoff(start)

	return cutoff, cutoff.Add(-ft.WeavingTime)
}

func (ft FacilityTime) getCutoff(from time.Time) time.Time {
	for _, cutoff := range ft.Cutoffs {
		if cutoff.Before(from) {
			return cutoff
		}
	}

	return ft.Cutoffs[len(ft.Cutoffs)-1]
}
