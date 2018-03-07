package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Register routes
func (r *RouterEngine) Register() {
	// init gorilla mux as router engine
	r.MuxRouter = mux.NewRouter()

	r.MuxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the petstore api.."))
	})

	r.MuxRouter.HandleFunc("/pet", container.PetController.Store).Methods("POST")
}
