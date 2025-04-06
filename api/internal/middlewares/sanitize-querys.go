package middlewares

import (
	"net/http"
	"regexp"
	"slices"

	"github.com/Blessed0314/tru-test/api/internal/utils"
	"github.com/gorilla/mux"
)


func SanitizeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        var sqlInjectionPattern = regexp.MustCompile(`(?i)(\b(SELECT|INSERT|DELETE|UPDATE|DROP|UNION|--|;|')\b)`)

        for _, values := range r.URL.Query() {
            if slices.ContainsFunc(values, sqlInjectionPattern.MatchString) {
                    utils.SendResponse(w, r, http.StatusBadRequest, "Invalid input detected", nil)
                    return
                }
        }

		vars := mux.Vars(r)
        for _, param := range vars {
            if sqlInjectionPattern.MatchString(param) {
                utils.SendResponse(w, r, http.StatusBadRequest, "Invalid input detected", nil)
                return
            }
        }

        next.ServeHTTP(w, r)
    })
}