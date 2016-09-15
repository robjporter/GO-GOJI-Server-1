package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"../../../modules/auth"
	"../../../modules/auth/jwt"
	"../../../modules/auth/store"
	"../../../render"
	"../../../system"
	"github.com/zenazn/goji/web"
)

func Validate(c web.C, w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Token")

	//conf := c.Env["Configuration"].(*system.Configuration)
	Options := c.Env["DBOptions"].(*jwt.Options)

	if succ, _ := auth.Authenticate(token, *Options); succ == true {
		fmt.Println("DONE")
	}
}

func Login(c web.C, w http.ResponseWriter, req *http.Request) {
	email := req.Header.Get("Email")
	pass := req.Header.Get("Password")
	accept := ""
	result := ""
	token := ""
	err := errors.New("")

	if len(req.Header["Accept"]) == 1 {
		accept = req.Header["Accept"][0]
	} else {
		accept = "application/json"
	}

	conf := c.Env["Configuration"].(*system.Configuration)
	DBType := conf.Database.Type
	Options := c.Env["DBOptions"].(*jwt.Options)

	if DBType == "bolt" {
		token, err = auth.LoginBolt(email, pass, c.Env["UserDBStore"].(*store.BoltStore), *Options)
	} else if DBType == "ledis" {
		token, err = auth.LoginLedis(email, pass, c.Env["UserDBStore"].(*store.LedisStore), *Options)
	} else if DBType == "mongo" {
		token, err = auth.LoginMongo(email, pass, c.Env["UserDBStore"].(*store.MongoStore), *Options)
	}

	if err == nil {
		timeout := conf.Certificate.Timeout
		result = `{"Login":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Duration":"` + timeout + `","Token":"` + token + `"}}`
	} else {
		result = `{ "CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}`
	}

	if accept == "application/json" {
		render.RenderJSON(w, http.StatusOK, result)
	} else if accept == "application/xml" {
		render.RenderXML(w, http.StatusOK, `<xml><token>`+token+`</token></xml>`)
	} else {
		fmt.Println("NOT SURE WHERE WE ARE!")
	}
}

func Impersonate(c web.C, w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Token")
	email := req.Header.Get("Email")
	accept := ""
	result := ""
	newToken := ""

	if len(req.Header["Accept"]) == 1 {
		accept = req.Header["Accept"][0]
	} else {
		accept = "application/json"
	}

	DBSession := c.Env["UserDBSession"].(*store.BoltStore)
	Options := c.Env["DBOptions"].(*jwt.Options)
	userId, err := auth.AuthenticateUserId(token, DBSession, *Options)
	cont := false

	fmt.Println(accept)

	if err == nil {
		valid, err2 := auth.ValidateScope(token, "IMPERSO", *Options)
		if err2 == nil {
			cont = valid
		} else {
			//result = system.GenerateResponse()
			result = ""
		}
	} else {
		result = ""
	}

	if cont {

	}

	if err == nil {
		fmt.Println(userId)
		newToken, err = auth.Impersonate(email, DBSession, *Options)
		if err == nil {
			conf := c.Env["Configuration"].(*system.Configuration)
			timeout := conf.Certificate.Timeout
			result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","User":"` + email + `","Duration":"` + timeout + `","Token":"` + newToken + `"}}`
			render.RenderJSON(w, http.StatusOK, result)
		} else {
			result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}}`
			render.RenderJSON(w, http.StatusUnauthorized, result)
		}
	} else {
		result = `{"Impersonate":{"CreatedAt":"` + strconv.Itoa(int(time.Now().Unix())) + `","Error":"` + err.Error() + `"}}`
		render.RenderJSON(w, http.StatusUnauthorized, result)
	}
}

func Signup(c web.C, w http.ResponseWriter, req *http.Request) {
	fmt.Println("Signup")
}

func RequestToJsonObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(jsonDoc)
	if err != nil {
		return err
	}
	return nil
}
