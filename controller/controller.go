package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/atulmirajkar/atgo/model"
)

var controllerObj *Controller

//Config represent configuration helper class
type Config struct {
	LogFile      string `json:"logFile"`
	BaseViewPath string `json:"baseViewPath"`
}

//Controller represents the controller class in MVC pattern
type Controller struct {
	basePath     string
	configObject Config
	modelObject  *model.Model
	logger       *log.Logger
}

func init() {
	controllerObj = &Controller{}
	//initialize logger
	var file *os.File
	var e error
	controllerObj.basePath = os.Getenv("GOPATH") + "/src/github.com/atulmirajkar/atgo/"
	if e = controllerObj.readConfig(controllerObj.basePath + "config.json"); e != nil {
		return
	}

	fmt.Println(controllerObj.configObject.LogFile)
	if file, e = os.Create(controllerObj.configObject.LogFile); e != nil {
		fmt.Println("File creation error:")
		fmt.Print(e)
		return
	}
	controllerObj.logger = log.New(file, "log:", log.LstdFlags|log.Lshortfile)

	//initialize model
	controllerObj.modelObject = &model.Model{}
	controllerObj.modelObject.InitializeModel(controllerObj.logger)
}

func (controllerObj *Controller) readConfig(configFilePath string) error {
	var file []byte
	var e error
	file, e = ioutil.ReadFile(configFilePath)
	if e != nil {
		fmt.Println("File error:")
		fmt.Print(e)
		return e
	}
	var configObject = Config{}
	if e := json.Unmarshal(file, &configObject); e != nil {
		fmt.Println("Unmarshalling error:")
		fmt.Print(e)
		return e
	}
	configObject.LogFile = controllerObj.basePath + configObject.LogFile
	configObject.BaseViewPath = controllerObj.basePath + configObject.BaseViewPath
	controllerObj.configObject = configObject
	return nil
}

//StartServer - function called by controller to start server
func StartServer() {
	fmt.Printf("Started Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", mux)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path[1:])

}

//handler for static files like images and css
func viewHandler(w http.ResponseWriter, r *http.Request) {
	controllerObj.logger.Println(controllerObj.basePath + r.URL.Path[1:])
	http.ServeFile(w, r, controllerObj.basePath+r.URL.Path[1:])
}

//handler for files with templates
func homeHandler(w http.ResponseWriter, r *http.Request) {
	var view *(model.View)
	var err error
	var t = &template.Template{}

	if view, err = controllerObj.modelObject.LoadPage(controllerObj.configObject.BaseViewPath + "/data/" + r.URL.Path[1:] + ".json"); err != nil {
		controllerObj.logger.Println(err)
		return
	}
	//initialie the base path from where to find static pages
	view.BaseViewPath = controllerObj.configObject.BaseViewPath
	controllerObj.logger.Println(view)

	//define template functions
	t = t.Funcs(template.FuncMap{"ConvertToId": view.Page.DefineSidebarId})
	if t, err = t.ParseFiles(view.BaseViewPath+"/html/"+r.URL.Path[1:], controllerObj.basePath+"/template/content.tmpl", controllerObj.basePath+"/template/header.tmpl", controllerObj.basePath+"/template/sidebar.tmpl"); err != nil {
		controllerObj.logger.Println("template parsing err:")
		controllerObj.logger.Print(err)
		return
	}

	//t.ExecuteTemplate(os.Stdout, "header", view)
	//t.ExecuteTemplate(os.Stdout, "sidebar", view.Page.Sidebar)
	//t.ExecuteTemplate(os.Stdout, "content", view)

	if err = t.ExecuteTemplate(w, r.URL.Path[1:], view); err != nil {
		controllerObj.logger.Println("template executing  err:")
		controllerObj.logger.Print(err)
		return
	}

}
