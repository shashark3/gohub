package sms

//Driver 将短信提供商抽象化，如果后面有需要，可以灵活更换短信商
type Driver interface {
	//发送短信
	Send(phone string, message Message, config map[string]string) bool
}
