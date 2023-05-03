package main

import (
    "testing"
    "net/http"
)

func TestGetZones(t *testing.T) {
    // Create a new DB instance
    db, err := NewDB("dns.db")
    if err != nil {
        panic(err)
    }
    
    // Create a new request
    req, err := http.NewRequest("GET", "/zones", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    // Create a new response recorder
    rr := httptest.NewRecorder()
    
    // Serve the request
    db.ServeHTTP(rr, req)
    
    // Check the status code
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code 200, got %d", rr.Code)
    }
    
    // Check the response body
    zones := []Zone{}
    err = json.Unmarshal(rr.Body.Bytes(), &zones)
    if err != nil {
        t.Fatal(err)
    }
    
    // Check that the zones are non-empty
    if len(zones) == 0 {
        t.Errorf("Expected at least one zone, got 0")
    }
}

func TestGetZone(t *testing.T) {
    // Create a new DB instance
    db, err := NewDB("dns.db")
    if err != nil {
        panic(err)
    }
    
    // Create a new request
    req, err := http.NewRequest("GET", "/zones/example.com", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    // Create a new response recorder
    rr := httptest.NewRecorder()
    
    // Serve the request
    db.ServeHTTP(rr, req)
    
    // Check the status code
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code 200, got %d", rr.Code)
    }
    
    // Check the response body
    zone := Zone{}
    err = json.Unmarshal(rr.Body.Bytes(), &zone)
    if err != nil {
        t.Fatal(err)
    }
    
    // Check that the zone name is correct
    if zone.Name != "example.com" {
        t.Errorf("Expected zone name 'example.com', got '%s'", zone.Name)
    }
}

func TestGetRecords(t *testing.T) {
    // Create a new DB instance
    db, err := NewDB("dns.db")
    if err != nil {
        panic(err)
    }
    
    // Create a new request
    req, err := http.NewRequest("GET", "/records?zone=example.com", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    // Create a new response recorder
    rr := httptest.NewRecorder()
    
    // Serve the request
    db.ServeHTTP(rr, req)
    
    // Check the status code
    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code 200, got %d", rr.Code)
    }
    
    // Check the response body
    records := []Record{}
    err = json.Unmarshal(rr.Body.Bytes(), &records)
    if err != nil {
        t.Fatal(err)
    }
    
    // Check that the records are non-empty
    if len(records) == 0 {
        t.Errorf("Expected at least one record, got 0")
    }
}
