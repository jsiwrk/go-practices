package main

import (
	"github.com/joansais/go-practices/wiki"
	"log"
	"flag"
)

const (
	DEFAULT_ADDR = ":8080"
	DEFAULT_ASSETS_DIR = "assets/wiki"
	DEFAULT_STORAGE_DIR = "data/wiki/pages"
)

func main() {
	addr := flag.String("addr", DEFAULT_ADDR, "network address to listen on")
	assetsDir := flag.String("assets", DEFAULT_ASSETS_DIR, "location of HTML templates")
	storageDir := flag.String("storage", DEFAULT_STORAGE_DIR, "storage directory for wiki pages")
	flag.Parse()

	store := wiki.NewDiskStore(*storageDir)
	syntax := wiki.NewMarkdownSyntax(store)
	server := wiki.NewServer(store, syntax, *assetsDir)
	err := server.Start(*addr)
	if err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}
}
