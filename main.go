package main

import (
	"fmt"
	"net/http"
)

// Repo for json GET from gosearch
type Repo struct {
	Repo string   `json:"repo"`
	Deps []string `json:"deps"`
}

func main() {
	url := "https://go-search.org/api?action=package_depends"
	// repo = Package string
	// deps = Imported []string

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	// Is there a get all?
	// Or have to cycle through the packages?

	// Package_depends Action:
	// returns array of the Package, and Imported

}
