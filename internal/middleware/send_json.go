package middleware

import (
	"encoding/json"
	"net/http"
)

func SendJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		jsonObj := r.Context().Value("json")
		if jsonObj == nil {
			return
		}

		json, errJson := json.Marshal(jsonObj)
		if errJson != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})
}
