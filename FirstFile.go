package FirstFile

import (
    "fmt"
    "net/http"
)

func Sub(x, y int) int {
	return x - y
}

func Mul(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}


func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Build & Deploy Success")
}