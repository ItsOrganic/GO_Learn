package main
import "fmt"

type CustomMap[K comparable, V any] struct {
    data map[K]V
}
func (m *CustomMap[K, V]) Insert(k K, v V) error {
    m.data[k] = v
    return nil
}
func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
    return &CustomMap[K, V]{
        data: make(map[K]V),
    }
}

func foo[T any, B any](val T,x B) {
    fmt.Println(val,x)
}

func main() {
    foo[string, int]("agg",1)
}