package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string, 12)
	msgChannel <- "A"
	msgChannel <- "B"
	msgChannel <- "C"
	msgChannel <- "D"
	close(msgChannel)
 // a for approach where the for is true and will run infintely but after the (ok) implementation it will know where to stop
	for {
		msg, ok := <-msgChannel
		if !ok {
			break
		}
		fmt.Println("The message is ->", msg)
	}

}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
