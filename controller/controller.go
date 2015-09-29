package controller

import (
	"encoding/json"
	"fmt"
	"github.com/atulmirajkar/atgo/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var controllerObj *Controller

type Config struct {
	LogFile      string `json:"logFile"`
	BaseViewPath string `json:"baseViewPath"`
}

type Controller struct {
	configObject Config
	modelObject  *model.Model
	logger       *log.Logger
}

func init() {
	controllerObj = &Controller{}
	//initialize logger
	var file *os.File
	var e error
	if e = controllerObj.readConfig("config.json"); e != nil {
		return
	}

	if file, e = os.Create(controllerObj.configObject.LogFile); e != nil {
		fmt.Println("File creation error: %v", e)
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
		fmt.Println("File error: %v", e)
		return e
	}
	var configObject = Config{}
	if e := json.Unmarshal(file, &configObject); e != nil {
		fmt.Println("Unmarshalling error %v", e)
		return e
	}
	controllerObj.configObject = configObject
	return nil
}

func StartServer() {
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
	http.ServeFile(w, r, r.URL.Path[1:])
}

//handler for files with templates
func homeHandler(w http.ResponseWriter, r *http.Request) {
	var view *(model.View)
	var err error
	var t = &template.Template{}

	if view, err = controllerObj.modelObject.LoadPage(r.URL.Path[1:]); err != nil {
		controllerObj.logger.Println(err)
		return
	}
	//initialie the base path from where to find static pages
	view.BaseViewPath = controllerObj.configObject.BaseViewPath

	controllerObj.logger.Println(view)
	if t, err = template.ParseFiles("./template/"+r.URL.Path[1:], "./template/content.tmpl", "./template/header.tmpl", "./template/sidebar.tmpl"); err != nil {
		controllerObj.logger.Println("template parsing err %s\n", err)
		return
	}

	//t.ExecuteTemplate(os.Stdout, "header", view)
	//t.ExecuteTemplate(os.Stdout, "sidebar", view.Page.Sidebar)
	//t.ExecuteTemplate(os.Stdout, "content", view)
	if err = t.Execute(w, view); err != nil {
		controllerObj.logger.Println("template executing  err %s\n", err)
		return
	}

}
