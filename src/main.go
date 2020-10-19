/*This is the hello world file*/
package main

import (
	"fmt"
	"net/http"
	"server/routes"
	"server/services"
)

func main() {
	fmt.Println("Hello world")

	services.ConnectMongoDB()

	routes.RegisterRoutes()
	http.ListenAndServe(":3002", routes.Mux)
}
