// concurrency and goroutines in golang

package main

import (
	"fmt"
	"time"
)
func concurrency() {
    for i:=1;i<=5;i++ {
        time.Sleep(200 * time.Millisecond )
        fmt.Println(i)
 }
}
func main(){

    go concurrency()

    for i:=6;i<=10;i++ {
        time.Sleep(200 * time.Millisecond)
        fmt.Println(i)
    }
}
