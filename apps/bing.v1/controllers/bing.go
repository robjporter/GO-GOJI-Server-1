package controllers

import (
	"errors"
	"net/http"
	"strings"

	"../../../render"
	"../../../system"
	"github.com/stretchr/objx"
	"github.com/zenazn/goji/web"
)

func BingHome(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates(c)
	test := c.Env["Settings"].(*system.Settings)
	test.Count += 1
	templates = append(templates, "apps/bing.v1/views/home.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BingAbout(c web.C, w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates(c)
	templates = append(templates, "apps/bing.v1/views/about.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "About"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BingDailyPhoto(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "

	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			url := "http://www.bing.com" + document.Get("images[0].url").Str()
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoEmbed(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphotoembed.html")
	body := ``
	status := 200

	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			url := "http://www.bing.com" + document.Get("images[0].url").Str()
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoRaw(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	body := ``
	status := 200

	result := `{"Bing":{"Status":"Failed","Data":"Failed to get information from Bing"}}`
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		result = `{"Bing":{"Status":"Success","Data":` + string(response[:]) + `}}`
	}

	render.RenderJSON(w, status, result)
}

func BingDailyPhotoSized(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	x := c.URLParams["x"]
	y := c.URLParams["y"]

	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByXY( x, y)
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoSVGA(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("svga")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoXGA(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("xga")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}


func BingDailyPhotoWXGA(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("wxga")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoWXGA2(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("wxga2")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoHD(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("hd")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoFHD(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("fhd")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoQHD(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("qhd")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoWQXGA(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("wqxga")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}

func BingDailyPhotoUHD(c web.C, w http.ResponseWriter, req *http.Request) {
	var err interface{} = nil
	headers := make(map[string]string)
	parameters := make(map[string]string)
	templates := []string{}
	values := map[string]string{}
	templates = append(templates, "apps/bing.v1/views/dailyphoto.html")
	body := ``
	status := 200
	values["Title"] = "Bing Daily - "
	
	response, status, err := system.GetURL("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1", "GET", "", headers, parameters, []byte(body))

	if err == nil {
		document, err2 := objx.FromJSON(string(response[:]))
		if err2 == nil {
			tmp := document.Get("images[0].url").Str()
			i := strings.LastIndex(tmp, "_")
			url := tmp[:i]
			url = "http://www.bing.com" + url + system.ReturnSizeByName("uhd")
			values["URL"] = url
			err := render.RenderTemplate(w, templates, "base", values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	render.RenderJSON(w, status, errors.New(""))
}