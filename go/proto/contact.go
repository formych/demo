package main

import (
	"io/ioutil"
	"log"

	"github.com/formych/demo/go/proto/pb"

	"github.com/golang/protobuf/proto"
)

const tmpFile = "contact.txt"

func runContact() {
	writeContact()
	readContact()
}

func writeContact() {
	p1 := &pb.Person{
		Id:   1,
		Name: "小张",
		Phones: []*pb.Phone{
			{Type: pb.PhoneType_Home, Number: "11111"},
			{Type: pb.PhoneType_WORK, Number: "12345"},
		},
	}

	book := &pb.ContactBook{}
	book.Persons = append(book.Persons, p1)

	data, _ := proto.Marshal(book)
	ioutil.WriteFile(tmpFile, data, 0644)
}

func readContact() {
	data, _ := ioutil.ReadFile(tmpFile)
	book := &pb.ContactBook{}

	proto.Unmarshal(data, book)

	for _, v := range book.Persons {
		log.Println(v.Id, v.Name)
		for _, v := range v.Phones {
			log.Println(v.Type, v.Number)
		}
	}
}
