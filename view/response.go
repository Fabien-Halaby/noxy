package view

import (
	"net/http"
	"noxy/model"
)

func SendResponse(w http.ResponseWriter, entry model.CacheEntry, fromCache bool) {
	for key, values := range entry.Headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	if fromCache {
		w.Header().Set("X-Cache", "HIT")
	} else {
		w.Header().Set("X-Cache", "MISS")
	}

	w.WriteHeader(entry.StatusCode)
	w.Write(entry.Body)
}
