package httpclient

//implementing own http client, which can define timeout

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func MyGetRequest(url string) {
	myclient := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	//resp, err := myclient.Get("https://httpbin.org/get")
	resp, err := myclient.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close() //always close Response Body to prevent resource leak

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}
