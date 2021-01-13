package leadtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestFBMBuildCaso1(t *testing.T) {
	start := time.Date(2019, time.March, 28, 12, 0, 0, 0, time.UTC)
	components := []timeComponent{
		FacilityTime{
			weavingTime: time.Hour * 3,
			cutoff:      []time.Time{time.Date(2019, time.March, 28, 22, 0, 0, 0, time.UTC)},
		},
		ShippingTime{
			speed: 2,
		},
	}

	v, _ := json.Marshal(calculatePromise(start, components, Dimensions{}))
	fmt.Println(string(v))
}

func TestFBMBuildCaso2(t *testing.T) {
	start := time.Date(2019, time.March, 28, 19, 0, 0, 0, time.UTC)
	components := []timeComponent{
		FacilityTime{
			weavingTime: time.Hour * 3,
			cutoff:      []time.Time{time.Date(2019, time.March, 28, 22, 0, 0, 0, time.UTC)},
		},
		ShippingTime{
			speed: 2,
		},
	}

	v, _ := json.Marshal(calculatePromise(start, components, Dimensions{}))
	fmt.Println(string(v))
}

func TestFBMBuildCaso3(t *testing.T) {
	start := time.Date(2019, time.March, 28, 20, 0, 0, 0, time.UTC)
	components := []timeComponent{
		FacilityTime{
			weavingTime: time.Hour * 3,
			cutoff:      []time.Time{time.Date(2019, time.March, 29, 2, 0, 0, 0, time.UTC)},
		},
		ShippingTime{
			speed: 2,
		},
	}

	v, _ := json.Marshal(calculatePromise(start, components, Dimensions{}))
	fmt.Println(string(v))
}

func TestFBMBuildCaso4(t *testing.T) {
	start := time.Date(2019, time.March, 28, 12, 0, 0, 0, time.UTC)
	components := []timeComponent{
		FacilityTime{
			weavingTime: time.Hour * 3,
			cutoff:      []time.Time{time.Date(2019, time.March, 28, 22, 0, 0, 0, time.UTC)},
		},
		ShippingTime{
			speed: 2,
		},
	}

	v, _ := json.Marshal(calculatePromise(start, components, Dimensions{}))
	fmt.Println(string(v))
}
