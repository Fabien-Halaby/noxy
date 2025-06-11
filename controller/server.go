package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"noxy/model"
	"noxy/view"
)

func StartServer(cache *model.Cache, args Args) {
	if args.Port == 0 || args.Origin == "" {
		log.Fatal("Error: --port and --origin are required.")
	}

	originURL, err := url.Parse(args.Origin)
	if err != nil || (originURL.Scheme != "http" && originURL.Scheme != "https") {
		log.Fatalf("Error: --origin must be a valid url (e.g.: http://example.com), got %s", args.Origin)
	}

	client := &http.Client{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		cacheKey := r.URL.Path
		if entry, exists := cache.Get(cacheKey); exists {
			view.SendResponse(w, entry, true)

			return
		}

		entry, err := model.ProxyRequest(client, originURL, r)
		if err != nil {
			http.Error(w, "Error forwarding request: "+err.Error(), http.StatusBadGateway)

			return
		}

		cache.Set(cacheKey, entry)
		view.SendResponse(w, entry, false)
	}

	view.ShowStartMessage(args.Port, args.Origin)

	addr := fmt.Sprintf(":%d", args.Port)
	if err := http.ListenAndServe(addr, http.HandlerFunc(handler)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
