package controllers

import (
	"fmt"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to GymE Backend!")
}

func UserController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User route accessed")
	// Add your user-related logic here
}
