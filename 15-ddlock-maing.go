package main
import (
    "fmt"
    "time"
)
func main() {
    msgChannel := make(chan string,128)
    msgChannel <- "A"
    msgChannel <- "B"
    msgChannel <- "C"
    msgChannel <- "D"


    //will create deadlock because it will iterate to all 128 as it doesnt know that its null or empty 

    for msg := range msgChannel {
        fmt.Println("the message is <-",msg)
    }

}

func fetchResource(n int) string {
    time.Sleep(time.Second * 2)
    return fmt.Sprintf("result %d",n)
}
