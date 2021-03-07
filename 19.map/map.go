package main

import "fmt"

func main() {

    temperature := map[string]int{
        "Earth": 15,
        "Mars": -65,
    }

    temperature["Earth"] = 16
    temperature["Venus"] = 464

    temp := temperature["Earth"]
    fmt.Println(temperature)
    fmt.Println(temp)


    // check for non-existent property

    if moon, ok := temperature["Moon"]; ok {
        fmt.Printf("On average the moon is %v C.\n", moon)
    } else {
        fmt.Println("Where is moon??!")
    }
}

