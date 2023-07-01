package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Plaform  string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "LearncodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "LearncodeOnline.in", "bcd123", []string{"fullstack-dev", "js"}},
		{"Angular Bootcamp", 299, "LearncodeOnline.in", "sur123", nil},
	}

	//package this in json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Reactjs Bootcamp",
		"Price": 299,
		"website": "LearncodeOnline.in",
		"tags": ["web-dev","js"]
	}
`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//different use case example

	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
