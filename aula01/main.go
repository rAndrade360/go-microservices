package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops", 400)
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.ListenAndServe(":3000", nil)
}
