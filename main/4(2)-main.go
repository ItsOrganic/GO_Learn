package main

import (
	"fmt"
)


var(
    floatVar32 float32 = 0.1
    floatVar64 float64 = 0.1
    name       string = "Foo"
    intVar32   int32 = 1
    intVar64   int64 = 1
    intVar     int = -1
    uintVar    uint = 1
    uintVar32  uint32 = 1
    uintVar64  uint64 = 10
    uint8Var   uint8 = 0x1
    byteVar    byte  = 0x2
    runVar      rune = 'a'
)

type Player struct{
    name    string
    health  int
    attackPower float64
}

func (player Player) getHealth() int {
return player.health
}


func main() {
    //slices creation
    numbers := []int{}
    numbers = append(numbers,4)
    numbers = append(numbers,5)

    //arrays creation
    fmt.Println(numbers)
    arr := [2]int{}

    arr[0] = 5
    arr[1] = 6
    fmt.Println(arr)

}
