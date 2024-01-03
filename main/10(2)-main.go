// using string functions and import and played around

package main

import ("fmt"
        "strings"
)
type Transform func (s string) string

func Uppercase(s string) string {
    return strings.ToUpper(s)
}
func Prefix(prefix string) Transform{
    return func(s string) string {
    return prefix + s
}
}
func transformString(s string,fn Transform) string {
    return fn(s)
}

func main() {
    fmt.Println(transformString("hey bro, how you doin", Uppercase))
    fmt.Println(transformString("Its me bro", Prefix("hey,")))
}
