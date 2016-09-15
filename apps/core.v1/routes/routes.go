package routes

import (
	"../controllers"
	"net/http"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Routes( m *web.Mux ) {
	goji.Handle( "/api/v1/core/*", m )
	goji.Get( "/api/v1/core", http.RedirectHandler( "/api/v1/core/", 301 ))

	m.Get( "/api/v1/core/", controllers.CoreHome )
	m.Get( "/api/v1/core/home", controllers.CoreHome )
	m.Get( "/api/v1/core/about", controllers.CoreAbout )
	goji.NotFound( controllers.NotFound )
}