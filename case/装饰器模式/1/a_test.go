package sitonggo

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

/*
装饰器模式考题说明：

题目：实现一个日志记录系统

背景：
你正在开发一个日志记录系统，需要支持多种日志记录方式，并且这些方式可以灵活组合。
比如：既需要记录到文件，又需要发送到监控系统，还需要进行数据脱敏等。

要求：
1. 实现基础的日志记录功能：
   - 记录到控制台
   - 记录到文件
   - 记录到远程服务器

2. 实现日志装饰功能：
   - 时间戳装饰：每条日志添加时间戳
   - 脱敏装饰：对敏感信息进行脱敏处理
   - 压缩装饰：对日志内容进行压缩
   - 加密装饰：对日志内容进行加密

3. 支持装饰器的灵活组合：
   - 可以任意组合多个装饰器
   - 装饰器的顺序可以调整
   - 可以方便地添加新的装饰器

提示：
- 使用装饰器模式来扩展日志记录功能
- 考虑如何设计基础日志记录器和装饰器
- 注意装饰器的组合方式

请先按照自己的想法实现，然后我会指导你如何优化。
*/

/*
装饰器模式适用场景说明：

1. 适用场景：
   - 需要动态地给一个对象添加额外的职责
   - 需要灵活地组合多个功能
   - 不想通过继承来扩展功能

2. 在日志系统中的常见应用：
   - 日志格式化：添加时间戳、日志级别等
   - 日志处理：脱敏、压缩、加密等
   - 日志输出：控制台、文件、远程等
   - 日志过滤：按级别、关键字等

3. 使用装饰器模式的好处：
   - 可以动态添加功能
   - 避免继承带来的类爆炸
   - 功能可以灵活组合
   - 符合开闭原则
*/

// Logger 是日志记录器的基本接口
// 所有日志记录器和装饰器都必须实现这个接口
type Logger interface {
	Write(message string) error
}

// LoggerDecorator 是装饰器接口
// 它继承了Logger接口，并添加了设置被装饰对象的方法
type LoggerDecorator interface {
	Logger
	SetLogger(logger Logger)
}

// BaseDecorator 是所有装饰器的基类
// 它持有一个被装饰的Logger对象
type BaseDecorator struct {
	logger Logger // 被装饰的Logger对象
}

func (d *BaseDecorator) SetLogger(logger Logger) {
	d.logger = logger
}

// 基础日志记录器实现
type ConsoleLogger struct{}

func (c *ConsoleLogger) Write(message string) error {
	fmt.Println("控制台日志:", message)
	return nil
}

type FileLogger struct {
	filePath string
}

func (f *FileLogger) Write(message string) error {
	fmt.Printf("文件日志[%s]: %s\n", f.filePath, message)
	return nil
}

type RemoteLogger struct {
	server string
}

func (r *RemoteLogger) Write(message string) error {
	fmt.Printf("远程日志[%s]: %s\n", r.server, message)
	return nil
}

// 具体装饰器实现
type TimestampDecorator struct {
	BaseDecorator
}

func (d *TimestampDecorator) Write(message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return d.logger.Write(fmt.Sprintf("[%s] %s", timestamp, message))
}

type DesensitizeDecorator struct {
	BaseDecorator
}

func (d *DesensitizeDecorator) Write(message string) error {
	// 简单的脱敏处理，实际项目中会更复杂
	desensitized := strings.ReplaceAll(message, "password", "****")
	desensitized = strings.ReplaceAll(desensitized, "token", "****")
	return d.logger.Write(desensitized)
}

type CompressDecorator struct {
	BaseDecorator
}

func (d *CompressDecorator) Write(message string) error {
	// 模拟压缩处理
	compressed := fmt.Sprintf("[COMPRESSED]%s", message)
	return d.logger.Write(compressed)
}

type EncryptDecorator struct {
	BaseDecorator
}

func (d *EncryptDecorator) Write(message string) error {
	// 模拟加密处理
	encrypted := fmt.Sprintf("[ENCRYPTED]%s", message)
	return d.logger.Write(encrypted)
}

// NewLoggerChain 创建装饰器链
// 参数说明：
//   - logger: 基础的日志记录器
//   - decorators: 要应用的装饰器列表，按顺序排列
//
// 返回值：最外层的装饰器，它包含了整个装饰器链
//
// 工作原理：
// 1. 装饰器链的构建是从内到外的
// 2. 最后一个装饰器装饰基础logger
// 3. 倒数第二个装饰器装饰最后一个装饰器
// 4. 以此类推，直到第一个装饰器
//
// 例如：NewLoggerChain(logger, A, B, C) 会创建如下链：
// A -> B -> C -> logger
// 当调用A.Write()时，调用顺序是：
// A.Write() -> B.Write() -> C.Write() -> logger.Write()
func NewLoggerChain(logger Logger, decorators ...LoggerDecorator) Logger {
	if len(decorators) == 0 {
		return logger
	}

	// 从后向前设置装饰器
	// 例如：decorators = [A, B, C]
	// 第一次循环：i=2, C装饰logger
	// 第二次循环：i=1, B装饰C
	// 第三次循环：i=0, A装饰B
	for i := len(decorators) - 1; i >= 0; i-- {
		if i == len(decorators)-1 {
			// 最后一个装饰器装饰基础logger
			decorators[i].SetLogger(logger)
		} else {
			// 其他装饰器装饰前一个装饰器
			decorators[i].SetLogger(decorators[i+1])
		}
	}

	// 返回最外层的装饰器
	return decorators[0]
}

func TestCase1(t *testing.T) {
	// 创建基础日志记录器
	consoleLogger := &ConsoleLogger{}
	fileLogger := &FileLogger{filePath: "app.log"}
	remoteLogger := &RemoteLogger{server: "log-server:8080"}

	// 测试基础日志
	consoleLogger.Write("这是一条普通日志")
	fileLogger.Write("这是一条普通日志")
	remoteLogger.Write("这是一条普通日志")

	// 测试装饰器组合
	// 1. 控制台日志 + 时间戳 + 脱敏
	// 创建装饰器链：DesensitizeDecorator -> TimestampDecorator -> ConsoleLogger
	consoleChain := NewLoggerChain(
		consoleLogger,
		&TimestampDecorator{},
		&DesensitizeDecorator{},
	)
	// 调用链：DesensitizeDecorator.Write() -> TimestampDecorator.Write() -> ConsoleLogger.Write()
	consoleChain.Write("用户登录，password=123456, token=abc123")

	// 2. 文件日志 + 时间戳 + 压缩
	// 创建装饰器链：CompressDecorator -> TimestampDecorator -> FileLogger
	fileChain := NewLoggerChain(
		fileLogger,
		&TimestampDecorator{},
		&CompressDecorator{},
	)
	fileChain.Write("这是一条需要压缩的日志")

	// 3. 远程日志 + 时间戳 + 脱敏 + 加密
	// 创建装饰器链：EncryptDecorator -> DesensitizeDecorator -> TimestampDecorator -> RemoteLogger
	remoteChain := NewLoggerChain(
		remoteLogger,
		&TimestampDecorator{},
		&DesensitizeDecorator{},
		&EncryptDecorator{},
	)
	remoteChain.Write("敏感信息：password=123456, token=abc123")
}
