package main

import "fmt"

func main() {

  var planets [8]string // array decleration

  planets[0] = "Mercury"
  planets[1] = "Venus"
  planets[2] = "Earth"

  earth := planets[2]
  fmt.Println(earth) // "Earth"

}
