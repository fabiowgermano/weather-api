package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []ItemHandler
	WebServerPort string
}

type ItemHandler struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []ItemHandler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	itemHandler := ItemHandler{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	s.Handlers = append(s.Handlers, itemHandler)
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		if handler.Method == http.MethodGet {
			s.Router.Get(handler.Path, handler.Handler)
		} else if handler.Method == http.MethodPost {
			s.Router.Post(handler.Path, handler.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
