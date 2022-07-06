package main

import "C"
import (
	"fmt"
	"github.com/donpark/pam"
	"log/syslog"
)

type mypam struct {
	// your pam vars
}

func (mp *mypam) Authenticate(hdl pam.Handle, args pam.Args) pam.Value {
	log("Authenticate: %v", args)
	user, err := hdl.GetUser()
	log("user: %v , err:%v", user, err)

	_, _ = hdl.Conversation(pam.Message{
		Style: pam.MessageTextInfo,
		Msg:   "-\n-\nAAAAAAAAAAAA\n-\n",
	})

	data, err := hdl.Conversation(pam.Message{
		Style: pam.MessageEchoOff,
		Msg:   "Password: ",
	})
	log("data: %v, err: %v", data, err)

	data, err = hdl.Conversation(pam.Message{
		Style: pam.MessageEchoOff,
		Msg:   "Security Code: ",
	})
	log("data: %v, err: %v", data, err)

	return pam.Success
}

func (mp *mypam) SetCredential(hdl pam.Handle, args pam.Args) pam.Value {
	log("SetCredential: %v", args)
	return pam.Success
}

var mp mypam

func init() {
	pam.RegisterAuthHandler(&mp)
}

func main() {
	// needed in c-shared buildmode
}

func log(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-demo")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}
