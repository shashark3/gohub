package sms

import (
	"gohub/pkg/config"
	"sync"
)

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

// SMS 发送短信的操作类
type SMS struct {
	Driver Driver
}

var once sync.Once

var internalSMS *SMS

func NewSms() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &aliyun{},
		}
	})
	return internalSMS
}

func (s *SMS) Send(phone string, message Message) bool {
	return s.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
