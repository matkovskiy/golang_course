package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// This type describes handler that can decorate another handler:
// decoratedHandler := Middleware(handler)
type Middleware func(fn http.HandlerFunc) http.HandlerFunc

// This function takes handler and a list of decorators and apply them
// one by one: ApplyMiddleware(handler, one, two) => two(one(handler))
func ApplyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, next := range middlewares {
		handler = next(handler)
	}
	return handler
}

// creates session and sets cookie for the authenticated user
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		next(w, req)

		// Wrapped handler should modify context
		// and set CONTEXT_AUTH_KEY key to true.
		// Otherwise we do not authenticate this request.
		// Also, next(w, req) should not write to the body
		// because SetCookie will write to header and it will be
		// useless
		isAuthenticated, exists := req.Context().Value(CONTEXT_AUTH_KEY).(bool)
		if !exists || !isAuthenticated {
			return
		}

		fmt.Println("Authenticate3")
		session := uuid.NewString()
		fmt.Println(session)

		authenticated[session] = true

		cookie := http.Cookie{
			Name:     SESSION_COOKIE,
			Value:    session,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteNoneMode,
		}

		http.SetCookie(w, &cookie)

		w.Write([]byte("you are logged in"))

	}
}

func Authorize() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {

			cookie, err := req.Cookie(SESSION_COOKIE)

			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					http.Error(w, "You are not logged in", http.StatusBadRequest)
				default:
					log.Println(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}

			if !authenticated[cookie.Value] {
				http.Error(w, "Your session is expired", http.StatusUnauthorized)
				return
			}

			next(w, req)
		}
	}
}
