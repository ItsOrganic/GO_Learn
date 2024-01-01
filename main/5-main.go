//enums in golang
package main

import (
	"fmt"
)

type WeaponType int

//kind of interface
func (w WeaponType) String() string{
    switch w{
        case Sword:
            return "Sword"
        case Axe:
            return "Axe"
    }
    return ""
}

//enums
const (
    Axe WeaponType=iota
    Sword
    WoodenStick
    Knife
)

func getDamage(weaponType WeaponType)int {
    switch weaponType {
        case Axe:
            return 100
        case Sword:
            return 90
        case WoodenStick:

            return 1
        case Knife:
            return 40

        default:
            panic("weapon does not exist")
    }

}
func main(){

    //interface example using the string of weaponType
    fmt.Printf("value of Axe is  %s and damage of Axe is %d \n",Axe,getDamage(Axe))

    //example of the enum creation
    fmt.Printf("value of Axe is  %d and damage of Axe is %d \n",Axe,getDamage(Axe))
    fmt.Printf("value of sword is %d and damage is %d \n",Sword,getDamage(Sword))

}
