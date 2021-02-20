package main


func main() {
  var planets [8]string

  planets[8] = "Pluto" // An error only detected at runtime. Results in panic.
  pluto := planets[8]
}
