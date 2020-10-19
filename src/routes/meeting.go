package routes

import (
	"fmt"
	"net/http"
	"server/controllers"
)

/*AddTestRoutes registers the test routes to the Mux*/
func AddTestRoutes(r *http.ServeMux) {

	fmt.Println("Hello from meeting.go")

	var createNewMeet controllers.CreateMeetingAPI
	r.Handle("/meetings", createNewMeet)

}
