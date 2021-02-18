package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"time"
	"xy/pkg/cfg"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "index")
	})
}

func NotifyStarted(config cfg.Config) {
	for {
		if _, err := net.DialTimeout("tcp", config.SocketAddress(), time.Millisecond); err == nil {
			break
		}
	}
	fmt.Println("Server is online!", "(" + config.SocketAddress() + ")")
}

func Start(config cfg.Config, router *mux.Router) {
	err := http.ListenAndServe(config.SocketAddress(), router)

	if err != nil {
		fmt.Println("The server failed to start")
		os.Exit(1)
	}
}
