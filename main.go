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

var config Config

func main() {
	fmt.Println("xy")

	LoadEnvironment()
	ParseFlags()

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

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println("Error parsing PORT environment variable")
		os.Exit(1)
	}

	config = Config{Host: host, Port: port}
}

func ParseFlags() {
	flag.IntVar(&config.Port, "port", config.Port, "Port the HTTP server will listen on")
	flag.StringVar(&config.Host, "host", config.Host, "Host the HTTP server will listen on")
	flag.Parse()
}

func NotifyServerStarted() {
	for {
		if _, err := net.DialTimeout("tcp", config.SocketAddress(), time.Millisecond); err == nil {
			break
		}
	}
	fmt.Println("Server is online!", "(" + config.SocketAddress() + ")")
}

func StartServer(router *mux.Router) {
	err := http.ListenAndServe(config.SocketAddress(), router)

	if err != nil {
		fmt.Println("The server failed to start")
		os.Exit(1)
	}
}
