import (
    "fmt"
    "time"
)
func main() {
    resultChannel := make(chan string)

    go func(){
    result := <-resultChannel
    fmt.Println(result)
}()
    resultChannel <- "foo"
}
func fetchResource(n int) string {
    time.Sleep(time.Second * 2)
    return fmt.Sprintf("result %d",n)
}
