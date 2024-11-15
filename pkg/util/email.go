package util

import (
	"bytes"
	"crypto/tls"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/util/logger"
	"github.com/nbb2025/distri-domain/pkg/util/validate"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strings"
)

// EmailTemplate 全局变量，用于存储邮件模板的内容
var EmailTemplate string

func init() {
	b, e := os.ReadFile("config/email_template.html")
	if e != nil {
		logger.Error(e.Error())
	}
	EmailTemplate = string(b)
}

// SendEmail cc为抄送人,bcc为暗抄,org为组织名(自定义)
func SendEmail(to []string, org, subject, text string, cc []string, bcc []string) {
	go func() {
		con := config.Conf.MailConfig

		m := gomail.NewMessage()
		m.SetHeader("From", org+"<"+con.Account+">") // 增加发件人别名
		m.SetHeader("To", to...)
		// 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
		if con.BCC != "" && validate.Email(con.BCC) {
			bcc = append(bcc, con.BCC)
		}
		if len(bcc) > 0 {
			m.SetHeader("Bcc", bcc...) // 暗送，可以多个
		}
		if con.CC != "" && validate.Email(con.CC) {
			cc = append(cc, con.CC)
		}
		if len(cc) > 0 {
			m.SetHeader("Cc", cc...) // 抄送，可以多个
		}

		if org != "" && !strings.Contains(subject, "["+org+"]") {
			subject = "[" + org + "]" + subject
		}

		m.SetHeader("Subject", subject) // 邮件主题

		m.SetBody("text/html", BuildEmailContent(text))
		d := gomail.NewDialer(
			con.SmtpServer,
			con.SmtpPort,
			con.Account,
			con.Password,
		)
		// 关闭SSL协议认证
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		if err := d.DialAndSend(m); err != nil {
			logger.Error(err.Error())
		}
	}()
}

func BuildEmailContent(text string) string {
	tmp := EmailTemplate
	// 加载邮件模板
	tmpl, err := template.New("email").Parse(tmp)
	if err != nil {
		logger.Error("Error parsing email template:" + err.Error())
		return text
	}

	// 准备模板数据
	data := struct {
		Body      template.HTML
		URLPrefix template.URL
	}{
		Body:      template.HTML(text), // 替换邮件正文内容
		URLPrefix: template.URL(config.Conf.MailConfig.URLPrefix),
	}

	// 执行模板替换
	var resultTemplate bytes.Buffer
	err = tmpl.Execute(&resultTemplate, data)
	if err != nil {
		logger.Error("Error executing email template:" + err.Error())
		return text
	}
	return resultTemplate.String()
}
