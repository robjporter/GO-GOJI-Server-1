package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"../../../render"
	"../../../system"
	"../../../middleware/stats"
	"github.com/zenazn/goji/web"
)

func StatsHome( c web.C, w http.ResponseWriter, req *http.Request ) {
	templates := render.GetBaseTemplates( c )
	templates = append( templates, "apps/stats.v1/views/home.html" )
	err := render.RenderTemplate( w, templates, "base", map[string]string{ "Title": "Home" })
	if err != nil {
		http.Error( w, err.Error(), http.StatusInternalServerError )
	}
}
func StatsDashboard( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	values := map[string]string{}
	templates := []string{}
	templates = append( templates, "apps/stats.v1/views/dashboard.html" )

	week, day, hour, min, sec := system.SecondsToDate( settings.Stats.Data().UpTimeSec )
	tmp2 := settings.Stats.Data().TotalStatusCodeCount
	tmp3 := settings.Stats.Data().RequestTypeCounts
	var tmp5 *stats.ResponseURL = settings.Stats.Data().MaxResponseTime
	tmp6 := ""

	values[ "Title" ] = "Home"

	if week != 0 {
		tmp6 = strconv.Itoa( week ) + " weeks "
		tmp6 += strconv.Itoa( day ) + " days "
	} else if day != 0 {
		tmp6 += strconv.Itoa( day ) + " days "
		tmp6 += strconv.Itoa( hour ) + " hours "
	} else if hour != 0 {
		tmp6 += strconv.Itoa( hour ) + " hours "
		tmp6 += strconv.Itoa( min ) + " mins "
	} else {
		tmp6 += strconv.Itoa( min ) + " mins "
		tmp6 += strconv.Itoa( sec ) + " secs "
	}

	values[ "ActiveTimes" ] = tmp6
	values[ "Requests" ] = strconv.Itoa( settings.Stats.Data().TotalCount )
	values[ "AverageResponse" ] = settings.Stats.Data().AverageResponseTime
    values[ "LongestRequest" ] = tmp5.ResponseDuration.String()
    values[ "GetRequests" ] = strconv.Itoa( tmp3[ "GET" ])
    values[ "PostRequests" ] = strconv.Itoa( tmp3[ "POST" ])
    values[ "HTTP200" ] = strconv.Itoa( tmp2[ "200" ])
    values[ "HTTP404" ] = strconv.Itoa( tmp2[ "404" ])
    values[ "HTTP500" ] = strconv.Itoa( tmp2[ "500" ])

	err := render.RenderTemplate( w, templates, "base", values )
	if err != nil {
		http.Error( w, err.Error(), http.StatusInternalServerError )
	}
}
func StatsRaw( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		render.RenderJSON( w, http.StatusOK, settings.Stats.Data() )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data() )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}


func StatsCalls( c web.C, w http.ResponseWriter, req *http.Request ) {}

func StatsRequestTypeCounts( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		result := `{ "STATS" : { "RESPONSE" : { "METHODCOUNTS" : `
    		for key, value := range settings.Stats.Data().RequestTypeCounts {
    			result += `{ "` + key + `" : "` + strconv.Itoa( value ) + `""}`
			}
    		result += `}}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data() )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}

func StatsCodes( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		render.RenderJSON( w, http.StatusOK, settings.Stats.Data() )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data() )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}

func StatsAverageResponseTime( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
			result := `{ "STATS" : { "RESPONSE" : { "AVERAGERESPONSETIME" : "` + settings.Stats.Data().AverageResponseTime + `" }}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data().AverageResponseTime )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}

func StatsUpTime( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
			result := `{ "STATS" : { "RESPONSE" : { "UPTIME" : "` + settings.Stats.Data().UpTime + `" }}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data().UpTime )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}
func StatsTime( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
			result := `{ "STATS" : { "RESPONSE" : { "SERVERTIME" : "` + settings.Stats.Data().Time + `" }}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data().Time )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}
func StatsCount( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		result := `{ "STATS" : { "RESPONSE" : { "STATSCOUNT" : `
    		for key, value := range settings.Stats.Data().TotalStatusCodeCount {
    			result += `{ "` + key + `" : "` + strconv.Itoa( value ) + `""}`
			}
    		result += `}}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, "A" )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}
func StatsRequests( c web.C, w http.ResponseWriter, req *http.Request ) {
	settings := c.Env[ "Settings" ].(*system.Settings)
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		result := `{ "STATS" : { "RESPONSE" : { "STATSREQUESTS" : "` + strconv.Itoa( settings.Stats.Data().TotalCount ) + `" }}}`
    		render.RenderJSON( w, http.StatusOK, result )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, settings.Stats.Data().TotalCount )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}
