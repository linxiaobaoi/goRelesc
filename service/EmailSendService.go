/**
* Created by GoLand
* User: lingm
* Date: 2023/6/12
* Time: 下午 01:57
* Author: 现在的努力是为了小时候吹过的NB
* Atom: 小白从不写注释！！！
 */

package service

import (
	"fmt"
	"github.com/linxiaobaoi/goRelesc.git/config"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"time"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList []string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

// MailboxConf 邮箱配置
type MailConf struct {
	// 邮件标题
	Title string
	// 邮件内容
	Body string
	// 收件人列表
	RecipientList string
	// 发件人账号
	Sender string
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string
	// SMTP端口 QQ邮箱是25
	SMTPPort int
}

// 发送邮箱验证码
func Send(recipientList string) (string, bool) {
	var mailConf MailConf
	//title: "goRelease"
	//fromAcoont: "1281102864@qq.com" #发送者邮箱号
	//password: "vuekfiekshnqbacg"
	//smtAddr: "smtp.qq.com"
	//smtpPort: 587  #端口号
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	title := viper.GetString("email.title")
	send := viper.GetString("email.fromAcoont")
	password := viper.GetString("email.password")
	smtAddr := viper.GetString("email.smtAddr")
	smtpPort := viper.GetInt("email.smtpPort")
	//title := "知了"
	mailConf.Title = title
	//这里就是我们发送的邮箱内容，但是也可以通过下面的html代码作为邮件内容
	// mailConf.Body = "坚持才是胜利，奥里给"
	//这里支持群发，只需填写多个人的邮箱即可，我这里发送人使用的是QQ邮箱，所以接收人也必须都要是

	//QQ邮箱
	mailConf.RecipientList = recipientList
	mailConf.Sender = send

	//这里QQ邮箱要填写授权码，网易邮箱则直接填写自己的邮箱密码，授权码获得方法在下面
	mailConf.SPassword = password
	//下面是官方邮箱提供的SMTP服务地址和端口
	// QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	// 雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	// 163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	// 126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	// 新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）
	mailConf.SMTPAddr = smtAddr
	mailConf.SMTPPort = smtpPort

	//产生六位数验证码
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	//vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	vcode := config.RandNumber()
	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的<strong>xxx</strong>，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)

	m := gomail.NewMessage()

	// 第三个参数是我们发送者的名称，但是如果对方有发送者的好友，优先显示对方好友备注名
	m.SetHeader(`From`, mailConf.Sender, "知了")
	m.SetHeader(`To`, mailConf.RecipientList)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	// m.Attach("./Dockerfile") //添加附件
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		return "发送失败", false
	} else {
		exp := config.ExpLates(4)
		config.SetRedis("email_"+recipientList, vcode, time.Duration(exp))
		return vcode, true
	}
}
