package email

/**
{ "to" : "imneo47@gmail.com",	"sub" : "test sub", "msg" : "test msg" }
*/

type Email struct {
	To      string `json:"to"`
	Subject string `json:"sub"`
	Msg     string `json:"msg"`
}
