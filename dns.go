package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Record struct {
	Name       string `json:name`
	RecordType string `json:recordtype`
	RecordData string `json:recorddata`
	TTL        int
}

type Zone struct {
	Name   string   `json:name`
	Record []Record `json:record`
}

func main() {

	fmt.Printf("hello")
	router := mux.NewRouter()
	router.HandleFunc("/record", GetRecord).Methods("GET")
	router.HandleFunc("/zone", GetZone).Methods("GET")
	router.HandleFunc("/record/{name}", SetRecord).Methods("POST")
	router.HandleFunc("/zone/{name}", DeleteZone).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetRecord(w http.ResponseWriter, r *http.Request) {}
func GetZone(w http.ResponseWriter, r *http.Request)   {}
func SetRecord(w http.ResponseWriter, r *http.Request) {}
func DeleteZone(w http.ResponseWriter, r *http.Request)   {}