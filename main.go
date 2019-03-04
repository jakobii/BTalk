package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func Authenticate(key string) bool {
	if key == "088d7646-3e16-11e9-b6e5-af6f2bafb279" {
		return true
	}
	return false
}

// Cmd is te request...
type Cmd struct {
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
	var cmd Cmd
	err := json.NewDecoder(req.Body).Decode(&cmd)
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// authenticate
	if !Authenticate(cmd.Key) {
		io.WriteString(w, "Authentication Error")
	}

	// make some coffee..
	Output, err := exec.Command(cmd.Command, cmd.Arguments...).Output()
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// forward output to client
	io.WriteString(w, string(Output))
}

type Bash struct {
	Key     string
	Command string
}

func (b Bash) EscapedCommand() string {
	var cmd string
	cmd = strings.ReplaceAll(b.Command, "'", "\\'")
	cmd = "$'" + cmd + "'"
	return cmd
}

func BashRunner(w http.ResponseWriter, req *http.Request) {

	//decode the json payload
	var bash Bash
	err := json.NewDecoder(req.Body).Decode(&bash)
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// authenticate
	if !Authenticate(bash.Key) {
		io.WriteString(w, "Authentication Error")
	}

	// make some coffee..
	Output, err := exec.Command("/bin/bash", "-c", bash.EscapedCommand()).Output()
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// forward output to client
	io.WriteString(w, string(Output))
}

type Powershell struct {
	Key     string
	Command string
}

func (ps Powershell) EscapedCommand() string {
	var cmd string
	cmd = strings.ReplaceAll(ps.Command, "'", "`'")
	cmd = "'" + cmd + "'"
	return cmd
}

func PsRunner(w http.ResponseWriter, req *http.Request) {

	//decode the json payload
	var ps Powershell
	err := json.NewDecoder(req.Body).Decode(&ps)
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// authenticate
	if !Authenticate(ps.Key) {
		io.WriteString(w, "Authentication Error")
	}

	// make some coffee..
	Output, err := exec.Command("pwsh", "-Command", ps.EscapedCommand()).Output()
	if err != nil {
		io.WriteString(w, fmt.Sprint(err))
	}

	// forward output to client
	io.WriteString(w, string(Output))
}

func main() {
	http.HandleFunc("/cmd", CmdRunner)
	http.HandleFunc("/bash", BashRunner)
	http.HandleFunc("/powershell", PsRunner)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
