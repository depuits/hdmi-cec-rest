package main

import (
	"net/http"
	"./webservice"
)

func main() {
	router := webservice.GetRouter()
	http.ListenAndServe(":5000", router)
}
