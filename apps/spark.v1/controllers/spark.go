package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"../../../render"
	"github.com/stretchr/objx"
	"github.com/zenazn/goji/web"
)

var rooms []byte

func SparkHome(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates(c)
	templates = append(templates, "apps/spark.v1/views/home.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SparkDisplay(c web.C, w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}

func SparkMembershipByRoomID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	status := 200
	token := req.Header.Get("SparkToken")
	id := strings.TrimSpace(c.URLParams["id"])
	result := `{"SparkRoom":{"Status":"Failed","Data":"` + id + `"}}`

	parameters["roomId"] = id

	response, status, err := GetURL("https://api.ciscospark.com/v1/memberships", "GET", token, headers, parameters, []byte(`{}`))

	fmt.Println("ERROR:", err)
	fmt.Println("STATUS:", status)
	fmt.Println("RESPONSE:", string(response[:]))

	render.RenderJSON(w, status, result)
}

func SparkMembershipByPersonID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	status := 200
	token := req.Header.Get("SparkToken")
	id := strings.TrimSpace(c.URLParams["id"])
	result := `{"SparkRoom":{"Status":"Failed","Data":"` + id + `"}}`

	parameters["personId"] = id

	response, status, err := GetURL("https://api.ciscospark.com/v1/memberships", "GET", token, headers, parameters, []byte(`{}`))

	fmt.Println("ERROR:", err)
	fmt.Println("STATUS:", status)
	fmt.Println("RESPONSE:", string(response[:]))

	render.RenderJSON(w, status, result)
}

func SparkMembershipByEmail(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	status := 200
	token := req.Header.Get("SparkToken")
	id := strings.TrimSpace(c.URLParams["id"])
	result := `{"SparkRoom":{"Status":"Failed","Data":"` + id + `"}}`

	parameters["personEmail"] = id

	response, status, err := GetURL("https://api.ciscospark.com/v1/memberships", "GET", token, headers, parameters, []byte(`{}`))

	fmt.Println("ERROR:", err)
	fmt.Println("STATUS:", status)
	fmt.Println("RESPONSE:", string(response[:]))

	render.RenderJSON(w, status, result)
}

func SparkRoomsDeleteRoom(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	response := rooms
	status := 200
	id := ""
	token := req.Header.Get("SparkToken")
	data := strings.TrimSpace(c.URLParams["name"])
	result := `{"SparkRoom":{"Status":"Failed","Data":"` + data + `"}}`

	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		if err == nil {
			rooms = response
		}
	}

	found, name := roomsExistsByID(response, data)
	if found {
		id = data
	} else {
		found, id = roomsExistsByName(response, data)
		name = data
	}

	if found {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms/"+id, "DELETE", token, headers, parameters, []byte(`{}`))
		if err == nil {
			result = `{"SparkRoom":{"Status":"Success","Name":"` + name + `","id":"` + id + `"}}`
		}
	} else {
		result = `{"SparkRoom":{"Status":"Failed","Name":"` + name + `","Error":"Room not found.","ID":"` + id + `"}}`
	}

	render.RenderJSON(w, status, result)
}

func SparkRoomsChangeName(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	response := rooms
	status := 200
	token := req.Header.Get("SparkToken")
	id := strings.TrimSpace(c.URLParams["name"])
	name2 := strings.TrimSpace(c.URLParams["name2"])
	result := `{"SparkRoom":{"Status":"Failed","ID":"` + id + `",Name":"` + name2 + `"}}`

	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		if err == nil {
			rooms = response
		}
	}

	found, name := roomsExistsByID(response, id)
	if !found {
		found, id = roomsExistsByName(response, id)
	}

	if found {
		body := `{"title":"` + name2 + `"}`
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms/"+id, "PUT", token, headers, parameters, []byte(body))

		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			id := document.Get("id").Str()
			result = `{"SparkRoom":{"Status":"Success","Name":"` + name + `","id":"` + id + `"}}`
		}
	} else {
		result = `{"SparkRoom":{"Status":"Failed","Name":"` + name + `","Error":"Room not found.","ID":"` + id + `"}}`
	}

	render.RenderJSON(w, status, result)
}

func SparkRoomsChangeNameByID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	response := rooms
	status := 200
	token := req.Header.Get("SparkToken")
	id := strings.TrimSpace(c.URLParams["id"])
	name2 := strings.TrimSpace(c.URLParams["name2"])
	result := `{"SparkRoom":{"Status":"Failed","ID":"` + id + `",Name":"` + name2 + `"}}`

	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		if err == nil {
			rooms = response
		}
	}
	found, name := roomsExistsByID(response, id)

	if found {
		body := `{"title":"` + name2 + `"}`
		response, status, err := GetURL("https://api.ciscospark.com/v1/rooms/"+id, "PUT", token, headers, parameters, []byte(body))

		fmt.Println("RESPONSE:", response)
		fmt.Println("STATUS:", status)
		fmt.Println("ERROR:", err)

		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			id := document.Get("id").Str()
			result = `{"SparkRoom":{"Status":"Success","Name":"` + name + `","id":"` + id + `"}}`
		}
	} else {
		result = `{"SparkRoom":{"Status":"Failed","Name":"` + name + `","Error":"Room not found.","ID":"` + id + `"}}`
	}
	render.RenderJSON(w, status, result)
}

//name = original name
//name2 = new name
//id = room id

func SparkRoomsCreateRoomByName(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)

	response := rooms
	status := 200
	token := req.Header.Get("SparkToken")
	name := strings.TrimSpace(c.URLParams["name"])
	result := `{"SparkRoom":{"Status":"Failed","Name":"` + name + `"}}`

	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		if err == nil {
			rooms = response
		}
	}
	found, id := roomsExistsByName(response, name)

	if found {
		result = `{"SparkRoom":{"Status":"Failed","Name":"` + name + `","Error":"Room already exists.","ID":"` + id + `"}}`
	} else {
		body := `{"title":"` + name + `"}`
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "POST", token, headers, parameters, []byte(body))
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			id := document.Get("id").Str()
			result = `{"SparkRoom":{"Status":"Success","Name":"` + name + `","id":"` + id + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func roomsExistsByID(response []byte, name string) (bool, string) {
	id := ""
	document, err2 := objx.FromJSON(string(response[:]))
	if err2 == nil {
		cont := true
		count := 0
		for cont {
			if document.Get("items["+strconv.Itoa(count)+"].id").Str() == "" {
				cont = false
			} else {
				if document.Get("items["+strconv.Itoa(count)+"].id").Str() == name {
					id = document.Get("items[" + strconv.Itoa(count) + "].title").Str()
					return true, id
				}
			}
			count += 1
			//Simple check to ensure we are not stuck in an endless loop....User is unlikely to be a member of 1000+ groups
			if count > 1000 {
				cont = false
			}
		}
	}
	return false, id
}

func roomsExistsByName(response []byte, name string) (bool, string) {
	id := ""
	document, err2 := objx.FromJSON(string(response[:]))
	if err2 == nil {
		cont := true
		count := 0
		for cont {
			if document.Get("items["+strconv.Itoa(count)+"].id").Str() == "" {
				cont = false
			} else {
				if document.Get("items["+strconv.Itoa(count)+"].title").Str() == name {
					id = document.Get("items[" + strconv.Itoa(count) + "].id").Str()
					return true, id
				}
			}
			count += 1
			//Simple check to ensure we are not stuck in an endless loop....User is unlikely to be a member of 1000+ groups
			if count > 1000 {
				cont = false
			}
		}
	}
	return false, id
}

func SparkPeopleFindMe(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	token := req.Header.Get("SparkToken")
	result := `{"SparkUser":{"Status":"Not Found","Name":"me"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people/me", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			fmt.Println(document)
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","ID":"` + document.Get("items[0].id").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindNameReturnCreated(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	result := `{"SparkUser":{"Status":"Not Found","Name":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Created":"` + document.Get("items[0].created").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmailReturnCreated(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	result := `{"SparkUser":{"Status":"Not Found","Email":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Created":"` + document.Get("items[0].created").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindNameReturnAvatar(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	result := `{"SparkUser":{"Status":"Not Found","Name":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Avatar":"` + document.Get("items[0].avatar").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmailReturnAvatar(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	result := `{"SparkUser":{"Status":"Not Found","Email":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Avatar":"` + document.Get("items[0].avatar").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindNameReturnName(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	result := `{"SparkUser":{"Status":"Not Found","Name":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmailReturnName(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	result := `{"SparkUser":{"Status":"Not Found","Email":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindNameReturnEmails(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	result := `{"SparkUser":{"Status":"Not Found","Name":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Emails":"` + document.Get("items[0].emails[0]").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmailReturnEmails(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	result := `{"SparkUser":{"Status":"Not Found","Email":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","Emails":"` + document.Get("items[0].emails[0]").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindNameReturnID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	result := `{"SparkUser":{"Status":"Not Found","Name":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Name":"` + document.Get("items[0].displayName").Str() + `","ID":"` + document.Get("items[0].id").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmailReturnID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	result := `{"SparkUser":{"Status":"Not Found","Email":"` + name + `"}}`
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			result = `{"SparkUser":{"Status":"Found","Email":"` + document.Get("items[0].displayName").Str() + `","ID":"` + document.Get("items[0].id").Str() + `"}}`
		}
	}
	render.RenderJSON(w, status, result)
}

func SparkPeopleFindEmail(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["email"]
	token := req.Header.Get("SparkToken")
	parameters["email"] = name
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		render.RenderJSON(w, status, string(response[:]))
	}
}

func SparkPeopleFindName(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	parameters["displayName"] = name
	response, status, err := GetURL("https://api.ciscospark.com/v1/people", "GET", token, headers, parameters, []byte(`{}`))

	if err == nil {
		render.RenderJSON(w, status, string(response[:]))
	}
}

func SparkRoomsCheckRoomExists(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	response := rooms
	status := 200
	name := c.URLParams["name"]
	token := req.Header.Get("SparkToken")
	id := ""
	result := `{"Status":"NotFound","Name":"` + name + `"}`
	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		rooms = response
	}
	if err == nil {
		if status == 200 {
			document, err2 := objx.FromJSON(string(response[:]))
			if err2 == nil {
				cont := true
				count := 0
				for cont {
					if document.Get("items["+strconv.Itoa(count)+"].id").Str() == "" {
						cont = false
					} else {
						if document.Get("items["+strconv.Itoa(count)+"].title").Str() == name {
							id = document.Get("items[" + strconv.Itoa(count) + "].id").Str()
							result = `{"Status":"Found","Name":"` + name + `","ID":"` + id + `"}`
						}
					}
					count += 1
					//Simple check to ensure we are not stuck in an endless loop....User is unlikely to be a member of 1000+ groups
					if count > 1000 {
						cont = false
					}
				}
			}
		}
	}
	render.RenderJSON(w, http.StatusOK, result)
}

func SparkRoomsGetRoomNameFromID(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	response := rooms
	status := 200
	name := c.URLParams["id"]
	token := req.Header.Get("SparkToken")
	id := ""
	result := `{"Status":"NotFound","ID":"` + name + `"}`
	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		rooms = response
	}
	if err == nil {
		if status == 200 {
			document, err2 := objx.FromJSON(string(response[:]))
			if err2 == nil {
				cont := true
				count := 0
				for cont {
					if document.Get("items["+strconv.Itoa(count)+"].id").Str() == "" {
						cont = false
					} else {
						if document.Get("items["+strconv.Itoa(count)+"].id").Str() == name {
							id = document.Get("items[" + strconv.Itoa(count) + "].title").Str()
							result = `{"Status":"Found","ID":"` + name + `","Name":"` + id + `"}`
						}
					}
					count += 1
					//Simple check to ensure we are not stuck in an endless loop....User is unlikely to be a member of 1000+ groups
					if count > 1000 {
						cont = false
					}
				}
			}
		}
	}
	render.RenderJSON(w, http.StatusOK, result)
}

func SparkRoomsListAll(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	response := rooms
	status := 200
	token := req.Header.Get("SparkToken")
	if rooms == nil {
		response, status, err = GetURL("https://api.ciscospark.com/v1/rooms", "GET", token, headers, parameters, []byte(`{}`))
		rooms = response
	}
	if err == nil {
		if status == 200 {
			render.RenderJSON(w, status, string(response[:]))
		}
	}
}

func GetURL(url string, method string, token string, headers map[string]string, parameters map[string]string, data []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	values := req.URL.Query()
	for key, value := range parameters {
		values.Add(key, value)
	}
	req.URL.RawQuery = values.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 500, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, 200, nil
}
