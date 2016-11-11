package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "OH NOES, got an error!")
		return
	}
	body := fmt.Sprintf(`<h1>Hello Öredev! ✖‿✖</h1>
	<h2>My host name is %s!</h2>`, hostname)
	fmt.Fprint(w, body)
}

func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", nil))
}
