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
	flagFiles := flag.String("initfiles", "/initfiles",
		"The directory containing initial files for initiating")
	flagAddr := flag.String("addr", ":80",
		"The address that the HTTP server will bind")
	flag.Parse()

	ctx, err := context.Open(*flagData)
	if err != nil {
		log.Panic(err)
	}
	defer ctx.Close()

	context.InitFiles(ctx, *flagFiles)

	log.Panic(web.ListenAndServe(*flagAddr, ctx))
}
