package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port to listen")
	flag.Parse()
	fmt.Printf("\nport is: %d\n\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
