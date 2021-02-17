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
var addr string

func main() {
	fmt.Println("xy")

	ParseFlags()
	addr = CreateSocketAddress()

	router := mux.NewRouter()

	fmt.Println("Server is starting...")

	go NotifyServerStarted()

	StartServer(router)
}

func ParseFlags() {
	flag.IntVar(&port, "port", port, "Port the HTTP server will listen on")
	flag.StringVar(&host, "host", host, "Host the HTTP server will listen on")
	flag.Parse()
}

func CreateSocketAddress() string {
	return host + ":" + fmt.Sprint(port)
}

func NotifyServerStarted() {
	for {
		if _, err := net.DialTimeout("tcp", addr, time.Millisecond); err == nil {
			break
		}
	}
	fmt.Println("Server is online!")
}

func StartServer(router *mux.Router) {
	err := http.ListenAndServe(addr, router)

	if err != nil {
		fmt.Println("The server failed to start")
		os.Exit(1)
	}
}
