package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("xy")

	router := mux.NewRouter()

	_ = http.ListenAndServe(":8000", router)
}
