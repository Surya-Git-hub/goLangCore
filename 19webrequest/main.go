package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type : %T\n", res)
	fmt.Println("respons:", res)
	defer res.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
