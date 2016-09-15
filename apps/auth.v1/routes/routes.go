package routes

import (
	"net/http"

	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Prefetch() {

}

func Routes(m *web.Mux) {
	goji.Handle("/api/v1/auth/*", m)
	goji.Get("/api/v1/auth", http.RedirectHandler("/api/v1/auth/", 301))

	m.Post("/api/v1/auth/login", controllers.Login)
	m.Post("/api/v1/auth/impersonate", controllers.Impersonate)
	m.Get("/api/v1/auth/validate", controllers.Validate)

	m.Get("/api/v1/auth/signup", controllers.Signup)
}
