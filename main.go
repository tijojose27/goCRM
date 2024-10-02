package main

import (
	"encoding/json"
	"fmt"
	"goCRM/model"
	"net/http"

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

	id := mux.Vars(r)["id"]

	customersWithID := []model.Customer{}
	for _, customer := range sampleCustomers {
		if customer.ID == id {
			customersWithID = append(customersWithID, customer)
		}
	}
	fmt.Printf("Length of all customers is : %v", len(customersWithID))
	if len(customersWithID) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customersWithID)
		return
	}
}

// add new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var postCustomerData struct {
		Name      string `json:"name"`
		Role      string `json:"role"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Contacted bool   `json:"contacted"`
	}

	err := json.NewDecoder(r.Body).Decode(&postCustomerData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newCustomer := model.NewCustomer(postCustomerData.Name, postCustomerData.Role, postCustomerData.Email, postCustomerData.Phone, postCustomerData.Contacted)

	sampleCustomers = append(sampleCustomers, newCustomer)

	err = json.NewEncoder(w).Encode(&newCustomer)
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

	id := mux.Vars(r)["id"]

	customersWithID := model.Customer{}

	for index, customer := range sampleCustomers {
		if customer.ID == id {
			customersWithID = sampleCustomers[index]
			sampleCustomers = append(sampleCustomers[:index], sampleCustomers[index+1:]...)
			break
		}
	}

	if len(customersWithID.Name) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customersWithID)
		return
	}
}

// update customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var customerToUpdate model.Customer

	err := json.NewDecoder(r.Body).Decode(&customerToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for index, customer := range sampleCustomers {
		if customer.ID == id {
			sampleCustomers[index] = customerToUpdate
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customerToUpdate)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func showLandingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", showLandingPage)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	fmt.Println("STARTING SERVER ON PORT localhost:3000   .....")

	http.ListenAndServe(":3000", router)

}
