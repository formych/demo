package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type handler struct{}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World From Custom ServeHTTP"))
}

func main() {
	// http包默认的
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World From Default"))
	})
	go http.ListenAndServe(":8080", nil)

	// 使用自定义的mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World From Default Custom Mux"))
	})
	go func() {
		log.Println(http.ListenAndServe(":8081", mux))
	}()

	// 自定义Server和mux
	server := &http.Server{
		Addr:         ":8082",
		WriteTimeout: 2 * time.Second,
		Handler:      mux,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server Closed Under request")
			} else {
				log.Println("Server Closed Under unexpected")
			}
		}
	}()

	// 自定义实现ServeHTTP
	go http.ListenAndServe(":8083", &handler{})

	// 关闭服务退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := server.Close(); err != nil {
		log.Fatalln("Closed seever failed: ", err)
	}
	time.Sleep(1 * time.Second)
}
