package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var host string = "0.0.0.0"
var port int = 8000

func main() {
	fmt.Println("xy")

	flag.IntVar(&port, "port", port, "Port the HTTP server will listen on")
	flag.StringVar(&host, "host", host, "Host the HTTP server will listen on")
	flag.Parse()

	router := mux.NewRouter()

	addr := host + ":" + fmt.Sprint(port)

	_ = http.ListenAndServe(addr, router)
}
