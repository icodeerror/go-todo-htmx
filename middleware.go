package main

import "net/http"

func bypassLocalTunnel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("bypass-tunnel-reminder", "perfect")
		w.Header().Set("User-Agent", "localtunnel")

		next.ServeHTTP(w, r)

	})
}

func restrictDirectAccessHTMX(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		_, ok := headers["Hx-Request"]
		if !ok {
			// w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			// renderTemplate(w, "index.page.html", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
