package main

import (
  "fmt"
  "strings"
)

type farenheit float64
type celcius float64

func (c celcius) farenheit() farenheit {
  return farenheit(c * 1.8 + 32)
}

func main() {

  line()
  fmt.Printf("| %-11v | %-11v |\n", "°C", "°F")
  line()
  for c := celcius(-40); c <= 100; c += 5 {
    print2cells(c, c.farenheit())
  }
  line()


}

func print2cells(c celcius, f farenheit) {
  fmt.Printf("| %-11v | %-11v|\n", c, f)
}

func line() {
  fmt.Println(strings.Repeat("=", 30))
}
