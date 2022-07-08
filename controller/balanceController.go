package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"seleksi-compfest-backend/entity"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity.StoreBalance)
}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var balance entity.Balance
	json.Unmarshal(requestBody, &balance)

	entity.StoreBalance = balance

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}

func AddBalance(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var balance entity.Balance
	json.Unmarshal(requestBody, &balance)

	log.Println(balance)

	entity.StoreBalance.Amount = entity.StoreBalance.Amount + balance.Amount

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.StoreBalance)
}

func SubstractBalance(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var balance entity.Balance
	json.Unmarshal(requestBody, &balance)

	entity.StoreBalance.Amount = entity.StoreBalance.Amount - balance.Amount

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.StoreBalance)
}
