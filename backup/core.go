package system

import (
	//"encoding/gob"

	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorilla/sessions"

	"../modules/auth/jwt"
	"../modules/auth/store"

	"../middleware/stats"
	"github.com/Sirupsen/logrus"

	"github.com/golang/glog"
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

func (application *Application) Init(filename *string) {
	//gob.Register(bson.ObjectId(""))
	application.Configuration = &Configuration{}
	err := application.Configuration.Load(*filename)
	if err != nil {
		glog.Fatalf("Can't read configuration file: %s", err)
		panic(err)
	}
	application.Settings = &Settings{Count: 0}
	application.Settings.Stats = stats.New()
	application.Log = logrus.New()
	application.Log.Formatter = new(logrus.JSONFormatter)
	application.Log.Level = logrus.DebugLevel
	f, err := os.OpenFile("./logs/testlogrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	application.Log.Out = f
	application.Store = sessions.NewCookieStore([]byte(application.Configuration.Secret))
}

func (application *Application) SetupJWT() {
	timeout, _ := strconv.Atoi(application.Configuration.Certificate.Timeout)
	options := jwt.Options{
		SigningMethod: application.Configuration.Certificate.SigningMethod,
		PrivateKey:    readFile(application.Configuration.Certificate.Private), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     readFile(application.Configuration.Certificate.Public),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    time.Duration(timeout) * time.Minute,
	}
	application.DBOptions = &options
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (application *Application) LoadTemplates() error {
	var templates []string
	fn := func(path string, f os.FileInfo, err error) error {

		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			fmt.Println(path)
			templates = append(templates, path)
		}
		return nil
	}

	var err error
	if exists(application.Configuration.TemplatePath) {
		err = filepath.Walk(application.Configuration.TemplatePath, fn)
	} else {
		panic("TEMPLATE DIRECTORY DOES NOT EXIST")
		os.Exit(1)
	}

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}

func (application *Application) ConnectToLedisDatabase() {
	create := false

	//TEMP
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)

	application.DB = db
}

func (application *Application) ConnectToBoltDatabase() {
	create := false
	if !Exists("data/usersdb") {
		create = true
	}
	db, err := bolt.Open("data/usersdb", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		glog.Fatalf("Can't connect to or create the database: %v", err)
		panic(err)
	}
	application.DB = db
	store2, err2 := store.NewBoltStore(application.DB, "users")
	if err2 != nil {
		panic("Can not create bolt store")
	}
	application.UserDBSession = store2
	if create {
		application.createDefaultAdminUser()
	}
}

func (application *Application) createDefaultAdminUser() {
	application.UserDBSession.Signin(application.Configuration.Admin.Email,
		application.Configuration.Admin.Password,
		application.Configuration.Admin.Scopes)
}

func (application *Application) Close() {
	glog.Info("Bye!")
	application.DB.Close()
}
