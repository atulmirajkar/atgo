package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Model struct {
	logger *log.Logger
}

type Header struct {
	Title string `json:"title"`
}

type BodyParts struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
}

type SidebarLink struct {
	Label string `json: "label"`
	Link  string `json: "link"`
}

type View struct {
	Page         *Page
	BaseViewPath string
}

type Page struct {
	Head    Header        `json:"header"`
	Body    []BodyParts   `json:"body"`
	Sidebar []SidebarLink `json:"-"`
}

func (m *Model) InitializeModel(inputLogger *log.Logger) {
	m.logger = inputLogger
}

func (m *Model) LoadPage(dataPath string) (*View, error) {
	//dataPath = "./view/data/" + dataPath + ".json"
	file, err := ioutil.ReadFile(dataPath)
	if err != nil {
		m.logger.Println("File error: %v\n", err)
		return nil, err
	}

	//unmarshall the json file
	var page Page
	if e := json.Unmarshal(file, &page); e != nil {
		m.logger.Println("Unmarshal error: %s for the file %s\n", err, dataPath)
		return nil, e
	}

	//make sidebar slice
	page.Sidebar = make([]SidebarLink, 0, 5)
	//read content and create the sidebar content
	for i := 0; i < len(page.Body); i++ {
		page.buildSideBarIds(page.Body[i].Heading)
	}
	return &View{Page: &page}, nil
}

func (p *Page) getAbsLink(inputLink string) {

}
func (page *Page) buildSideBarIds(heading string) {
	//add to sidebar links
	id := strings.Join(strings.Split(heading, " "), "")
	page.Sidebar = append(page.Sidebar, SidebarLink{Label: heading, Link: page.Head.Title + "#" + id})

}

//concatenate the input string and return
func (page *Page) DefineSidebarId(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	id := strings.Join(strings.Split(s, " "), "")
	return id
}
