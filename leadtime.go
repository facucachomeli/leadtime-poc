package leadtime

import (
	"time"
)

//LeadTime representa la estimacion del tiempo desde se concreta una compra hasta que se entrega un envio
type LeadTime struct {
	Start      time.Time
	Promise    time.Time
	PayBefore  time.Time
	Components []ComponentTime
}

//ComponentTime representa el inicio y fin del componente de una promesa de entrega
type ComponentTime struct {
	Type  string
	Start time.Time
	End   time.Time
}

//Dimensions determina las dimensiones para construir una promesa de entrega
type Dimensions struct {
	Facility     string
	Destination  string
	Service      int
	LogisticType string
}

type TimeComponent interface {
	GetTimeFrom(start time.Time, dimensions Dimensions) (end time.Time, payBefore time.Time)
	GetName() string
}

//Estimate crea una promesa de entrega
func Estimate(dimensions Dimensions) LeadTime {
	components := getComponentsBy(dimensions.LogisticType)
	return CalculatePromiseBy(time.Now(), components, dimensions)
}

func getComponentsBy(logisticType string) []TimeComponent {

	ht := HandlingTime{}
	ft := FacilityTime{}
	fmt := FirstMileTime{}
	st := ShippingTime{}

	switch logisticType {
	case "FBM":
		return []TimeComponent{ft, st}
	case "XD":
		return []TimeComponent{ht, fmt, ft, st}
	case "DS", "SS":
		return []TimeComponent{ht, st}
	}

	return []TimeComponent{}
}

//CalculatePromiseBy devulve la promesa por fecha de inicio, componentes y dimensiones
func CalculatePromiseBy(start time.Time, components []TimeComponent, dimensions Dimensions) LeadTime {
	leadtime := LeadTime{Start: start}
	for index, component := range components {
		end, paybefore := component.GetTimeFrom(start, dimensions)

		if isFirstComponent(index) {
			leadtime.PayBefore = paybefore
		}

		leadtime.Components = append(leadtime.Components, ComponentTime{component.GetName(), start, end})

		if isLastComponent(index, components) {
			leadtime.Promise = end
		} else {
			start = end
		}
	}

	return leadtime
}

func isFirstComponent(index int) bool {
	return index == 0
}

func isLastComponent(index int, components []TimeComponent) bool {
	return index == len(components)-1
}
