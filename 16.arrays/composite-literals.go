package main
import "fmt"

func main() {
  // pre-populate arrays with curly braces

  dwarfs := [5]string{"Ceres", "Pluto", "Humea", "Makemake", "Eris"}

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

  fmt.Println(len(dwarfs))
  fmt.Println(len(planets))

  // 'range' keyword provides index/val iteration pairs

  for i, dwarf := range dwarfs {
    fmt.Println(i, dwarf)
  }
}
