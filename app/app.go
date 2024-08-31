package app

import (
	"fmt"
	"net/http"
	"os"

	"gyme_backend/routes"
)

func Run() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    // Register the routes
    routes.RegisterRoutes()

    fmt.Printf("Server is running on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
        os.Exit(1)
    }
}
