package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadolibre/leadtime"
)

//Case test
type Case struct {
	caso       string
	start      time.Time
	components []leadtime.TimeComponent
}

func main() {
	casos := []Case{
		Case{
			caso:  "FF antes del 1er cutoff",
			start: time.Date(2019, time.March, 28, 12, 0, 0, 0, time.UTC),
			components: []leadtime.TimeComponent{
				leadtime.FacilityTime{
					WeavingTime: time.Duration(3) * time.Hour,
					Cutoffs:     []time.Time{time.Date(2019, time.March, 28, 22, 0, 0, 0, time.UTC)},
				},
				leadtime.ShippingTime{
					Speed: 2,
				},
			},
		},
		Case{
			caso:  "FF antes del 2do cutoff",
			start: time.Date(2019, time.March, 28, 20, 0, 0, 0, time.UTC),
			components: []leadtime.TimeComponent{
				leadtime.FacilityTime{
					WeavingTime: time.Duration(3) * time.Hour,
					Cutoffs: []time.Time{
						time.Date(2019, time.March, 28, 21, 0, 0, 0, time.UTC),
						time.Date(2019, time.March, 28, 23, 0, 0, 0, time.UTC),
					},
				},
				leadtime.ShippingTime{
					Speed: 2,
				},
			},
		},
		Case{
			caso:  "FF antes cutoff del otro dia antes del horario operativo del servicio",
			start: time.Date(2019, time.March, 28, 20, 0, 0, 0, time.UTC),
			components: []leadtime.TimeComponent{
				leadtime.FacilityTime{
					WeavingTime: time.Duration(3) * time.Hour,
					Cutoffs: []time.Time{
						time.Date(2019, time.March, 28, 21, 0, 0, 0, time.UTC),
						time.Date(2019, time.March, 29, 2, 0, 0, 0, time.UTC),
					},
				},
				leadtime.ShippingTime{
					Speed: 2,
				},
			},
		},
		Case{
			caso:  "FF antes cutoff del otro dia antes del horario operativo del servicio",
			start: time.Date(2019, time.March, 28, 20, 0, 0, 0, time.UTC),
			components: []leadtime.TimeComponent{
				leadtime.FacilityTime{
					WeavingTime: time.Duration(3) * time.Hour,
					Cutoffs: []time.Time{
						time.Date(2019, time.March, 28, 21, 0, 0, 0, time.UTC),
						time.Date(2019, time.March, 29, 2, 0, 0, 0, time.UTC),
					},
				},
				leadtime.ShippingTime{
					Speed: 2,
				},
			},
		},
	}
	for _, caso := range casos {
		v, _ := json.Marshal(leadtime.CalculatePromiseBy(caso.start, caso.components, leadtime.Dimensions{}))
		fmt.Println("//caso:", caso.caso)
		fmt.Println(string(v))
	}

}
