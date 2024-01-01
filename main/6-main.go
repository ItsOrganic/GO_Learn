package main

import (
	"fmt"
)
func main(){
    //traversing in slices
    numbers := []int{1,2,3,4,5,6,7}
    for i := 0; i<len(numbers);i++{
        fmt.Println("Slices number:",numbers[i])
    }

    for _,name := range numbers {
        fmt.Println(name)
    }

    for i := 0; i<10;i++{
        fmt.Println("Iteration: ",i)
    }

    //traversing in map[string]int
    users := map[string]int{
        "foo": 1,
        "bar": 2,
        "car": 3,
    }
    for key,value := range users{
        fmt.Println(key, value)
    }

    //example for a switch case
    name := "random"

    switch name {
    case "Alice":
        fmt.Print("Alice")
    case "Bob":
        fmt.Print("Bob")
    default:
        fmt.Print("Name is default")
    }

}
