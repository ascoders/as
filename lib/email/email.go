package email

import (
	"github.com/ascoders/as"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strconv"
)

// 发送
func Send(address []string, title string, html string) error {
	e := email.NewEmail()
	e.From = as.Conf.EmailFrom
	e.To = address
	e.Subject = title
	e.Text = []byte("邮件无法显示")
	e.HTML = []byte(`
		<div style="border-bottom:3px solid #d9d9d9; background:url(http://www.wokugame.com/static/img/email_bg.gif) repeat-x 0 1px;">
			<div style="border:1px solid #c8cfda; padding:40px;">
				` + html + `
				<p>&nbsp;</p>
				<div>我酷游戏团队 祝您游戏愉快</div>
				<div>Powered by wokugame</div>
				<img src="http://www.wokugame.com/static/img/logo.png">
				</div>
			</div>
		</div>
	`)
	return e.Send(as.Conf.EmailHost+":"+strconv.Itoa(as.Conf.EmailPort),
		smtp.PlainAuth("", as.Conf.EmailFrom, as.Conf.EmailPassword, strconv.Itoa(as.Conf.EmailPort)))
}
