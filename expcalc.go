// If skills I'm learning can be quantified by the number of hours
// practiced, whats an exponential way to quantify the number I'm
// doing so that I can level up 10 (or 20) times?


package main

import (
    "fmt"
    "math"
)


func main() {
    mini()
}


func mini() {
    // calculation for mini skill, to level 10.

    // 10(10 + n)

    levels := make(map[float64]float64, 10)
    var total float64 = 0
    for n := 0.0; n < 10; n++ {
        exp := 10 * (10 + math.Round( (n*n*n) / ((n+1)/4) )  )
        fmt.Println(exp)
        levels[n] = exp

        total += exp
    }

    fmt.Println(levels)
    fmt.Println(total)
}

