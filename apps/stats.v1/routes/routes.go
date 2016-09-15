package routes

import (
	"../controllers"
	"net/http"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Routes( m *web.Mux ) {
	goji.Handle( "/api/v1/stats/*", m )
	goji.Get( "/api/v1/stats", http.RedirectHandler( "/api/v1/stats/", 301 ))

	m.Get( "/api/v1/stats/", controllers.StatsHome )
	m.Get( "/api/v1/stats/dashboard", controllers.StatsDashboard )
	m.Get( "/api/v1/stats/raw", controllers.StatsRaw )
	m.Get( "/api/v1/stats/averageresponsetime", controllers.StatsAverageResponseTime )
	m.Get( "/api/v1/stats/methodtypes", controllers.StatsRequestTypeCounts )
	m.Get( "/api/v1/stats/uptime", controllers.StatsUpTime )
	m.Get( "/api/v1/stats/time", controllers.StatsTime )
	m.Get( "/api/v1/stats/count", controllers.StatsCount )
	m.Get( "/api/v1/stats/calls", controllers.StatsCalls )
	m.Get( "/api/v1/stats/requests", controllers.StatsRequests )
	m.Get( "/api/v1/stats/codes", controllers.StatsCodes )
}