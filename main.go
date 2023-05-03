package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashutosh001s/mongoapi/router"
)

func main() {
	fmt.Println("MongoDb API in golang")
	r:= router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at : http://localhost:4000")
}