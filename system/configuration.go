package system

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigurationStats struct {
	Threshold int `json:"threshold"`
}

type ConfigurationServer struct {
	CDN string `json:"cdn"`
}

type ConfigurationDatabase struct {
	Hosts    string `json:"hosts"`
	Database string `json:"database"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Bolt     ConfigurationDatabaseBolt
	Ledis    ConfigurationDatabaseLedis
	Mongo    ConfigurationDatabaseMongo
}

type ConfigurationDatabaseBolt struct {
	Path string `json:"path"`
}

type ConfigurationDatabaseMongo struct {
	Monotonic  bool   `json:"monotonic"`
	DBName     string `json:"name"`
	UsersTable string `json:"users"`
}

type ConfigurationDatabaseLedis struct {
	Path        string `json:"path"`
	Count       int    `json:"count"`
	Backend     string `json:"backend"`
	readbuffer  int    `json:"readbuffer"`
	writebuffer int    `json:"writebuffer"`
}

type ConfigurationHost struct {
	Port  string `json:"port"`
	Https bool   `json:"https"`
}

type ConfigurationAdmin struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Scopes   []string `json:"scopes"`
}

type ConfigurationCertificate struct {
	Private       string `json:"private"`
	Public        string `json:"public"`
	CertPem       string `json:"certificate"`
	CertKey       string `json:"certkey"`
	Timeout       string `json:"timeout"`
	SigningMethod string `json:"signing"`
}

type Configuration struct {
	Secret       string `json:"secret"`
	PublicPath   string `json:"public_path"`
	PluginPath   string `json:"plugin_path"`
	TemplatePath string `json:"template_path"`
	Style        string `json:"style"`
	Development  bool   `json:"development"`
	Debug        bool   `json:"debug"`
	Templates    []string
	Server       ConfigurationServer
	Database     ConfigurationDatabase
	Host         ConfigurationHost
	Stats        ConfigurationStats
	Admin        ConfigurationAdmin
	Certificate  ConfigurationCertificate
}

func (configuration *Configuration) Load(filename string) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = configuration.Parse(data)
	return
}

func (configuration *Configuration) Parse(data []byte) (err error) {
	err = json.Unmarshal(data, &configuration)
	return
}
