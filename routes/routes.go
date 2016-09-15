package routes

import (
	//bookmarkRoutes "../apps/bookmarksv1/routes"
	authRoutes "../apps/auth.v1/routes"
	bingRoutes "../apps/bing.v1/routes"
	coreRoutes "../apps/core.v1/routes"
	funRoutes "../apps/fun.v1/routes"
	sparkRoutes "../apps/spark.v1/routes"
	statsRoutes "../apps/stats.v1/routes"
	systemRoutes "../apps/system.v1/routes"
	"github.com/zenazn/goji/web"
)

func Include() {
	// Stats app
	stats := web.New()
	statsRoutes.Routes(stats)

	// Stats app
	sys := web.New()
	systemRoutes.Routes(sys)

	// Fun app
	fun := web.New()
	funRoutes.Routes(fun)

	//Bing app
	bing := web.New()
	bingRoutes.Routes(bing)

	//Spark app
	spark := web.New()
	sparkRoutes.Routes(spark)

	// Core app
	auth := web.New()
	authRoutes.Routes(auth)

	// Core app
	core := web.New()
	coreRoutes.Routes(core)
}
