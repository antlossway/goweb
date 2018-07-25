package main

//use http.Get , http.Post, http.PostForm, defaultCLient

import (
	"github.com/antlossway/goweb/httpclient"
)

func main() {
	//GetRequest()
	//PostForm()
	httpclient.MyGetRequest("https://rest.nexmo.com/sms/json")
}
