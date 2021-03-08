package main

import "fmt"

func main() {
  planets := []string{
    "Mercury", "Venus", "Earth", "Mars",
    "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
  }

  terrestrial := planets[0:4:4] // capacity 4 (copies the backed array?)
  fmt.Println(terrestrial)

  terrestrial2 := planets[0:4] // capacety len(planets) (8)
  worlds := append(terrestrial2, "CERES") // overwrites planets[4]
  fmt.Println(worlds)
  fmt.Println(planets)

}
