package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

//   receiver   method    result
func (k kelvin) celsius() celsius {
  return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
  return fahrenheit((k - 273.15) * 1.8 + 32)
}

func (c celsius) kelvin() kelvin {
  return kelvin(c + 273.15)
}
func (c celsius) fahrenheit() fahrenheit {
  return fahrenheit(c * 1.8 + 32)
}

func (f fahrenheit) kelvin() kelvin {
  return kelvin((f - 32 / 1.8) + 273.15)
}
func (f fahrenheit) celsius() celsius {
  return celsius((f - 32) / 1.8)
}

func main() {
  var k kelvin = 294.0
  c := k.celsius()
  fmt.Println("The answer is", c)

  var mc celsius = 127
  mk := mc.kelvin()
  fmt.Println("The surface temp is ", mk, " degrees Kelvin or ", mk.fahrenheit(), "F")
}
