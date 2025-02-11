```go
func main() {
    var m map[string]int
    if m == nil {
        fmt.Println("Map is nil") // Handle the nil case
    } else {
        fmt.Println(m["a"])
    }
}
```