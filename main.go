package main

import (
	"encoding/json"
	"fmt"
	"goCRM/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// creating mock data
var sampleCustomers = model.CreateSampleCustomers()

// GetALlCustomers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sampleCustomers)
}

// get customer with id
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)["id"]

	id, err := strconv.Atoi(vars)

	customersWithID := []model.Customer{}
	if err == nil {
		for _, customer := range sampleCustomers {
			if customer.ID == id {
				customersWithID = append(customersWithID, customer)
			}
		}
		fmt.Printf("Length of all customers is : %v", len(customersWithID))
		if len(customersWithID) == 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customersWithID)
			return
		}
	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
}

// add new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customerToAdd model.Customer

	err := json.NewDecoder(r.Body).Decode(&customerToAdd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sampleCustomers = append(sampleCustomers, customerToAdd)

	err = json.NewEncoder(w).Encode(&customerToAdd)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}

// remove customer with id
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)["id"]

	id, err := strconv.Atoi(vars)

	customersWithID := model.Customer{}
	if err == nil {

		for index, customer := range sampleCustomers {
			if customer.ID == id {
				customersWithID = sampleCustomers[index]
				sampleCustomers = append(sampleCustomers[:index], sampleCustomers[index+1:]...)
				break
			}
		}

		if len(customersWithID.Name) == 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customersWithID)
			return
		}
	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
}

// update customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)["id"]

	var customerToUpdate model.Customer

	err := json.NewDecoder(r.Body).Decode(&customerToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, idErr := strconv.Atoi(vars)
	if idErr == nil {
		for index, customer := range sampleCustomers {
			if customer.ID == id {
				sampleCustomers[index] = customerToUpdate
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(customerToUpdate)
				return
			}
		}
	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	fmt.Println("STARTING SERVER ON PORT localhost:3000   .....")

	http.ListenAndServe(":3000", router)

}
