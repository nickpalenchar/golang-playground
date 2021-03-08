package main

import "fmt"

func main() {
  dwarfs1 := []string{"Ceres", "pluto", "Haumea", "Makemake", "Eris"} // allocates Array A
  dwarfs2 := append(dwarfs1, "Orcus") // no room, copied to Array B (2x size)
  dwarfs3 := append(dwarfs2, "Salcia", "Quaoar", "Sedna") // room, stays with Array B

  dwarfs3[1] = "Pluto!"

  // what is the result of dwarfs1 and dwarfs2?

  fmt.Println(dwarfs1) // Array A
  fmt.Println(dwarfs2) // Array B

}
