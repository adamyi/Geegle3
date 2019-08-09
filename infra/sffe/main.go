package main

import (
	"flag"
	"log"

	"geegle.org/infra/sffe/context"
	"geegle.org/infra/sffe/web"
)

func main() {
	flagData := flag.String("data", "data",
		"The location to use for the data store")
	flagAddr := flag.String("addr", ":80",
		"The address that the HTTP server will bind")
	flag.Parse()

	ctx, err := context.Open(*flagData)
	if err != nil {
		log.Panic(err)
	}
	defer ctx.Close()

	log.Panic(web.ListenAndServe(*flagAddr, ctx))
}
