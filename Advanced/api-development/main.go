package main

import (
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/", homepage)
	http.HandleFunc("/product/", returnProduct)
	http.HandleFunc("/products", returnAllProducts)
	http.ListenAndServe("localhost:10000", nil)
}

func returnProduct(writer http.ResponseWriter, r *http.Request) {
	log.Println("Product")
	key := r.URL.Path[len("/product/"):]
	log.Println(key)
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
