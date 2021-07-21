package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	. "../helpers"
	"github.com/gorilla/mux"

	. "../models"
)

var productStore = make(map[string]Product)
var id int = 0

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product
	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Contenr-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}
	data, err := json.Marshal(products)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}
	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]
	var prodUpd Product
	err = json.NewDecoder(r.Body).Decode(&prodUpd)
	CheckError(err)
	if _, ok := productStore[key]; ok {
		prodUpd.ID, _ = strconv.Atoi(key)
		prodUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpd

	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
