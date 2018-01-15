package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/co0p/construe/storage"
	"github.com/co0p/construe/syndication"
	"github.com/co0p/construe/translate"
	"github.com/co0p/construe/usecases"
	"github.com/co0p/construe/web"
)

var port string
var googleKey string

func init() {
	flag.StringVar(&port, "port", "8080", "port to run this application on")
	flag.StringVar(&googleKey, "googleKey", "", "The google key to use for the translation service")
}

func main() {
	flag.Parse()
	_, ok := strconv.Atoi(port)
	if (googleKey == "") || ok != nil {
		flag.Usage()
		os.Exit(0)
	}

	// construct clients handlers will use
	encoder := syndication.Encoder{}
	decoder := syndication.Decoder{}
	translationClient := translate.NewGoogleClient(googleKey)
	storageClient := storage.MemoryStorage{}

	// usecases
	importUsecase := usecases.Import{
		Decoder:           decoder,
		TranslationClient: translationClient,
		StorageClient:     storageClient,
	}

	exportUsecase := usecases.Export{
		Encoder:       encoder,
		StorageClient: storageClient,
	}

	// configure handlers
	importHandler := web.ImportHandler(importUsecase)
	exportHandler := web.ExportHandler(exportUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/import", importHandler)
	mux.HandleFunc("/api/export/", exportHandler)

	// start application
	log.Println("starting application on port: " + port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
