package main

import (
	"databaseconnection/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.ServeHome).Methods("GET")
	r.HandleFunc("/getallrecords", controller.GetAllRecords).Methods("GET")
	r.HandleFunc("/addrecord", controller.AddARecord).Methods("POST")               
	r.HandleFunc("/deleterecord/{name}", controller.DeleteByName).Methods("DELETE") 
	r.HandleFunc("/getrecord/{name}", controller.GetByName).Methods("GET")
	r.HandleFunc("/updaterecord/{name}", controller.UpdateByName).Methods("PUT")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+"3000", r))
}

