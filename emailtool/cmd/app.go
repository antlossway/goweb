package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
)

// App Define an struct to hold the application-wide dependencies and configuration settings for our web application.
type App struct {
	ListenAddr   string //http server listen address IP:Port
	HTMLDir      string //hold html template
	StaticDir    string //hold static files
	CfgDir       string //hold smtp.cfg, contact.cfg
	From         string //Email From
	Cc           string //Email Cc
	Signature    string //Email Signature
	SMTPhost     string
	SMTPaddr     string
	SMTPusername string
	SMTPpassword string
	Tab          []string //website top level navigation tab
}

// GetInit initialize globle configurations
func GetInit() *App {
	addr := flag.String("addr", ":4000", "HTTP Server Listen Address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	cfgDir := flag.String("cfg-dir", "./cfg", "Path to configuration files")
	staticDir := flag.String("static-dir", "./ui/static", "Path to Static assets")
	flag.Parse()

	//get general email field: From and Cc
	emailcfg := filepath.Join(*cfgDir, "email.cfg")
	bs, err := ioutil.ReadFile(emailcfg)
	if err != nil {
		log.Fatalln("Can't open", emailcfg)
	}

	var myEmailVar EmailVar
	err = json.Unmarshal(bs, &myEmailVar)

	if err != nil {
		log.Fatalln("json.Unmarshal myEmailVar: ", err.Error())
	}

	//get SMTP setting
	smtpcfg := filepath.Join(*cfgDir, "smtp.cfg")
	bs, err = ioutil.ReadFile(smtpcfg)
	if err != nil {
		log.Fatalln("Can't open", smtpcfg)
	}
	//debug
	//log.Println("result from reading smtpcfg", string(bs))

	var mysmtp SMTP
	err = json.Unmarshal(bs, &mysmtp)

	if err != nil {
		log.Fatalln("json.Unmarshal mysmtp:", err.Error())
	}

	//debug
	//log.Println("mysmtp:", mysmtp)

	app := &App{
		ListenAddr:   *addr,
		HTMLDir:      *htmlDir,
		CfgDir:       *cfgDir,
		StaticDir:    *staticDir,
		From:         myEmailVar.From,
		Cc:           myEmailVar.Cc,
		Signature:    myEmailVar.Signature,
		SMTPaddr:     mysmtp.Addr,
		SMTPhost:     mysmtp.Host,
		SMTPusername: mysmtp.Username,
		SMTPpassword: mysmtp.Password,
		Tab:          []string{"Promotion", "Notification"},
	}

	return app
}
