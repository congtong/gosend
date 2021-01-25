package main
import (
	"gopkg.in/gomail.v2"
)
const (
	// 邮件服务器地址
	MAIL_HOST	= "smtp.qq.com"
	// 端口
	MAIL_PORT	= 465
	// 发送邮件用户账号
	MAIL_USER	= ""
	// 授权密码
	MAIL_PWD	= ""
)

func SendGoMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "gomail"

	m.SetHeader("From",nickname + "<" + MAIL_USER + ">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MAIL_HOST, MAIL_PORT, MAIL_USER, MAIL_PWD)
	err := d.DialAndSend(m)
	return err
}

func Alert() {
	address := []string{"276620500@qq.com"}
	SendGoMail(address, "ETH算力为0预警", "ETH的电脑是不是断电了还是断网了")
}
