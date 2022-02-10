package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

type RootResponse struct {
	Method string `json:"method"`
	Index  int    `json:"index"`
}

type HelloResponse struct {
	Method   string `json:"method"`
	Index    int    `json:"index"`
	Greeting string `json:"greet"`
}

func main() {
	var idx int
	flag.IntVar(&idx, "i", 0, "Index munber (default: 0)")
	flag.Parse()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HelloResponse{
			Method:   r.Method,
			Index:    idx,
			Greeting: "Keep greeting to world hoever!!",
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RootResponse{
			Method: r.Method,
			Index:  idx,
		})
	})

	addr := 8080 + idx
	http.ListenAndServe(":"+strconv.Itoa(addr), nil)
}
