package routes

import (
	"net/http"

	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Routes(m *web.Mux) {
	goji.Handle("/api/v1/bing/*", m)
	goji.Get("/api/v1/bing", http.RedirectHandler("/api/v1/bing/", 301))

	m.Get("/api/v1/bing/", controllers.BingHome)
	m.Get("/api/v1/bing/home", controllers.BingHome)
	m.Get("/api/v1/bing/about", controllers.BingAbout)
	m.Get("/api/v1/bing/daily/photo", controllers.BingDailyPhoto)
	m.Get("/api/v1/bing/daily/photo/embed", controllers.BingDailyPhotoEmbed)
	m.Get("/api/v1/bing/daily/photo/raw", controllers.BingDailyPhotoRaw)
	
	m.Get("/api/v1/bing/daily/photo/svga", controllers.BingDailyPhotoSVGA)    //800x600
	m.Get("/api/v1/bing/daily/photo/xga", controllers.BingDailyPhotoXGA)      //1024x768
	m.Get("/api/v1/bing/daily/photo/wxga", controllers.BingDailyPhotoWXGA)    //1280x720
	m.Get("/api/v1/bing/daily/photo/hd", controllers.BingDailyPhotoHD)        //1366x768
	m.Get("/api/v1/bing/daily/photo/fhd", controllers.BingDailyPhotoFHD)      //1920x1080
	m.Get("/api/v1/bing/daily/photo/qhd", controllers.BingDailyPhotoQHD)      //2560x1440
	m.Get("/api/v1/bing/daily/photo/wqxga", controllers.BingDailyPhotoWQXGA)  //2560x1600
	m.Get("/api/v1/bing/daily/photo/uhd", controllers.BingDailyPhotoUHD)      //3840x2160
	
	m.Get("/api/v1/bing/daily/photo/sized/:x/:y", controllers.BingDailyPhotoSized)
}
