package main

import (
    "fmt"
    "time"
)
func main() {
    msgChannel := make(chan string,12)
    msgChannel <- "A"
    msgChannel <- "B"
    msgChannel <- "C"
    msgChannel <- "D"
    close(msgChannel)

    // It will great because the channel was closed so it will know how many values are there im the final code
    for msg := range msgChannel {
        fmt.Println("the message is <-",msg)
    }

}

func fetchResource(n int) string {
    time.Sleep(time.Second * 2)
    return fmt.Sprintf("result %d",n)
}
it
