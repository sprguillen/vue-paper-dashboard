package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
    "net/http"
    "math/rand"
    "time"
    "strconv"
)

type account struct {
	Email 		string `json:"Email"`
	Password 	string `json:"Password"`
}

type dateRange struct {
	FromDate 	string `json:"fromDate"`
	ToDate		string `json:"toDate"`
}

type stats struct {
	StatType 	string
	Icon 		string
	Title 		string
	Value 		string
	FooterText 	string
	FooterIcon 	string
}

var accounts []account

func main() {
    router := mux.NewRouter()
    accounts = append(accounts, account {
		Email: "john@smith.com",
		Password: "mypassword",
	})
    router.HandleFunc("/login", login).Methods("POST")
    router.HandleFunc("/statcards/capacity", getCapacity).Methods("GET")
    handler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST"},
        AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type"},
    }).Handler(router)

    log.Print("Connecting to server 8000")
    log.Print("Successfully connected to server 8000")
    log.Fatal(http.ListenAndServe(":8000", handler))
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var a account

	err := decoder.Decode(&a)
    if err != nil {
        panic(err)
    }
	for _, item := range accounts {
		if (item.Email == a.Email && item.Password == a.Password) {
			json.NewEncoder(w).Encode(true)
			return
		}
	}
	json.NewEncoder(w).Encode(false)
}

func getCapacity(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

    fromDate := keys.Get("fromDate")
    toDate := keys.Get("toDate")

    if (fromDate != "" && toDate != "") {
    	t, err := time.Parse(time.RFC3339Nano, fromDate)
    	if err != nil {
    		panic(err)
    	}
    	rand.Seed(t.Unix())
    	fromDateRand := 10 + rand.Intn(200-10)
    	toDateRand := 10 + rand.Intn(200-10)
    	totalCapacity := strconv.Itoa(fromDateRand + toDateRand)

    	var capacityStats stats
    	capacityStats = stats {
			StatType: "warning",
			Icon: "ti-server",
			Title: "Capacity",
			Value: totalCapacity + "GB",
			FooterText: "Updated now",
			FooterIcon: "ti-reload",
		}

		json.NewEncoder(w).Encode(capacityStats)
	} else {
		json.NewEncoder(w).Encode(false)
	}
}