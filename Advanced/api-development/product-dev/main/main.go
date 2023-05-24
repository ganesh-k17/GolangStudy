/*
Using gorilla/mux router
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/product/{id}", returnProduct)
	myRouter.HandleFunc("/products", returnAllProducts)
	http.ListenAndServe("localhost:10000", myRouter)
}

func returnProduct(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(writer).Encode(product)
		}
	}
}

func returnAllProducts(writer http.ResponseWriter, _ *http.Request) {
	log.Println("Endpoint all products hit...")

	for i := 0; i < len(Products); i++ {
		fmt.Println(Products[i].Name)
	}

	json.NewEncoder(writer).Encode(Products)
}

func homepage(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "welcome to home")
	fmt.Println("Welcome endpoint hit...")
}

func main() {

	Products = []Product{
		Product{Id: "1", Name: "Raja", Quantity: 6, Price: 200},
		Product{Id: "2", Name: "Vijay", Quantity: 7, Price: 100},
	}

	handleRequests()

	fmt.Println("end")
}
