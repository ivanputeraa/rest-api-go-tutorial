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

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", getAllData)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	mypackage.PrintHello()
	BuildData = []BuildDatum{
		BuildDatum{ModuleName: "ShipperPicker", Duration: 10},
		BuildDatum{ModuleName: "Checkout", Duration: 20},
	}
	handleRequests()
}
