package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		r.ParseMultipartForm(2 << 10)

		log.Printf("%+v, %+v, %s\n", r.PostForm, r.Form["name"], r.PostFormValue("name"))

		s, _ := ioutil.ReadAll(r.Body)
		log.Printf("body: %s\n", string(s))
	})
	http.ListenAndServe(":8080", nil)

	// curl -X POST '127.0.0.1:8080' \
	// --header 'Content-Type: application/x-www-form-urlencoded' \
	// --data-urlencode 'name=formych'

	// curl -X POST '127.0.0.1:8080' \
	// --form 'name=formych'

	// curl -X 'POST' http://127.0.0.1s:8080 \
	// -d '{"name": "formych"}' -H 'Content-Type:application/json'

}
