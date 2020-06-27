package global

import "github.com/astaxie/beego/utils"

//发送邮件帮助类
func sendEmail(toEmail []string, content string) error {
	//config := `{"username":"1451670604@qq.com","password":"urqtuhejvxqdhfdb","host":"smtp.qq.com","port":587}`
	config := `{"username":"renjucode@163.com", "password":"FKRJBQMLSMEKBISM", "host":"smtp.163.com", "port":25}`
	email := utils.NewEMail(config)
	email.To = toEmail
	email.From = "renjucode@163.com"
	email.Subject = "【五子连珠】验证码"
	email.Text = content
	return email.Send()
}
