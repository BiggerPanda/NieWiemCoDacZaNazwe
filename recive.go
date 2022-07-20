package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func redisView(w http.ResponseWriter, r *http.Request) {
	redisPool := newPool().Get()
	var value, err2 = redisPool.Do("GET", "TestKey")
	fmt.Fprintf(w, "%s \n", value)

	if err2 != nil {
		fmt.Fprintf(w, "%s \n", err2)
	}
}

func CreateClient(w http.ResponseWriter, r *http.Request) {

	redisPool := newPool().Get()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	redisPool.Do("SET", name, "TestKey123")
	var value, err2 = redisPool.Do("GET", name)
	fmt.Fprintf(w, "%s \n", value)

	if err2 != nil {
		fmt.Fprintf(w, "%s \n", err2)
	}

}
