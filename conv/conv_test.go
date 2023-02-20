package conv

import (
    "reflect"
    "testing"
)

/*

    Mal for testfunksjoner
    Du skal skrive alle funksjonene basert på denne malen
    For alle konverteringsfunksjonene (tilsammen 6)
    kan du bruke malen som den er (du må selvsagt endre
    funksjonsnavn og testverdier)
*/

func TestFarhenheitToCelsius(ttesting.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 134, want: 56.67},
    }

    for , tc := range tests {
        got := FahrenheitToCelsius(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}

func TestCelsiusToFahrenheit(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 40, want: 104},
    }

    for , tc := range tests {
        got := CelsiusToFahrenheit(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}

func TestKelvinToFarhenheit(t *testing.T) {
    type test struct {
        input float64
        want  float64
    }

    tests := []test{
        {input: 400, want: 260.33},
    }

    for _, tc := range tests {
        got := KelvinToFahrenheit(tc.input)
        if !reflect.DeepEqual(tc.want, got) {
            t.Errorf("expected: %v, got: %v", tc.want, got)
        }
    }
}

// De andre testfunksjonene implementeres her
// ...