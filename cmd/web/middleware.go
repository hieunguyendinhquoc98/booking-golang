package main

import (
	"github.com/hieunguyendinhquoc98/bookings/internal/helpers"
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !helpers.IsAuthenticated(request) {
			session.Put(request.Context(), "error", "Logged in first!")
			http.Redirect(writer, request, "/user/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
