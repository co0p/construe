package web

import (
	"net/http"

	"github.com/co0p/construe/usecases"
)

func getIdFromRequest(r *http.Request) string {
	return r.URL.Query().Get("id")
}

func ExportHandler(export usecases.Export) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := getIdFromRequest(r)

		if r.Method == "GET" && id != "" {
			if article, err := export.ById(id); err != nil {
				v, ok := err.(T)
				switch v {
				case ArticleNotFoundError:
					w.WriteHeader(http.StatusNotFound)
					break
				}
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func ImportHandler(Import usecases.Import) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
