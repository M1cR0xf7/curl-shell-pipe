package main

import (
	"os"
	"fmt"
	"net/http"
	"strings"
)

func loadFile(path string) string {
	body, err := os.ReadFile(path)
	if err != nil {
		panic("Error reading file")
	}
	return string(body[:])
}


func handler(w http.ResponseWriter, r *http.Request)  {
	ua := r.UserAgent()
	var file string

	evil := (strings.Contains(ua, "curl") || strings.Contains(ua, "wget"))

	if evil {
		file = loadFile("evil.sh")
	} else {
		file = loadFile("good.sh")
	}

	fmt.Fprintf(w, "%s\n", file)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
