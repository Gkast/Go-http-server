package handlers

import (
	"fmt"
	"http-server/errors"
	"http-server/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `<h1>Welcome to My Dynamic Go Server</h1>
			<p>This is a dynamic Go web server serving an HTML page</p>`
	w.Header().Set("Content-Type", "text/html")
	if err := renderHTML(w, htmlContent); err != nil {
		HandleError(w, err)
		return
	}
}

func renderHTML(w http.ResponseWriter, content string) error {
	_, err := fmt.Fprint(w, template.HtmlTemplate(content))
	if err != nil {
		return errors.Wrap(err, "failed to render HTML")
	}
	return nil
}
