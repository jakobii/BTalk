package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

// CmdRequest is te request...
type CmdRequest struct {
	Key       string
	Command   string
	Arguments []string
}

// CmdRunner makes the magic happen.
func CmdRunner(w http.ResponseWriter, req *http.Request) {

	//if req.Method != "POST" {
	//	return
	//}

	//decode the json payload
	var Cmd CmdRequest
	err := json.NewDecoder(req.Body).Decode(&Cmd)
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// authenticate
	if Cmd.Key != "088d7646-3e16-11e9-b6e5-af6f2bafb279" {
		io.WriteString(w, "Authentication Error")
	}

	// make some coffee..
	Output, err := exec.Command(Cmd.Command, Cmd.Arguments...).Output()
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// forward output to client
	io.WriteString(w, string(Output))
}

func main() {
	http.HandleFunc("/cmd", CmdRunner)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
