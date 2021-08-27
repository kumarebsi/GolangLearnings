package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	l := log.Logger()
	handlers.NewHello(l)
	http.ListenAndServe(":9090", nil)
}
