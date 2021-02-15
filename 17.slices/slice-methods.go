package main

import (
  "fmt"
  "sort"
  "strings"
)

func main() {
  planets := []string{
    "Mercury",
    "Venus",
    "Earth",
    "Mars",
    "Jupiter",
    "Saturn",
    "Uranus",
    "Neptune",
  }
  sort.Strings(planets)
  fmt.Println(planets)
  fmt.Println(strings.Join(planets, ", "))
  terraform(planets)
  fmt.Println(planets)
}

func terraform (p []string) {
  for i := range p {
    p[i] = "New " + p[i]
  }
}
