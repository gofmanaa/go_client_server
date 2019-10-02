package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/?ls=-alh")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if len(responseData) > 0 {
		body := string(responseData)
		fmt.Printf("%v", body)
	}

}
