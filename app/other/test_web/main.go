package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(HandleReq))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func HandleReq(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	w.Write([]byte(id))
}
