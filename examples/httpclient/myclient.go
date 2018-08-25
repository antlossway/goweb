package httpclient

//implementing own http client, which can define timeout

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const mytimeout = 5 * time.Second

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

//simple get request, return normal body
func MyGet(url string) {
	myclient := http.Client{
		//		Timeout: time.Duration(5 * time.Second),
		Timeout: mytimeout,
	}
	//resp, err := myclient.Get("https://httpbin.org/get")
	resp, err := myclient.Get(url)

	defer resp.Body.Close() //always close Response Body to prevent resource leak

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	log.Println(string(body))

}

//result in Json
func MyGetJsonResp(url string) {
	myclient := http.Client{
		//		Timeout: time.Duration(5 * time.Second),
		Timeout: mytimeout,
	}
	//resp, err := myclient.Get("https://httpbin.org/get")
	resp, err := myclient.Get(url)
	check(err)

	defer resp.Body.Close() //always close Response Body to prevent resource leak

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)

}
