package main

import "fmt"
var f = func() { // anonymous function (function literal)
  fmt.Println("Dress up for the masquerade.")
}

func main() {
  f()
}
