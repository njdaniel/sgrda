package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"

	"redtoad.co/sgrda/cassandra"
)

// Repo for json GET from gosearch
type Repo struct {
	Repo string   `json:"repo"`
	Deps []string `json:"deps,omitempty"`
}

// Repos slice of struct repo
type Repos []Repo

func main() {
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	// url := "https://go-search.org/api?action=package_depends"
	// repo = Package string
	// deps = Imported []string

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Is there a get all?
	// Or have to cycle through the packages?

	// Get all packages
	resp, err := http.Get("https://go-search.org/api?action=packages")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Resp: %v\n", resp)
	fmt.Print("Resp.body: ")
	fmt.Println(resp.Body)
	// rData, _ := ioutil.ReadAll(resp.Body)
	// var data Repos
	var data []string
	err = json.NewDecoder(resp.Body).Decode(&data)
	fmt.Print("Repos: ")
	fmt.Println(data)

	// TODO: Insert slice of repos to cassandra to ref later
	var created = false
	var gocqlUUID gocql.UUID
	gocqlUUID = gocql.TimeUUID()
	if err := cassandra.Session.Query(`INSERT INTO repolist (id, repos) VALUES (?,?)`, gocqlUUID, data).Exec(); err != nil {
		log.Fatal(err)
	} else {
		created = true
	}

	// Check Cassandra
	fmt.Println(created)
	if created {
		fmt.Println("user_id", gocqlUUID)
	} else {
		fmt.Println("errors")
	}
	// TODO: Cycle through repos and pop cassandra db for each repo and who imports it

	// Package_depends Action:
	// returns array of the Package, and Imported

}
