package main

import (
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/formych/go-proto/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	http.HandleFunc("/proto", func(w http.ResponseWriter, r *http.Request) {
		x, err := ioutil.ReadAll(r.Body)
		user := &pb.User{}
		err = proto.Unmarshal(x, user)
		log.Println(user, err)
	})

	http.ListenAndServe(":8080", nil)
}
