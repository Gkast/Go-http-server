package handlers

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	contentHtml := `<h1>Hello User!!!</h1>`

	w.Header().Set("Content-Type", "text/html")

	if err := renderHTML(w, contentHtml); err != nil {
		HandleError(w, err)
		return
	}
}
