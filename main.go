package main

import (
	"fmt"
	"net/http"
	"github.com/depuits/hdmi-cec-rest/webservice"
)

func main() {
	fmt.Println("Starting server")
	router := webservice.GetRouter()
	http.ListenAndServe(":5000", router)
}
