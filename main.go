package main

import (
	"fmt"
	"net/http"
	"./webservice"
)

func main() {
	fmt.Println("Starting server")
	router := webservice.GetRouter()
	http.ListenAndServe(":5000", router)
}
