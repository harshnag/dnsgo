package main

import (
    "database/sql"
    "fmt"
    "net/http"
)

// Zone represents a DNS zone.
type Zone struct {
    Name string `db:"name"`
}

// Record represents a DNS record.
type Record struct {
    Name string `db:"name"`
    Type string `db:"type"`
    Data string `db:"data"`
    TTL int `db:"ttl"`
}

// DB is a wrapper around a SQLite database connection.
type DB struct {
    *sql.DB
}

// NewDB creates a new DB instance.
func NewDB(file string) (*DB, error) {
    db, err := sql.Open("sqlite3", file)
    if err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

// CreateZone creates a new DNS zone.
func (db *DB) CreateZone(name string) error {
    stmt, err := db.Prepare("INSERT INTO zones (name) VALUES (?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(name)
    return err
}

// DeleteZone deletes a DNS zone.
func (db *DB) DeleteZone(name string) error {
    stmt, err := db.Prepare("DELETE FROM zones WHERE name = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(name)
    return err
}

// GetZones returns a list of all DNS zones.
func (db *DB) GetZones() ([]Zone, error) {
    stmt, err := db.Prepare("SELECT name FROM zones")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query()
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var zones []Zone
    for rows.Next() {
        var zone Zone
        err = rows.Scan(&zone.Name)
        if err != nil {
            return nil, err
        }
        zones = append(zones, zone)
    }
    return zones, nil
}

// CreateRecord creates a new DNS record.
func (db *DB) CreateRecord(zone string, name string, type string, data string, ttl int) error {
    stmt, err := db.Prepare("INSERT INTO records (zone, name, type, data, ttl) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(zone, name, type, data, ttl)
    return err
}

// DeleteRecord deletes a DNS record.
func (db *DB) DeleteRecord(zone string, name string) error {
    stmt, err := db.Prepare("DELETE FROM records WHERE zone = ? AND name = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(zone, name)
    return err
}

// GetRecords returns a list of all DNS records for a given zone.
func (db *DB) GetRecords(zone string) ([]Record, error) {
    stmt, err := db.Prepare("SELECT name, type, data, ttl FROM records WHERE zone = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(zone)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var records []Record
    for rows.Next() {
        var record Record
        err = rows.Scan(&record.Name, &record.Type, &record.Data, &record.TTL)
        if err != nil {
            return nil, err
        }
        records = append(records, record)
    }
    return records, nil
}

func main() {
    // Create a new DB instance
    db, err := NewDB("dns.db")
    if err != nil {
        panic(err)
    }
    
    // Create a new router
    router := http.NewServeMux()
    
    // Register the routes
    router.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
        // Get the zones
        zones, err := db.GetZones()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        // Marshal the zones to JSON
        json.NewEncoder(w).Encode(zones)
    })
    
    router.HandleFunc("/zones/{name}", func(w http.ResponseWriter, r *http.Request) {
        // Get the zone name
        name := r.URL.Path[len("/zones/"):]
        
        // Get the zone
        zone, err := db.GetZone(name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        
        // Marshal the zone to JSON
        json.NewEncoder(w).Encode(zone)
    })
    
    router.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {
        // Get the zone name
        name := r.URL.Query().Get("zone")
        
        // Get the records
        records, err := db.GetRecords(name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        // Marshal the records to JSON
        json.NewEncoder(w).Encode(records)
    })
    
    // Serve the API
    http.ListenAndServe(":8080", router)
}
