package conv

import "math"

func FahrenheitToCelsius(value float64) float64 {
    celsius := (value - 32) * (5.0 / 9.0)
    return math.Round(celsius100) / 100
}

func CelsiusToFahrenheit(value float64) float64 {
    fahrenheit := (value (9.0 / 5.0)) + 32
    return math.Round(fahrenheit100) / 100
}

func KelvinToCelsius(value float64) float64 {
    celsius := value - 273.15
    return math.Round(celsius100) / 100
}

func CelsiusToKelvin(value float64) float64 {
    kelvin := value + 273.15
    return math.Round(kelvin100) / 100
}

func KelvinToFahrenheit(value float64) float64 {
    fahrenheit := (value-273.15)(9.0/5.0) + 32
    return math.Round(fahrenheit100) / 100
}

func FahrenheitToKelvin(value float64) float64 {
    kelvin := (value-32)(5.0/9.0) + 273.15
    return math.Round(kelvin*100) / 100
}






package conv

import (
	"reflect"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
