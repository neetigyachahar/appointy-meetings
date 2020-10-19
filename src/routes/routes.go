package routes

import (
	"net/http"
)

/*Mux is the root of the routes*/
var Mux *http.ServeMux = http.NewServeMux()

/*RegisterRoutes registers the routes*/
func RegisterRoutes() {
	AddTestRoutes(Mux)
}
