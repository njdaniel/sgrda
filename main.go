package main

type Repo struct {
	Repo string   `json:"repo"`
	Deps []string `json:"deps"`
}

func main() {
	url := "https://go-search.org/api"
	// repo = Package string
	// deps = Imported []string

}
