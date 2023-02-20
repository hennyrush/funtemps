package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
  Celsius := (value - 32)*(5/9)
  return Celsius
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return 56.67
}

func CelciusToFahrenheit(value float64) float64 {
  Farhenheit := value*(9/5) + 32
  return Farhenheit
}

func KelvinToFarhenheit(value float64) float64 {
  Kelvin := (value-32)*(5/9) + 273.15
  return Kelvin
}

func KelvinToCelsius(value float64) float64 {
  Kelvin := (value -273.15)
  return Kelvin
}

func FarhenheitToKelvin(value float64) float64 {
 Farhenheit := (value-32)*(5/9)+273.15
  return Farhenheit
}

func CelsiusToKelvin(value float64) float64 {
  Celsius := (value+273.15)
  return Celsius
}

// De andre konverteringsfunksjonene implementere her
// ...
