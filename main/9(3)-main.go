// inherited a type and used to print names of the PLayers

package main
import "fmt"

type Player struct{
    player1, player2 string
}

type Position struct {
    Player
}

func (p Player) Name(i,j string) {
    fmt.Println("the names of the player are:",i,j)
}

func main() {
    p := Player{}
    p.Name("pele","messi")
}
