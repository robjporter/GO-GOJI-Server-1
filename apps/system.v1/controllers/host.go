package controllers

import(
	"fmt"
	"net/http"
	"time"
	"strconv"
	"../../../render"
	//"../../../system"
	"github.com/zenazn/goji/web"
	"github.com/shirou/gopsutil/host"
)

func HostBootTime( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, err := host.BootTime()
	if err != nil { tmp2 = 0 }

	tmp3 := time.Unix( int64( tmp2 ), 0).Format( "02/01/2006 - 15:04:05" )

	result := `{ "HOST" : { "BOOTTIME" : "` + tmp3 + `"}}`
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

func HostVersion( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := host.HostInfo()

	result := `{ "HOST" : { "KERNEL" : { "VERSION": "` + tmp2.PlatformVersion + `"}}}`
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

func HostOS( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _, _, _ := host.GetPlatformInformation()

	result := `{ "HOST" : { "OPERATING_SYSTEM" : "` + tmp2 + `"}}`
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

func HostActiveUsers( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := host.Users()

	result := `{ "HOST" : { "USERS" : {`
	for i := 0; i < len( tmp2 ); i++ {
		result += `"USER" : {`
		result += `"USERNAME" : "` + tmp2[ i ].User + `",`
		result += `"TERMINAL" : "` + tmp2[ i ].Terminal + `",`
		result += `"HOST" : "` + tmp2[ i ].Host + `",`
		result += `"STARTED" : "` + strconv.Itoa( tmp2[ i ].Started ) + `"`
		result += `}`
		if i + 1 != len( tmp2 ) { result += `,` }
	}
	result += `}}}`

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

func HostVirtualisationSystem( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp9, _ := host.HostInfo()

	result := `{ "HOST" : { "VIRTUALISATION" : { "SYSTEM" : "` + tmp9.VirtualizationSystem + `"}}}`
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

func HostVirtualisationRole( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp9, _ := host.HostInfo()

	result := `{ "HOST" : { "VIRTUALISATION" : { "ROLE" : "` + tmp9.VirtualizationRole + `"}}}`
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
