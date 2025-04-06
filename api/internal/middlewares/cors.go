package middlewares

import (
    "github.com/gorilla/handlers"
    "net/http"
)

func ConfigureCORS(handler http.Handler) http.Handler {
    return handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(handler)
}