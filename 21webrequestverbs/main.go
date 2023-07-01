package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web request verbs")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println("content : ", string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json payload
	requestBody := strings.NewReader(`
	{
		"course":"Lets go with golang",
		"price":0,
		"platform":"learnCodeOnline.in"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
