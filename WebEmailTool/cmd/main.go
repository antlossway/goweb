package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
)

func main() {
	//initialize general settings
	app := GetInit()

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Promotion)
	mux.HandleFunc("/Promotion", app.Promotion)
	mux.HandleFunc("/Notification", app.Notification)
	mux.HandleFunc("/Submit", app.Submit)

	fileServer := http.FileServer(http.Dir(app.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server listen on ...", app.ListenAddr)
	err := http.ListenAndServe(app.ListenAddr, mux)
	log.Fatal(err)
}

// getListFromJson read from json file which contain list of Contact, and return []Contact
func (app *App) getListFromJSON(cfg string) []Contact {
	mycfg := filepath.Join(app.CfgDir, cfg)
	bs, err := ioutil.ReadFile(mycfg)
	if err != nil {
		log.Fatalln("Can't open", mycfg)
	}

	var contactList []Contact
	err = json.Unmarshal(bs, &contactList)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return contactList
}

// Promotion send promotion email to contact list
func (app *App) Promotion(w http.ResponseWriter, r *http.Request) {
	//mylist is []Contact
	mylist := app.getListFromJSON("promotion_contact_json.cfg")

	li := List{
		Tab:         app.Tab,
		Title:       "Send Promotion",
		ContactList: mylist,
	}
	app.RenderForm(w, r, "base", li)
}

// Notification send maintenance schedule, or service outage/recovery to contact list
func (app *App) Notification(w http.ResponseWriter, r *http.Request) {
	//mylist is []Contact
	mylist := app.getListFromJSON("notification_contact_json.cfg")

	//log.Println("debug notification list: ", mylist)

	li := List{
		Tab:         app.Tab,
		Title:       "Send Notification",
		ContactList: mylist,
	}
	app.RenderForm(w, r, "base", li)
}

func (app *App) Submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("%s\n%s", err.Error(), debug.Stack())
		http.Error(w, "ParseForm error", http.StatusInternalServerError)
		return
	}

	subject := r.PostForm.Get("subject")
	content := r.PostForm.Get("content")

	msg := Message{
		Subject:   subject,
		Content:   content,
		Lines:     strings.Split(content, "\n"),
		Signature: app.Signature,
	}

	params := r.PostForm //result of r.PostForm is type url.Values => map[string] []string

	re := regexp.MustCompile("company.*$")

	for key, value := range params { //key: company{{.Company}}, value: {{ .Company}}:{{.Emails}}

		if re.Match([]byte(key)) { //catch the field which contain the company name and email

			strValue := strings.Join(value, ",") // convert from []string to string
			items := strings.Split(strValue, ":")

			msg.Attention = items[0]
			msg.To = items[1]

			if err := app.SendHtmlEmail(&msg); err != nil {
				log.Printf("%s\n%s", err.Error(), debug.Stack())
				http.Error(w, "sendHtmlEmail error", http.StatusInternalServerError)
				return
			}

			app.RenderForm(w, r, "confirmation", msg)

		}
	}

}
