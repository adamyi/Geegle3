package main

import (
	"flag"
	"log"

	"github.com/adamyi/geegle-sffe/context"
	"github.com/adamyi/geegle-sffe/web"
)

func main() {
	flagData := flag.String("data", "data",
		"The location to use for the data store")
	flagAddr := flag.String("addr", ":8067",
		"The address that the HTTP server will bind")
	flag.Parse()

	ctx, err := context.Open(*flagData)
	if err != nil {
		log.Panic(err)
	}
	defer ctx.Close()

	log.Panic(web.ListenAndServe(*flagAddr, ctx))
}
