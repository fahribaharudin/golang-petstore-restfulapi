package app

import (
	"net/http"

	"github.com/urfave/negroni"
)

// Register middleware
func (m *MiddlewareEngine) Register() {
	// init negroni classic middleware as default middleware engine
	m.Engine = negroni.Classic()

	// example of custom middleware
	m.Engine.Use(negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// do some stuff before
		next(w, r)
		// do some stuff after
	}))

	m.Engine.UseHandler(kernel.Router.MuxRouter)
}
