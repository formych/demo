package main

import (
	"encoding/json"
	"log"

	"github.com/formych/demo/go/proto/pb"

	"github.com/golang/protobuf/proto"
)

func runUser() {
	info := &pb.Info{
		Name:  "Formych",
		Email: "welcome@gmail.com",
		Addr:  "CN, BeiJing",
		NickInfo: []*pb.NickInfo{
			{Name: "红薯", Time: 1},
			{Name: "我兴我中华", Time: 2},
			{Name: "zaiyehubugelie", Time: 3},
		},
		Likes: map[string]string{"book": "平凡的世界", "sport": "run"},
	}

	u := &pb.User{
		Id:   2011118120,
		Info: info,
	}

	log.Println("---------------log.Println---------------")
	log.Println(u)
	log.Println()
	log.Println(*u)
	log.Println()

	log.Println("---------------log.Printf +v u----------")
	log.Printf("%+v\n", u)
	log.Println()
	log.Printf("%+v\n", *u)
	log.Println()

	log.Println("---------------protobuf------------------")
	x, _ := proto.Marshal(u)
	log.Printf("%x\n", x)
	log.Println(string(x))
	log.Println()

	log.Println("---------------json----------------------")
	jdata, _ := json.Marshal(u)
	log.Printf("%x\n", jdata)
	log.Println(string(jdata))

	log.Println("---------------size----------------------")
	log.Printf("size\t%s\t%s\n", "protobuf", "json")
	log.Printf("size\t%d\t\t%d\n", len(x), len(jdata))
}
