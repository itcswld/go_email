package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	/*--compnay email server
	  from := "a_acc@domain.com.tw"
	  to := "b_acc@gmail.com"
	  user := "eve"
	  pwd := "your_pwd"
	  host := `smtp.domainName.com.tw`
	  port := 25
	*/
	var (
		from = "a_acc@gmail.com"
		to   = from
		user = from
		pwd  = "your_pwd"
		host = `smtp.gmail.com`
		port = 587
	)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Eve</b>!")
	d := gomail.NewDialer(host, port, user, pwd)
	//x509: certificate signed by unknown authority的解决方法
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("mail send successfully!")
}
