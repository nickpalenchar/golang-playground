package greetings

import "fmt"

func Hello(name string) string {
    message := fmt.Printf("Hi, %v. Welcome!", name)
    return message
}
