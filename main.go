package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a Customer struct
type Customer struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Role      string `json:"Role"`
	Email     string `json:"Email"`
	Phone     int    `json:"Phone"`
	Contacted bool   `json:"Contacted"`
}

// create the Customers variable of Customer's type
var Customers []Customer

// homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Welcome in the CRM backend!</H1>")
	fmt.Println("Endpoint Hit: homePage")
}

// create the function for to obtain all the list of customer
func getCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getcustomers")
	json.NewEncoder(w).Encode(Customers)
}

// create the function for to select a customer
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, Customer := range Customers {
		if Customer.Id == key {
			json.NewEncoder(w).Encode(Customer)
		}
	}
}

// create the function for to create a new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var customer Customer
	json.Unmarshal(reqBody, &customer)
	Customers = append(Customers, customer)
	json.NewEncoder(w).Encode(customer)
}

// create the function for to delete the customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, customer := range Customers {
		if customer.Id == id {
			Customers = append(Customers[:index], Customers[index+1:]...)
		}
	}

}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

// The handleRequests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/customers", getCustomers)
	myRouter.HandleFunc("/customer/{id}", getCustomer)
	myRouter.HandleFunc("/customer", addCustomer).Methods("POST")
	myRouter.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
	myRouter.HandleFunc("/customers", updateCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

// The main function
func main() {

	Customers = []Customer{
		Customer{Id: "1", Name: "Yoan", Role: "contributor", Email: "yoan@email.com", Phone: 123654, Contacted: true},
		Customer{Id: "2", Name: "Daniel", Role: "admin", Email: "daniel@email.com", Phone: 852456, Contacted: false},
		Customer{Id: "3", Name: "Elcid", Role: "contributor", Email: "elcid@email.com", Phone: 753951, Contacted: false},
		Customer{Id: "4", Name: "Antonio", Role: "contributor", Email: "antonio@email.com", Phone: 951753, Contacted: true},
	}
	handleRequests()
}
