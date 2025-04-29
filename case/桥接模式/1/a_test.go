package sitonggo

import (
	"fmt"
	"testing"
)

/*
桥接模式考题说明：

题目：实现一个跨平台的消息通知系统

背景：
你正在开发一个消息通知系统，需要支持多种消息类型（如文本消息、图片消息、视频消息等）和多种发送渠道（如邮件、短信、微信等）。
系统需要能够灵活地组合不同的消息类型和发送渠道。

要求：
1. 实现一个基础的消息通知系统，支持以下功能：
   - 发送文本消息
   - 发送图片消息
   - 发送视频消息
2. 支持通过以下渠道发送消息：
   - 邮件
   - 短信
   - 微信
3. 系统需要能够方便地扩展新的消息类型和发送渠道

提示：
- 使用桥接模式来解耦消息类型和发送渠道
- 考虑如何设计抽象层和实现层
- 注意消息类型和发送渠道的独立性

请先按照自己的想法实现，然后我会指导你如何重构为桥接模式。
*/

func TestCase1(t *testing.T) {
	SendMessage(1, "text test")
	SendImageMessage(1, "img test ")
	SendVideoMessage(1, "video test")
}

type Notifier interface {
	SendTextMessage(message string) error
	SendImageMessage(imageURL string) error
	SendVideoMessage(videoURL string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) SendTextMessage(message string) error {
	fmt.Println("发送邮件", message)
	return nil
}

func (e *EmailNotifier) SendImageMessage(imageURL string) error {
	fmt.Println("发送邮件", imageURL)
	return nil
}

func (e *EmailNotifier) SendVideoMessage(videoURL string) error {
	fmt.Println("发送邮件", videoURL)
	return nil
}

type SmsNotifier struct{}

func (s *SmsNotifier) SendTextMessage(message string) error {
	fmt.Println("发送短信", message)
	return nil
}

func (s *SmsNotifier) SendImageMessage(imageURL string) error {
	fmt.Println("发送短信", imageURL)
	return nil
}

func (s *SmsNotifier) SendVideoMessage(videoURL string) error {
	fmt.Println("发送短信", videoURL)
	return nil
}

type WechatNotifier struct{}

func (w *WechatNotifier) SendTextMessage(message string) error {
	fmt.Println("发送微信", message)
	return nil
}

func (w *WechatNotifier) SendImageMessage(imageURL string) error {
	fmt.Println("发送微信", imageURL)
	return nil
}

func (w *WechatNotifier) SendVideoMessage(videoURL string) error {
	fmt.Println("发送微信", videoURL)
	return nil
}

func SendMessage(atype int, data string) {
	switch atype {
	case 1:
		a := EmailNotifier{}
		a.SendTextMessage(data)
	case 2:
		a := SmsNotifier{}
		a.SendTextMessage(data)
	case 3:
		a := WechatNotifier{}
		a.SendTextMessage(data)
	}
}

func SendImageMessage(atype int, data string) {
	switch atype {
	case 1:
		a := EmailNotifier{}
		a.SendImageMessage(data)
	case 2:
		a := SmsNotifier{}
		a.SendImageMessage(data)
	case 3:
		a := WechatNotifier{}
		a.SendImageMessage(data)
	}
}

func SendVideoMessage(atype int, data string) {
	switch atype {
	case 1:
		a := EmailNotifier{}
		a.SendVideoMessage(data)
	case 2:
		a := SmsNotifier{}
		a.SendVideoMessage(data)
	case 3:
		a := WechatNotifier{}
		a.SendVideoMessage(data)
	}
}
