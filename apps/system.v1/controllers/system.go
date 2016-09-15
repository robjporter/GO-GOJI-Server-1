package controllers

import(
	"fmt"
	"net/http"
	"time"
	"../../../render"
	//"../../../system"
	"github.com/zenazn/goji/web"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
)

func SysHome( c web.C, w http.ResponseWriter, req *http.Request ) {
}

func SysCoreRaw( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	if len( tmp ) == 1 {
		if tmp[ 0 ] == "application/json" {
    		render.RenderJSON( w, http.StatusOK, "AllData: Data" )
		} else if tmp[ 0 ] == "application/xml" {
			render.RenderXML( w, http.StatusOK, "AllData: Data" )
		} else {
			fmt.Println( "NOT SURE WHERE WE ARE!" )
		}
	}
}

func SysCoreCount( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUCounts( true )
	if err != nil { tmp2 = 0 }
	result := `{ "CPU" : { "CORES" : { "COUNT" : "` + strconv.Itoa( tmp2 ) + `"}}}`
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

func SysCPUPercent1Sec( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUPercent( time.Second, true )
	if err != nil {
		panic( "No return while checking CPU" )
		return
	}

	result := `{ "CPU" : { "USAGE" : {`
	for i := 0; i < len( tmp2 ); i++ {
		result += `"CORE" : "` + strconv.Itoa( int( tmp2[ i ])) + `"`
		if i + 1 != len( tmp2 ) { result +=  `,` }
	}
	result += "}}}"
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

func SysCPUPercent1SecTot( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUPercent( time.Second, false )
	if err != nil {
		panic( "No return while checking CPU" )
		return
	}

	result := `{ "CPU" : { "USAGE" : {`
	for i := 0; i < len( tmp2 ); i++ {
		result += `"CORE" : "` + strconv.Itoa( int( tmp2[ i ])) + `"`
		if i + 1 != len( tmp2 ) { result +=  `,` }
	}

	result += "}}}"
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

func SysCPUInfo( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	//result := `{ "CPU" : { "INFO" : ` + tmp3 + `}}`
	result := tmp2[ 0 ]
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

func SysCPUInfoVendor( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "VENDOR_ID" : "` + tmp2[ 0 ].VendorID + `"}}}`
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

func SysCPUInfoFamily( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "FAMILY" : "` + tmp2[ 0 ].Family + `"}}}`
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

func SysCPUInfoModel( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "MODEL" : "` + tmp2[ 0 ].Model + `"}}}`
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

func SysCPUInfoCores( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "CORES" : "` + strconv.FormatInt( int64( tmp2[ 0 ].Cores ), 10 ) + `"}}}`
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

func SysCPUInfoName( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "NAME" : "` + tmp2[ 0 ].ModelName + `"}}}`
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

func SysCPUInfoSpeed( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "SPEED_Mhz" : "` + strconv.FormatInt( int64( tmp2[ 0 ].Mhz), 10 ) + `"}}}`
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

func SysCPUInfoCache( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}
	result := `{ "CPU" : { "INFO" : { "CACHE" : "` + strconv.FormatInt( int64( tmp2[ 0 ].CacheSize ), 10 ) + `"}}}`
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

func SysCPUInfoFlags( c web.C, w http.ResponseWriter, req *http.Request ){
	tmp := req.Header[ "Accept" ]
	tmp2, err := cpu.CPUInfo()
	if err != nil { 
		panic( "No CPU Info received." )
		return
	}

	result := `{ "CPU" : { "INFO" : { "FLAGS" : {`
	flags := tmp2[ 0 ].Flags
	for i := 0; i < len( flags ); i++ {
		result += `"FLAG` + strconv.Itoa( i + 1 ) + `" : "` + flags[ i ] + `"`
		if i + 1 != len( flags ) { result +=  `,` }
	}
	result += "}}}}"

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

func SysRamCount( c web.C, w http.ResponseWriter, req *http.Request ) {
}