package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	"./middleware/logging2"
	"./routes"
	"./system"
	"github.com/Sirupsen/logrus"
	"github.com/golang/glog"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	gojiweb "github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

//openssl req -newkey rsa:2048 -new -nodes -keyout key.pem -out csr.pem -days 3650 -subj "/C=GB/ST=London/L=Earth/O=RJPDesigns/OU=IT/CN=NA/emailAddress=NA"
//openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem -days 3650 -subj "/C=GB/ST=London/L=Earth/O=RJPDesigns/OU=IT/CN=NA/emailAddress=NA"

func main() {
	// User all cpu cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	filename := flag.String("config", "conf/config.json", "Path to configuration file")
	flag.Parse()
	glog.Info("Prepare to repel boarders")

	// Any defer functions we need to finish up
	defer glog.Flush()

	// Setup main Application
	var application = &system.Application{}
	application.ClearTerminal()
	application.Init(filename)
	application.Log.WithFields(logrus.Fields{"animal": "walrus"}).Info("A walrus appears")
	application.DisplayIntro()
	application.SetupJWT()
	if application.Configuration.Database.Type == "ledis" {
		application.ConnectToLedisDatabase()
	} else if application.Configuration.Database.Type == "bolt" {
		application.ConnectToBoltDatabase()
	} else if application.Configuration.Database.Type == "mongo" {
		application.ConnectToMongoDatabase()
	}
	application.ReadStyleTemplates() // Replaces application.LoadTemplates()

	// Setup main static path to route asset folder
	static := gojiweb.New()
	static.Get("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir(application.Configuration.PublicPath))))
	http.Handle("/assets/", static)

	// Setup asset path for each of the installed plugins
	var plugins []system.Module
	plugins = application.GetPlugins()
	for i := 0; i < len(plugins); i++ {
		mod := plugins[i]
		tmp := gojiweb.New()
		tmp.Get("/assets/"+mod.Name+"/*", http.StripPrefix("/assets/"+mod.Name, http.FileServer(http.Dir(mod.Path+"/public"))))
		http.Handle("/assets/"+mod.Name+"/", tmp)
	}

	//application.Log.Debug("Useful debugging information.")
	//application.Log.Info("Something noteworthy happened!")
	//application.Log.Warn("You should probably take a look at this.")
	//application.Log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	// Calls panic() after logging
	//log.Panic("I'm bailing.")

	// Remove Logger Middleware so we can use our own
	goji.Abandon(middleware.Logger)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)
	goji.Use(application.ApplyDatabase)
	goji.Use(application.ApplySettings)
	goji.Use(application.ApplyConfiguration)
	goji.Use(application.ApplyIsXhr)
	goji.Use(application.ApplyLog)
	//goji.Use( application.ApplyCsrfProtection )
	goji.Use(logging2.LoggingMiddleWare)
	goji.Use(application.Settings.Stats.Handler)

	// Include all plugin routes
	routes.Include()

	// Any core routes
	//goji.Get( "/robots", controllers.Robots )
	//goji.Get( "/favicon", controllers.FavIcon )

	// Ensure we are shutting down gracefully
	graceful.PostHook(func() {
		application.Close()
	})

	// Set any parameters
	flag.Set("bind", ":"+application.Configuration.Host.Port)

	// Start server
	//goji.Serve()

	// Make TLS config

	conf, err := application.SetupSSLCertificate()
	if err == nil {
		fmt.Println("SERVING HTTPS")
		goji.ServeTLS(&conf)
	} else {
		fmt.Println("SERVING HTTP")
		goji.Serve()
	}
}
