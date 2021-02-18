package env

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"xy/pkg/cfg"
)

func LoadEnvironment(config *cfg.Config) {
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

	config.Host = host
	config.Port = port
}

func ParseFlags(config *cfg.Config) {
	flag.IntVar(&config.Port, "port", config.Port, "Port the HTTP server will listen on")
	flag.StringVar(&config.Host, "host", config.Host, "Host the HTTP server will listen on")
	flag.Parse()
}
