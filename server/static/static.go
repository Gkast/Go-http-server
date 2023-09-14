package static

import (
	"http-server/config" // Import your configuration package
	"net/http"
)

// ServeStaticFiles serves static files from the configured directory
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()

	// Get the file server for the static directory
	fs := http.FileServer(http.Dir(cfg.StaticDir))

	// Serve the static files
	fs.ServeHTTP(w, r)
}
