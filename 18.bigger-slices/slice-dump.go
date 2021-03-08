package main

import "fmt"

func sliceDump(label string, slice []string) {
  fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)

}

func main() {
  dwarfs := []string{"Ceres", "pluto", "Haumea", "Makemake", "Eris"}
  sliceDump("dwarfs", dwarfs)
  sliceDump("dwarfs[1:2]", dwarfs[1:2])
}
