// Advanced type techniques 
package main

import (
	"fmt"
)

type Entity struct {
    name string
    id string
    version string
    posx int
    posy int
}

type SpecialEntity struct {
    Entity  //inherited the complete entity struct above
    specialField float64
}

func main() {
    e:=&Entity{
        name:"My-entity",
        id: "id 1",
        version: "v1",
        posx: 100,
        posy: 200,
    }
    fmt.Printf("%+v",e)
}
