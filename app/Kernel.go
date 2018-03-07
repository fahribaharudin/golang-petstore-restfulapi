package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Kernel is the main container to wrap the application
type Kernel struct {
	Router     RouterEngine
	Middleware MiddlewareEngine
}

// RouterEngine as container for the router engine (gorilla mux)
type RouterEngine struct {
	MuxRouter *mux.Router
}

// MiddlewareEngine as container for the middleware engine (negroni)
type MiddlewareEngine struct {
	Engine *negroni.Negroni
}

// kernel is the main app Object
var kernel Kernel

func init() {
	kernel = Kernel{
		Router:     RouterEngine{},
		Middleware: MiddlewareEngine{},
	}
}

// Run the app!
func (app *Kernel) Run() {
	kernel.Router.Register()
	kernel.Middleware.Register()

	s := &http.Server{
		Addr:         ":3000",
		Handler:      kernel.Middleware.Engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server running in http://localhost:3000")
	s.ListenAndServe()
}
