package engine

type IOCom struct {
	Outmsg chan string
	Errmsg chan string
	Done   chan bool
}