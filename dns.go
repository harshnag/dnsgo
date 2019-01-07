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

var records []Record
var zones []Zone

func main() {
	fmt.Printf("hello")

	db, err := sql.Open("sqlite3", "./dns.db")
	if err != nil {
		log.Errorf("unable to create sqlite database: %v", err)
	}
	// Insert snippet here to handle JSON -> DB conversion
	// for zones and records relationships

	// This defines the operations we want in the api,
	// namely creating/deleting zones and
	// view and manage all of the records in those zones
	router := mux.NewRouter()
	router.HandleFunc("/record", GetRecord).Methods("GET")
	router.HandleFunc("/zone", GetZone).Methods("GET")
	router.HandleFunc("/record/{name}", SetRecord).Methods("POST")
	router.HandleFunc("/zone/{name}", DeleteZone).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Validate that Record Type is either A/CNAME
// For A, validate that the data is an IPV4 address
// For CNAME, validate that the data is a valid domain name
// For Record Name, validate either root domain if @ present or subdomain
func validateRecord() {
}

// Validate that the name is a valid domain name
func validateZone() {
}

// All of these should have DB commits at the end of each operation
func GetRecord(w http.ResponseWriter, r *http.Request)  {}
func GetZone(w http.ResponseWriter, r *http.Request)    {}
func SetRecord(w http.ResponseWriter, r *http.Request)  {}
func DeleteZone(w http.ResponseWriter, r *http.Request) {}
