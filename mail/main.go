package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	msg := gomail.NewMessage()
	//1. 设置发件人信息
	//msg.SetHeader("From", "zczhai@baai.ac.cn")
	msg.SetHeader("From", msg.FormatAddress("zczhai@baai.ac.cn", "算力平台"))
	// 2. 设置收件人信息 值为  ...string 可设置多个
	msg.SetHeader("To", "zczhai@baai.ac.cn")
	// 3. 如果需要可以填充 cc，也就是抄送
	//msg.SetHeader("Cc", "cc_address@example.com")
	// 4. 设置邮件标题
	msg.SetHeader("Subject", "你好")
	// 5. 设置要发送的邮件正文
	// 第一个参数是类型，第二个参数是内容
	// 如果是 html，第一个参数则是 `text/html`
	msg.SetBody("text/html", "<h1>test</h1>")
	// 6. 设置附件信息，直接传入附件的绝对路径
	//msg.Attach("C:/Users/ifigh/Desktop/Temps/timg.jpg")

	// 传入 服务地址、端口、账号、密码 等参数，初始化 dialer实例
	dialer := gomail.NewDialer("smtp.mxhichina.com", 465, "zczhai@baai.ac.cn", "73ca58ed")

	// 将 msg信息 实例作为参数 传入 dialer实例，进行发送
	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println(err)
	}

	fmt.Println("发送成功！！！")

}
