package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r1 := mux.NewRouter()
	r2 := mux.NewRouter()
	r3 := mux.NewRouter()
	r4 := mux.NewRouter()

	r1.HandleFunc("/", sayHello1)
	r2.HandleFunc("/", sayHello2)
	r3.HandleFunc("/", sayHello3)
	r4.HandleFunc("/", sayHello4)

	r1.HandleFunc("/app1", sayHelloR1)
	r2.HandleFunc("/app1", sayHelloR2)

	r3.HandleFunc("/app2", sayHelloR3)
	r4.HandleFunc("/app2", sayHelloR4)

	go func() {
		log.Fatal(http.ListenAndServe(":8082", r2))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8083", r3))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8084", r4))
	}()

	log.Fatal(http.ListenAndServe(":8081", r1))
}

func sayHello1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from 8081"))
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from 8082"))
}

func sayHello3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from 8083"))
}

func sayHello4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from 8084"))
}

func sayHelloR1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application1 on 8081"))
}

func sayHelloR2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application1 on 8082"))
}

func sayHelloR3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application2 on 8083"))
}

func sayHelloR4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application2 on 8084"))
}
