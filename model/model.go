package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	Sidebar []SidebarLink `json:"sidebar"`
}

func (m *Model) InitializeModel(inputLogger *log.Logger) {
	m.logger = inputLogger
}

func (m *Model) LoadPage(dataPath string) (*View, error) {
	dataPath = "./view/data/" + dataPath + ".json"
	file, err := ioutil.ReadFile(dataPath)
	if err != nil {
		m.logger.Println("File error: %v\n", err)
		return nil, err
	}

	//unmarshall the json file
	var page Page
	if e := json.Unmarshal(file, &page); e != nil {
		m.logger.Println("Unmarshal error: %s\n", err)
		return nil, e
	}

	return &View{Page: &page}, nil
}

func (p *Page) getAbsLink(inputLink string) {

}
