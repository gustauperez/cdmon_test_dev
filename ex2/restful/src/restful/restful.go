package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Hosting struct {
    ID        string   `json:"id"`
    Name      string   `json:"name"`
    Cores     string   `json:"cores"`
    Memory    string   `json:"memory"`
    Disc      string   `json:"disc"`
}

var hostings []Hosting

func GetHostings(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(hostings)
}

func GetHosting(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range hostings {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Hosting{})
}

func CreateHosting(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    found := false
    var hosting Hosting
    _ = json.NewDecoder(r.Body).Decode(&hosting)
    for id, item := range hostings {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            hosting.ID = params["id"]
            hostings[id] = hosting
            found = true
            return
        }
    }
    if found == false {
        hosting.ID = params["id"]
        hostings = append(hostings, hosting)
    }
    json.NewEncoder(w).Encode(hosting)
}

func DeleteHosting(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    for index, item := range hostings {
        if item.ID == params["id"] {
            hostings = append(hostings[:index], hostings[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(hostings)
    }
}

// our main function
func main() {
    hostings = append(hostings, Hosting{ID: "1", Name: "Hosting1", Cores: "2", Memory: "4096", Disc: "1TB"})

    router := mux.NewRouter()
    router.HandleFunc("/hostings", GetHostings).Methods("GET")
    router.HandleFunc("/hosting/{id}", GetHosting).Methods("GET")
    router.HandleFunc("/hosting/{id}", CreateHosting).Methods("POST")
    router.HandleFunc("/hosting/{id}", DeleteHosting).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
