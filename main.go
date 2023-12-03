package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"remote-filesystem/html"
)

var user_address string

func main() {
	l, close := createListener()
	defer close()
	http.Handle("/", http.HandlerFunc(defaultForm))
	http.Handle("/login", http.HandlerFunc(connectForm))

	log.Println("Servidor iniciado")
	log.Printf("http://localhost:%v\n", l.Addr().(*net.TCPAddr).Port)
	http.Serve(l, nil)
}

func defaultForm(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, html.Index)
}

type current_dir struct {
    Dir   []string `json:"dir"`
	Files []string `json:"files"`
	Path  string   `json:"path"`
}

func connectForm(rw http.ResponseWriter, r *http.Request) {
	res, err := http.Get(r.FormValue("user_address"))
	if err != nil {
		fmt.Fprint(rw, html.ReturnErrorForm(err.Error()))
		return
	}
    user_address = r.FormValue("user_address")

    body, err := io.ReadAll(res.Body)
    if err != nil {
		fmt.Fprint(rw, html.ReturnErrorForm(err.Error()))
		return
    }
    log.Println(string(body))
    var dir current_dir
    err = json.Unmarshal(body, &dir)
    if err != nil {
		fmt.Fprint(rw, html.ReturnErrorForm(err.Error()))
		return
    }
    fmt.Println(dir)
}

func createListener() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	return l, func() {
		_ = l.Close()
	}
}
