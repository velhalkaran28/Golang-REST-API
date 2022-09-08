package app

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Karan", "Chiplun", "415605"},
		{"ABC", "Mumbai", "123456"},
		{"XYZ", "Pune", "411021"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
