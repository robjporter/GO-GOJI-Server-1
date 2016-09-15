package controllers

import(
	"fmt"
	"net/http"
	//"strconv"
	"../../../render"
	//"../../../system"
	"github.com/zenazn/goji/web"
	"github.com/shirou/gopsutil/load"
)

func LoadInfoSummary( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := load.LoadAvg()

	result := tmp2 //`{ "HOST" : { "NET" : "` + tmp3 + `"}}`
	fmt.Println( tmp2 )
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, result )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}