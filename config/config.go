package config

import (
	"log"
	"net/http"
	"strconv"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

// Config stores server configuration settings
type options struct {
	addr   string
	port   string
	routes []Route
}

type Option func(*options)

func WithAddr(addr string) Option {
	return func(o *options) {
		o.addr = addr
	}
}

func WithPort(port string) Option {
	return func(o *options) {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			o.port = "8080"
		} else if portInt < 0 || portInt > 65535 {
			o.port = "8080"
		} else {
			o.port = port
		}
	}
}

func WithRoutes(routes *[]Route) Option {
	return func(o *options) {
		if routes == nil {
			log.Fatal("routes is nil")
		}
		o.routes = *routes
	}
}

// NewAppConfig returns a new app configuration with default settings
func NewAppConfig(o ...Option) http.Server {

	// Add the options to the struct
	var options options
	for _, opt := range o {
		opt(&options)
	}

	mux := http.NewServeMux()
	// Add static file server to serve the public files
	// This is used to serve the CSS, JS, and other static files
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// Register the routes to the mux
	for _, route := range options.routes {
		mux.HandleFunc(route.Path, route.Handler)
	}

	return http.Server{
		Addr:    options.addr + ":" + options.port,
		Handler: mux,
	}
}
