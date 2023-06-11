package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mymodule/mypackage"
	"net/http"

	"github.com/gorilla/mux"
)

// MARK: - Structs -
type BuildDatum struct {
	Id         string `json:"id"`
	ModuleName string `json:"module_name"`
	Duration   int    `json:"duration"`
}

var BuildData []BuildDatum

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint being hit")
}

func getAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint being hit: getAllData")
	json.NewEncoder(w).Encode(BuildData)
}

func getSingleData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint being hit: getSingleData")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, data := range BuildData {
		if data.Id == key {
			json.NewEncoder(w).Encode(data)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/data", getAllData)
	myRouter.HandleFunc("/data/{id}", getSingleData)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	mypackage.PrintHello()
	BuildData = []BuildDatum{
		{Id: "0", ModuleName: "ShipperPicker", Duration: 10},
		{Id: "1", ModuleName: "Checkout", Duration: 20},
	}
	handleRequests()
}
