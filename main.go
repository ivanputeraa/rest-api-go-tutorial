package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func getAllBuildData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint being hit: getAllBuildData")
	json.NewEncoder(w).Encode(BuildData)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/buildData", getAllBuildData)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	BuildData = []BuildDatum{
		BuildDatum{ModuleName: "ShipperPicker", Duration: 10},
		BuildDatum{ModuleName: "Checkout", Duration: 20},
	}
	handleRequests()
}
