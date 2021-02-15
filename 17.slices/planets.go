package main

import "fmt"

func main() {
  planets := [...]string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  terrestrial := planets[0:4] // slice
  gasGiants := planets[4:6]
  iceGiants := planets[6:8]

  fmt.Println(terrestrial, gasGiants, iceGiants)


  // slices *maintain reference* to the underlying array
  iceGiants[1] = "Poseidon"

  fmt.Println(planets)
}
