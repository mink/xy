package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var host string
var port int
var addr string

func main() {
	fmt.Println("xy")

	LoadEnvironment()
	ParseFlags()

	addr = CreateSocketAddress()
	router := mux.NewRouter()

	fmt.Println("Server is starting...")

	go NotifyServerStarted()

	StartServer(router)
}

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	host = os.Getenv("HOST")
	port, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println("Error parsing PORT environment variable")
		os.Exit(1)
	}
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
