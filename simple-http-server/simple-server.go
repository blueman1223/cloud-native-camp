package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("MY_SERVICE_PORT")
	if port == "" {
		fmt.Println("Env MY_SERVICE_PORT not set, default is 8080")
		port = "8080"
	}

	flag.Set("v", "4")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthz handler")

	var msgArr []string
	defer func() { // print log
		for _, msg := range msgArr {
			fmt.Println(msg)
		}
	}()

	msgArr = append(msgArr, r.RemoteAddr) // client ip

	ver := os.Getenv("VERSION")
	fmt.Println(ver)
	// copy header
	for k, vArr := range r.Header {
		for _, v := range vArr {
			w.Header().Add(k, v)
		}
	}
	w.Header().Add("VERSION", ver)
	w.WriteHeader(http.StatusOK)
	msgArr = append(msgArr, http.StatusText(http.StatusOK))
	io.WriteString(w, "ok\n")

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
