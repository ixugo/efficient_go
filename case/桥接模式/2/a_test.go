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

// 消息接口 - 抽象层
type Message interface {
	Send(notifier Notifier) error
}

// 通知器接口 - 实现层
type Notifier interface {
	SendMessage(content string) error
}

// 具体消息类型
type TextMessage struct {
	content string
}

func (m *TextMessage) Send(notifier Notifier) error {
	return notifier.SendMessage(m.content)
}

type ImageMessage struct {
	imageURL string
}

func (m *ImageMessage) Send(notifier Notifier) error {
	return notifier.SendMessage(m.imageURL)
}

type VideoMessage struct {
	videoURL string
}

func (m *VideoMessage) Send(notifier Notifier) error {
	return notifier.SendMessage(m.videoURL)
}

// 具体通知器实现
type EmailNotifier struct{}

func (e *EmailNotifier) SendMessage(content string) error {
	fmt.Println("发送邮件:", content)
	return nil
}

type SmsNotifier struct{}

func (s *SmsNotifier) SendMessage(content string) error {
	fmt.Println("发送短信:", content)
	return nil
}

type WechatNotifier struct{}

func (w *WechatNotifier) SendMessage(content string) error {
	fmt.Println("发送微信:", content)
	return nil
}

// 消息发送服务
type MessageService struct {
	notifier Notifier
}

func NewMessageService(notifier Notifier) *MessageService {
	return &MessageService{notifier: notifier}
}

func (s *MessageService) SendMessage(message Message) error {
	return message.Send(s.notifier)
}

func TestCase1(t *testing.T) {
	// 使用邮件发送器
	emailService := NewMessageService(&EmailNotifier{})
	emailService.SendMessage(&TextMessage{content: "text test"})
	emailService.SendMessage(&ImageMessage{imageURL: "img test"})
	emailService.SendMessage(&VideoMessage{videoURL: "video test"})

	// 使用短信发送器
	smsService := NewMessageService(&SmsNotifier{})
	smsService.SendMessage(&TextMessage{content: "text test"})

	// 使用微信发送器
	wechatService := NewMessageService(&WechatNotifier{})
	wechatService.SendMessage(&TextMessage{content: "text test"})
}
