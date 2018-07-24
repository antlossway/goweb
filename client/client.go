package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func main(){
	GetRequest()
}

func GetRequest(){
	resp, err := http.Get("https://rest.nexmo.com/sms/json")
	if err != nil { // error would be reported only if there was an issue connecting to the server 
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	log.Println(string(body))

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
}