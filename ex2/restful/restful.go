package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "./model"
    "net/http"
)

var hostings []hosting.Hosting

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
    http.Error(w, "", http.StatusNotFound)
}

func CreateHosting(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    found := false
    var hosting hosting.Hosting
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

// Because a DELETE should be idempotent, deleting an non-existing entity should
// return a 204, the same as an existing deletion
func DeleteHosting(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    for index, item := range hostings {
        if item.ID == params["id"] {
            hostings = append(hostings[:index], hostings[index+1:]...)
            break
        }
    }
    http.Error(w, "", http.StatusNoContent)
}

// our main function
func main() {
    hostings = append(hostings, hosting.Hosting{ID: "1", Name: "Hosting1", Cores: "2", Memory: "4096", Disc: "1TB"})
    hostings = append(hostings, hosting.Hosting{ID: "2", Name: "Hosting2", Cores: "4", Memory: "8192", Disc: "500MB"})

    router := mux.NewRouter()
    router.HandleFunc("/hostings", GetHostings).Methods("GET")
    router.HandleFunc("/hosting/{id}", GetHosting).Methods("GET")
    router.HandleFunc("/hosting/{id}", CreateHosting).Methods("PUT")
    router.HandleFunc("/hosting/{id}", DeleteHosting).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}

