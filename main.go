package main

import (
	"fmt"
	"log"
	"net/http"
	router "github.com/Aditya7880900936/postgres-golang/routers"
)

func main(){
	r := router.Router()
	fmt.Println("Starting Server on the PORT 8080.....")
	log.Fatal(http.ListenAndServe(":8080",r))
}