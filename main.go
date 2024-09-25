package main

import (
	"encoding/json"
	"fmt"
	"goCRM/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var sampleCustomers = model.CreateSampleCustomers()

// GetALlCustomers
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sampleCustomers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// fmt.Println("Getting customer")

	vars := mux.Vars(r)["id"]

	id, err := strconv.Atoi(vars)

	// fmt.Println("id is " + strconv.Itoa(id))

	customersWithID := []model.Customer{}
	if err == nil {
		// fmt.Println("id is not nil")
		for _, customer := range sampleCustomers {
			// fmt.Println("Comparing " + strconv.Itoa(customer.ID) + " to " + strconv.Itoa(id))
			if customer.ID == id {
				customersWithID = append(customersWithID, customer)
				// json.NewEncoder(w).Encode(customer)
			}
		}
		fmt.Printf("Length of all customers is : %v", len(customersWithID))
		if len(customersWithID) == 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customersWithID)
		}
	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

}

func main() {

	// fmt.Println(sampleCustomers)

	router := mux.NewRouter().StrictSlash(true)
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	// http.Handle("/", r)

	router.HandleFunc("/customers", GetAllCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", GetCustomer).Methods("GET")

	fmt.Println("STARTING SERVER ON PORT localhost:3000   .....")

	http.ListenAndServe(":3000", router)

}
