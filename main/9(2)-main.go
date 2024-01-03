package main
import "fmt"

type Color int

func (c Color) String() string {
    switch c {
        case ColorBlue:
            return "Blue"
        case ColorBlack:
            return "Black"
        case ColorYellow:
            return "Yellow"
        case ColorPink:
            return "Pink"
        default:
            panic("color not found")
    }
}

const (
    ColorBlue Color = iota
    ColorBlack
    ColorYellow
    ColorPink
)

func main() {
    fmt.Println("The color is:",ColorBlack)
}
