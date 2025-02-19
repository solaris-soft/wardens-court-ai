package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// HTTPHandler is a modified HandlerFunc to also provide a log error if error executing the route
type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

// MakeRoute encapsulates error logging to the http.HandlerFunc function signature (adapter pattern)
func MakeRoute(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

// A function to render the HTML templates by using templ
// This is a wrapper around the templ.Component.Render method
func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
