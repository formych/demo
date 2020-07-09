obj = go-proto

build :
	protoc --go_out=. proto3/*
	go build -o $(obj)

run : build
	./go-proto

clean :
	rm -rf contact.txt $(obj)
