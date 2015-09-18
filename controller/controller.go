package controller

import (
	//"errors"
	"fmt"
	"github.com/atulmirajkar/atgo/model"
	"html/template"
	"net/http"
)

type Controller struct {
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	//mux.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8080", mux)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path[1:])

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var page = &model.Page{}
	var err error
	var t = &template.Template{}
	if page, err = model.LoadPage(r.URL.Path[1:]); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(page)
	if t, err = template.ParseFiles("./template/testTemplate"); err != nil {
		fmt.Println("template parsing err %s\n", err)
		return
	}
	if err = t.Execute(w, page); err != nil {
		fmt.Println("template executing  err %s\n", err)
		return
	}
	//fmt.Fprintf(w, "loading Page %s,%s", page.Head.Title, t)

}
