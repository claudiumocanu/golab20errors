package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (chiRouter) GET(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (chiRouter) POST(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (chiRouter) DELETE(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Delete(uri, f)
}
func (chiRouter) SERVE(port string) {
	fmt.Printf("Server running chiRouter on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
