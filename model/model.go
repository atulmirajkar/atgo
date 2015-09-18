package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Header struct {
	Title string `json:"title"`
}

type BodyParts struct {
	Heading string `json:"heading"`
	Content string `jsom:"content"`
}

type Page struct {
	Head Header      `json:"header"`
	Body []BodyParts `json:"body"`
}

func LoadPage(dataPath string) (*Page, error) {
	dataPath = "./view/data/" + dataPath + ".json"
	file, err := ioutil.ReadFile(dataPath)
	if err != nil {
		fmt.Println("File error: %v\n", err)
		return nil, err
	}

	fmt.Println(string(file))
	//unmarshall the json file
	var page Page
	if e := json.Unmarshal(file, &page); e != nil {
		fmt.Println("Unmarshal error: %s\n", err)
		return nil, e
	}
	return &page, nil
}
