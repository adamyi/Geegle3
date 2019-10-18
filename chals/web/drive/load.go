package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	in, err := ioutil.ReadFile("./b")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	req := &FileRequest{}
	if err := proto.Unmarshal(in[64:], req); err != nil {
		log.Fatalln("Failed to parse the request:", err)
	}
	fmt.Println(req)
}
