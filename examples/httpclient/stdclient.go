package httpclient

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

//standard http client, using DefaultClient

//GetRequest: send HTTP GET
func GetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil { // error would be reported only if there was an issue connecting to the server
		log.Fatalln(err)
	}

	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	*/
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
}

//send HTTP PSOT
func PostForm() {
	formData := url.Values{
		"api_key":    {"toto"},
		"api_secret": {"secret"},
	}

	resp, err := http.PostForm("https://rest.nexmo.com/sms/json", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["message-count"])

	//println(result["form"])

}
