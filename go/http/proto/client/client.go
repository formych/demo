package main

import (
	"bytes"
	"log"
	"net/http"

	pb "github.com/formych/go-proto/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := &pb.User{
		Id: 2011118120,
		Info: &pb.Info{
			Name:  "Formych",
			Email: "welcome@gmail.com",
			Addr:  "CN, BeiJing",
			NickInfo: []*pb.NickInfo{
				{Name: "红薯", Time: 1},
				{Name: "我兴我中华", Time: 2},
				{Name: "zaiyehubugelie", Time: 3},
			},
			Likes: map[string]string{"book": "平凡的世界", "sport": "run"},
		},
	}

	x, _ := proto.Marshal(u)
	log.Printf("二进制: %08b\n\n", x)
	log.Printf("八进制: %o\n\n", x)
	log.Printf("十进制: %d\n\n", x)
	log.Printf("十六进制: %x\n\n", x)
	log.Printf("字符串: %s\n\n", string(x))

	log.Println("发起请求")
	resp, err := http.Post("http://127.0.0.1:8080/proto", "", bytes.NewReader(x))
	defer resp.Body.Close()
	log.Printf("请求结果: %+v, err: %v", *resp, err)
}
