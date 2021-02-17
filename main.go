package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"time"
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

	fmt.Println("Server is starting...")

	go NotifyServerStarted(addr)

	err := http.ListenAndServe(addr, router)

	if err != nil {
		fmt.Println("The server failed to start")
		os.Exit(1)
	}
}

func NotifyServerStarted(addr string) {
	for {
		if _, err := net.DialTimeout("tcp", addr, time.Millisecond); err == nil {
			break
		}
	}

	fmt.Println("Server is online!")
}
