package laser

import (
	"net/http"
)

type Route struct {
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) GET(path string, h http.HandlerFunc) {
	r.bindMethod("GET", path, h)
}

func (r *Route) bindMethod(method string, path string, h http.HandlerFunc) {

}
