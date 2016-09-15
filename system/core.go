package system

import (
	//"encoding/gob"

	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"

	"../modules/auth/jwt"
	"../modules/auth/store"

	"../middleware/stats"
	"github.com/Sirupsen/logrus"

	"github.com/boltdb/bolt"
	"github.com/golang/glog"
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
	mgo "gopkg.in/mgo.v2"
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

func (application *Application) SetupSSLCertificate2() (tls.Config, error) {
	var config tls.Config
	roots := x509.NewCertPool()

	pem := readFile(application.Configuration.Certificate.CertPem)
	//key := readFile(application.Configuration.Certificate.CertKey)
	ok := roots.AppendCertsFromPEM([]byte(pem))

	if !ok {
		panic("failed to parse root certificate")
	}

	config = tls.Config{
		RootCAs:            roots,
		InsecureSkipVerify: true,
		ServerName:         "localhost",
	}

	return config, errors.New("")
}

func (application *Application) SetupSSLCertificate() (tls.Config, error) {
	var config tls.Config

	pem := readFile(application.Configuration.Certificate.CertPem)
	key := readFile(application.Configuration.Certificate.CertKey)

	cert, err := tls.LoadX509KeyPair(pem, key)
	fmt.Println(err)

	if err != nil {
		return config, errors.New("Can't load keypair: " + err.Error())
	}

	config = tls.Config{
		ClientAuth:         tls.RequireAndVerifyClientCert,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		ServerName:         "localhost",
		IsCA:               true,
	}

	return config, errors.New("")
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

func (application *Application) ConnectToMongoDatabase() {
	session, err := mgo.Dial(application.Configuration.Database.Hosts)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, application.Configuration.Database.Mongo.Monotonic)
	application.Mongo.UserDBStore = store.NewMongoStore(session.DB(application.Configuration.Database.Mongo.DBName).C(application.Configuration.Database.Mongo.UsersTable))

	_, err = application.createDefaultAdminUser()
	if err == nil {

	}
}

func (application *Application) ConnectToLedisDatabase() {
	//create := false

	//TEMP
	cfg := lediscfg.NewConfigDefault()
	cfg.DBPath = application.Configuration.Database.Ledis.Path
	cfg.Databases = application.Configuration.Database.Ledis.Count
	cfg.DBName = application.Configuration.Database.Ledis.Backend
	cfg.ConnReadBufferSize = application.Configuration.Database.Ledis.readbuffer
	cfg.ConnWriteBufferSize = application.Configuration.Database.Ledis.writebuffer
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)

	application.Ledis.DB = l
	application.Ledis.UserDBStore = store.NewLedisStore(db)

	_, err := application.createDefaultAdminUser()
	if err == nil {

	}
}

func (application *Application) ConnectToBoltDatabase() {
	create := false
	path := application.Configuration.Database.Bolt.Path + "/" + application.Configuration.Database.Name
	if !Exists(path) {
		create = true
	}
	db, err := bolt.Open(path, 0600, &bolt.Options{})
	if err != nil {
		glog.Fatalf("Can't connect to or create the database: %v", err)
		panic(err)
	}
	application.Bolt.DB = db
	store2, err2 := store.NewBoltStore(application.Bolt.DB, "users")
	application.Bolt.UserDBStore = store2
	if err2 != nil {
		panic("Can not create bolt store")
	}
	if create {
		application.createDefaultAdminUser()
	}
}

func (application *Application) createDefaultAdminUser() (string, error) {
	if application.Configuration.Database.Type == "bolt" {
		return application.Bolt.UserDBStore.Signin(application.Configuration.Admin.Email, application.Configuration.Admin.Password, application.Configuration.Admin.Scopes)
	} else if application.Configuration.Database.Type == "ledis" {
		found, email, err := application.Ledis.UserDBStore.AdminExists(application.Configuration.Admin.Email)
		if found {
			return email, err
		} else {
			return application.Ledis.UserDBStore.Signin(application.Configuration.Admin.Email, application.Configuration.Admin.Password, application.Configuration.Admin.Scopes)
		}
	} else if application.Configuration.Database.Type == "mongo" {
		found, email, err := application.Mongo.UserDBStore.AdminExists(application.Configuration.Admin.Email)
		if found {
			return email, err
		} else {
			return application.Mongo.UserDBStore.Signin(application.Configuration.Admin.Email, application.Configuration.Admin.Password, application.Configuration.Admin.Scopes)
		}
	}
	return "", errors.New("")
}

func (application *Application) Close() {
	glog.Info("Bye!")
	if application.Configuration.Database.Type == "bolt" {
		application.Bolt.DB.Close()
	} else if application.Configuration.Database.Type == "ledis" {
		application.Ledis.DB.Close()
	} else if application.Configuration.Database.Type == "mongo" {

	}
}
