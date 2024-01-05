package main
import (
    "fmt"
    "time"
)
func main() {
    resultChannel := make(chan string)
    //resultChannel := make(chan string, 10) // 10 here is the buffer
    resultChannel <- "foo"
    result := <-resultChannel
    fmt.Println(result)
}
func fetchResource(n int) string {
    time.Sleep(time.Second * 2)
    return fmt.Sprintf("result %d",n)
}
