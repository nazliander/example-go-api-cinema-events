package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CinemaEvent struct {
    Date  string `json:"date"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

var CinemaEvents []CinemaEvent

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!\nSupports only GET for '/events' and POST for '/event' endpoints.")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent CinemaEvent

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEvent)
	CinemaEvents = append(CinemaEvents, newEvent)
	json.NewEncoder(w).Encode(newEvent)
}

func returnAllEvents(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(CinemaEvents)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/events", returnAllEvents)
    router.HandleFunc("/event", createEvent).Methods("POST")
    log.Fatal(http.ListenAndServe(":6661", router))
}

func main() {
    CinemaEvents = []CinemaEvent{
        CinemaEvent{Date: "2021-06-01", Id: "1", Title: "Hitchcock had 1000 birds"},
        CinemaEvent{Date: "2021-01-01", Id: "2", Title: "Tarkovsky is alive and still making movies"},
    }
    handleRequests()
}
