package controllers

import(
	"fmt"
	"net/http"
	"strconv"
	"../../../render"
	//"../../../system"
	"github.com/zenazn/goji/web"
	"github.com/shirou/gopsutil/net"
)

func NetInfoAdapterSummary( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := "TMP" //`{ "HOST" : { "NET" : "` + tmp3 + `"}}`
	fmt.Println( tmp3 )
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

func NetInfoAdapterSummaryBytesSent( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "SENT" : { "BYTES" : "` + strconv.Itoa( int( tmp3.BytesSent )) + `"}}}}`
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

func NetInfoAdapterSummaryBytesRecv( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "RECEIVED" : { "BYTES" : "` + strconv.Itoa( int( tmp3.BytesRecv )) + `"}}}}`
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

func NetInfoAdapterSummaryPacketsSent( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "SENT" : { "PACKETS" : "` + strconv.Itoa( int( tmp3.PacketsSent )) + `"}}}}`
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

func NetInfoAdapterSummaryPacketsRecv( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "RECEIVED" : { "PACKETS" : "` + strconv.Itoa( int( tmp3.BytesRecv )) + `"}}}}`
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

func NetInfoAdapterSummaryErrorIn( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "ERROR" : { "IN" : "` + strconv.Itoa( int( tmp3.Errin )) + `"}}}}`
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

func NetInfoAdapterSummaryErrorOut( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "ERROR" : { "OUT" : "` + strconv.Itoa( int( tmp3.Errout )) + `"}}}}`
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

func NetInfoAdapterSummaryDropsIn( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "ERROR" : { "IN" : "` + strconv.Itoa( int( tmp3.Dropin )) + `"}}}}`
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

func NetInfoAdapterSummaryDropsOut( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	tmp2, _ := net.NetIOCounters( false )
	tmp3 := tmp2[ 0 ]

	result := `{ "HOST" : { "NET" : { "ERROR" : { "OUT" : "` + strconv.Itoa( int( tmp3.Dropout )) + `"}}}}`
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

func NetInfoAdaptersSummary( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	adapter := c.URLParams[ "interface" ]
	tmp2, _ := net.NetIOCounters( true )

	result := `{ "HOST" : { "NET" : { "ADAPTERS" : {`
	for i := 0; i < len( tmp2 ); i++ {
		do := true
		if adapter != "" && adapter != tmp2[ i ].Name {
			do = false
		}
		if do {
			result += `"ADAPTER" : {`
			result += `"NAME" : "` + tmp2[ i ].Name + `",`
			result += `"BYTESSENT" : "` + strconv.Itoa( int( tmp2[ i ].BytesSent )) + `",`
			result += `"BYTESRECV" : "` + strconv.Itoa( int( tmp2[ i ].BytesRecv )) + `",`
			result += `"PACKETSSENT" : "` + strconv.Itoa( int( tmp2[ i ].PacketsSent )) + `",`
			result += `"PACKETSRECV" : "` + strconv.Itoa( int( tmp2[ i ].PacketsRecv )) + `",`
			result += `"ERRORSIN" : "` + strconv.Itoa( int( tmp2[ i ].Errin )) + `",`
			result += `"ERRORSOUT" : "` + strconv.Itoa( int( tmp2[ i ].Errout )) + `",`
			result += `"DROPSIN" : "` + strconv.Itoa( int( tmp2[ i ].Dropin )) + `",`
			result += `"DROPSOUT" : "` + strconv.Itoa( int( tmp2[ i ].Dropout )) + `"`
			result += `},`
		}
	}
	result = result[:len( result ) - 1 ]
	result += `}}}}`

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

func NetInfoAdaptersSummaryDetail( c web.C, w http.ResponseWriter, req *http.Request ) {
	tmp := req.Header[ "Accept" ]
	adapter := c.URLParams[ "interface" ]
	tmp2, _ := net.NetInterfaces()

	result := `{ "HOST" : { "NET" : { "ADAPTERS" : {`
	for i := 0; i < len( tmp2 ); i++ {
		do := true
		if adapter != "" && adapter != tmp2[ i ].Name {
			do = false
		}
		if do {
			result += `"ADAPTER" : {`
			result += `"NAME" : "` + tmp2[ i ].Name + `",`
			result += `"MTU" : "` + strconv.Itoa( int( tmp2[ i ].MTU )) + `",`
			result += `"HARDWAREADDR" : "` + tmp2[ i ].HardwareAddr + `",`
			result += `"FLAGS" : {`
			for j := 0; j < len( tmp2[ i ].Flags ); j++ {
				result += `"FLAG" : "` + tmp2[ i ].Flags[ j ] + `"`
				if j + 1 != len( tmp2[ i ].Flags ) { result += `,` }
			}
			result += `},`
			result += `"ADDRESSES" : {`
			for k := 0; k < len( tmp2[ i ].Addrs ); k++ {
				result += `"ADDRESS" : "` + tmp2[ i ].Addrs[ k ].Addr + `"`
				if k + 1 != len( tmp2[ i ].Addrs ) { result += `,` }
			}
			result += `}},`
		}
	}
	result = result[:len( result ) - 1 ]
	result += `}}}}`

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
