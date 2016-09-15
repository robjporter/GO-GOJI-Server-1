package system

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	color "../modules/colour"
)

func (application *Application) DisplayIntro() {
	configRuntime(application)
}

func configRuntime(application *Application) {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	proc := color.Red(nuCPU)
	ver := runtime.Version()
	var spacer = 50
	var mode = ""
	var port = ""
	var ip = getLocalIP()

	fmt.Println(getLine(spacer, "-"))
	t := time.Now().Format("02/01/2006 - 15:04:05")
	fmt.Println(getLine(getSpacer(spacer, t), " "), color.Yellow(t))
	fmt.Println(getLine(spacer, "-"))
	fmt.Printf("CPU Count:      %s\n", proc)
	fmt.Printf("Go Version:     %s\n", color.Red(ver[2:]))
	fmt.Printf("IP Address:     %s\n", color.Red(ip))
	//fmt.Printf( "Operating mode: %s\n", color.Red( coreConfig.GetString( "environment.mode" )))
	if application.Configuration.Host.Https {
		mode = "https"
	} else {
		mode = "http"
	}
	port = application.Configuration.Host.Port
	fmt.Printf("Style Template: %s\n", color.Red(application.Configuration.Style))
	fmt.Printf("Listening mode: %s\n", color.Red(mode))
	fmt.Printf("Listening port: %s\n", color.Red(port))
	fmt.Printf("Listening URL:  %s://%s:%s/\n", color.Red(mode), color.Red(ip), color.Red(port))
	fmt.Println(getLine(spacer, "-"))
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (application *Application) GetPlugins() []Module {
	var modules []Module
	if Exists(application.Configuration.PluginPath) {
		files, _ := ioutil.ReadDir(application.Configuration.PluginPath)
		for _, f := range files {
			path := application.Configuration.PluginPath + "/" + f.Name()
			if IsDirectory(path) {
				a := Module{Name: f.Name(), Path: path}
				modules = append(modules, a)
			}
		}
		return modules
	} else {
		panic("Application path does not exist.....Please correct and try again.")
		os.Exit(1)
	}
	return nil
}

func (application *Application) ReadStyleTemplates() {
	tmpFolder := "views" + string(filepath.Separator) + application.Configuration.Style

	fileList := []string{}

	err := filepath.Walk(tmpFolder, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			if strings.HasSuffix(f.Name(), ".html") {
				fileList = append(fileList, path)
			}
		}
		return nil
	})

	if err != nil {
		panic("There was an error parsing the template directory.  Please check and rerun the application.")
		os.Exit(1)
	}

	application.Configuration.Templates = fileList
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func getLine(num int, char string) string {
	str := ""
	for i := 0; i < num; i++ {
		str += char
	}
	return str
}

func getSpacer(num int, word string) int {
	var space = 0
	space = (num - len(word)) / 2
	return space
}

func (application *Application) ClearTerminal() {
	ops := runtime.GOOS
	switch ops {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GenerateUUID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}

func IsDirectory(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

func SecondsToDate(seconds2 float64) (int, int, int, int, int) {
	seconds := int(seconds2)
	weeks := seconds / (7 * 24 * 60 * 60)
	days := seconds/(24*60*60) - 7*weeks
	hours := seconds/(60*60) - 7*24*weeks - 24*days
	minutes := seconds/60 - 7*24*60*weeks - 24*60*days - 60*hours
	seconds3 := seconds - 7*24*60*60*weeks - 24*60*60*days - 60*60*hours - 60*minutes

	return weeks, days, hours, minutes, seconds3
}

func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func ReturnSizeByName(name string) string {
	switch name {
	case "svga":
		return ReturnSizeByXY("800", "600")
	case "xga":
		return ReturnSizeByXY("1024", "768")
	case "wxga":
		return ReturnSizeByXY("1280", "720")
	case "wxga2":
		return ReturnSizeByXY("1280", "800")
	case "sxga":
		return ReturnSizeByXY("1280", "1024")
	case "hd":
		return ReturnSizeByXY("1366", "768")
	case "uxga":
		return ReturnSizeByXY("1600", "1200")
	case "fhd":
		return ReturnSizeByXY("1920", "1080")
	case "qhd":
		return ReturnSizeByXY("2560", "1440")
	case "wqxga":
		return ReturnSizeByXY("2560", "1600")
	case "uhd":
		return ReturnSizeByXY("3840", "2160")

	}
	return "UNKNOWN"
}

func ReturnSizeByXY(x string, y string) string {
	return "_" + x + "x" + y + ".jpg"
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func GetURL(url string, method string, token string, headers map[string]string, parameters map[string]string, data []byte) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if token == "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
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
