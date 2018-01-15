package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/co0p/construe/usecases"
)

func httpError(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, http.StatusText(statusCode))
}

func logRequest(r *http.Request) {
	log.Printf("%s %s by ip=%s", r.Method, r.URL.Path, r.RemoteAddr)
}

func ExportHandler(export usecases.Export) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		if r.Method != "GET" {
			httpError(w, http.StatusMethodNotAllowed)
			return
		}

		httpError(w, http.StatusNotImplemented)
		return
	}
}

func ImportHandler(Import usecases.Import) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		if r.Method == "POST" {
			httpError(w, http.StatusMethodNotAllowed)
			return
		}

		httpError(w, http.StatusNotImplemented)
		return
	}
}
