package routes

import (
	"gyme_backend/controllers"
	"net/http"
)
func RegisterRoutes() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/users", controllers.UserController)
}
