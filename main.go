package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func sendSound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := "https://" + os.Getenv("auth") + "@audio.chirp.io/v3/" + vars["freq"] + "/" + vars["payload"]
	http.Redirect(w, r, url, 301)
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func handleRequests() {

	addr := determineListenAddress()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/v1/audio/{freq}/{payload}", sendSound).Methods("GET")
	log.Fatal(http.ListenAndServe(addr, myRouter))
}

func main() {
	handleRequests()
}
