package system

import (
	"html/template"

	"../middleware/stats"
	"../modules/auth/jwt"
	"../modules/auth/store"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/gorilla/sessions"
	"github.com/siddontang/ledisdb/ledis"
	mgo "gopkg.in/mgo.v2"
)

type Module struct {
	Path string
	Name string
}

type CsrfProtection struct {
	Key    string
	Cookie string
	Header string
	Secure bool
}

type Settings struct {
	Count int
	Stats *stats.Stats
}

type BoltStruct struct {
	DB          *bolt.DB
	UserDBStore *store.BoltStore
}

type LedisStruct struct {
	DB          *ledis.Ledis
	UserDBStore *store.LedisStore
}

type MongoStruct struct {
	DB          *mgo.Database
	UserDBStore *store.MongoStore
}

type Application struct {
	Configuration  *Configuration
	Template       *template.Template
	Store          *sessions.CookieStore
	Ledis          LedisStruct
	Mongo          MongoStruct
	Bolt           BoltStruct
	DBOptions      *jwt.Options
	Settings       *Settings
	Log            *logrus.Logger
	CsrfProtection *CsrfProtection
}
