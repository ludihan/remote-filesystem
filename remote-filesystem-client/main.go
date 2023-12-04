package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"remote-filesystem-client/html"
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
    fmt.Fprint(rw, html.ReturnDir(body))
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
