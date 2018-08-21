//here keep important stuct type
package main

//record of Company and emails, input source can be file or database
type Contact struct {
	Company string `json:"company"`
	Emails  string `json:"emails"` //emails separated by comma,
}

//List holds the values filled in form
type List struct {
	Tab         []string
	Title       string
	ContactList []Contact
	Subject     string
	Content     string
	Result      string
}

// email related info
type Message struct {
	To        string //To: email address list separated by comma
	From      string
	Cc        string
	Subject   string
	Content   string
	Attention string   //Dear xxxxx team
	Lines     []string //content separate to []string, to be processed by template emailbody and add <br/>
	Signature string
}

type EmailVar struct {
	From      string `json:"from"`
	Cc        string `json:"cc"`
	Signature string `json:"signature"`
}

type SMTP struct {
	Host     string `json:"host"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}
