package main

import "fmt"

func main() {
  dwarfs := []string{"Ceres", "pluto", "Haumea", "Makemake", "Eris"}

  dwarfs = append(dwarfs, "Orcus")
  fmt.Println(dwarfs)

  // append is *variadic, so multiple elements can be passed
  dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
  fmt.Println(dwarfs)

  fmt.Println(len(dwarfs))
}
