package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://covid19.th-stat.com/api/open/today")
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(robots)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getAll", getAll).Methods(http.MethodGet)

	log.Println("listen port :1234")
	http.ListenAndServe(":1234", r)

}
