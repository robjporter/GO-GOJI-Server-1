package controllers

import(
	"fmt"
	"net/http"
	//"../../../render"
	//"../../../system"
	"github.com/zenazn/goji/web"
)

var guess string

func TMP( c web.C, w http.ResponseWriter, req *http.Request ) {
	guess = "SET"
}

func TMP2( c web.C, w http.ResponseWriter, req *http.Request ) {
	fmt.Println( guess )
}

func NotFound( c web.C, w http.ResponseWriter, req *http.Request ) {
	fmt.Println( "HERE" )
}