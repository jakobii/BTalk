package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"encoding/json"
	"io"
)

type Bash struct {
	Auth    string
	Command string
	Arguments []string
}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		
		var b Bash
		err := json.NewDecoder(req.Body).Decode(&b)
		if err != nil {
			io.WriteString(w, fmt.Sprint(err))
		}
		
		// authenticate
		if b.Auth != "Jacob" {
			io.WriteString(w, "Authentication Error")
		}

		out, err := exec.Command(b.Command,b.Arguments...).Output()
		if err != nil {
			io.WriteString(w, fmt.Sprint(err))
		}
		io.WriteString(w, string(out))
	}

	http.HandleFunc("/bash", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
