package sitonggo

import (
	"io"
	"os"
	"testing"
)

/*
桥接模式考题说明：

题目：实现一个数据导出系统

背景：
你正在开发一个数据管理平台，需要支持多种数据格式（如Excel、CSV、JSON等）的导出功能，
同时需要支持不同的数据源（如MySQL、MongoDB、Redis等）。

要求：
1. 实现一个数据导出系统，支持以下功能：
   - 导出Excel格式
   - 导出CSV格式
   - 导出JSON格式
2. 支持从以下数据源导出：
   - MySQL
   - MongoDB
   - Redis
3. 系统需要能够方便地扩展新的数据格式和数据源

提示：
- 使用桥接模式来解耦数据格式和数据源
- 考虑如何设计抽象层和实现层
- 注意数据格式和数据源的独立性

请先按照自己的想法实现，然后我会指导你如何重构为桥接模式。
*/

/*
桥接模式适用场景说明：

1. 适用场景：
   - 当一个类存在两个独立变化的维度，且这两个维度都需要进行扩展时
   - 当一个系统需要在构件的抽象化角色和具体化角色之间增加更多的灵活性时
   - 当不希望使用继承或因为多层继承导致系统类的个数急剧增加时

2. 在CRUD管理平台中的常见应用场景：
   - 数据导出系统：数据格式（Excel/CSV/JSON）和数据源（MySQL/MongoDB）的分离
   - 权限管理系统：权限类型（菜单/按钮/数据）和验证方式（RBAC/ABAC）的分离
   - 日志系统：日志类型（操作/错误/性能）和存储方式（文件/数据库/ES）的分离
   - 缓存系统：缓存策略（LRU/LFU）和存储介质（内存/Redis）的分离
   - 消息通知：消息类型（文本/图片）和发送渠道（邮件/短信）的分离
   - 数据验证：验证规则（必填/格式）和验证方式（前端/后端）的分离

3. 使用桥接模式的好处：
   - 解耦抽象和实现，使它们可以独立变化
   - 提高系统的可扩展性
   - 减少代码重复
   - 提高代码的可维护性
*/

// 导出接口
type Exporter interface {
	export(DataSource) error
}

// 数据源接口
type DataSource interface {
	read() ([]byte, error)
}

type Excel struct {
	w io.Writer
}

func (e *Excel) export(ds DataSource) error {
	data, err := ds.read()
	if err != nil {
		return err
	}
	e.w.Write([]byte("导出Excel\t"))
	e.w.Write(data)
	return nil
}

type CSV struct {
	w io.Writer
}

func (c *CSV) export(ds DataSource) error {
	data, err := ds.read()
	if err != nil {
		return err
	}
	c.w.Write([]byte("导出CSV\t"))
	c.w.Write(data)
	return nil
}

type JSON struct {
	w io.Writer
}

func (j *JSON) export(ds DataSource) error {
	data, err := ds.read()
	if err != nil {
		return err
	}
	j.w.Write([]byte("导出JSON\t"))
	j.w.Write(data)
	return nil
}

type MySQL struct{}

func (m *MySQL) read() ([]byte, error) {
	return []byte("MySQL数据\n"), nil
}

type MongoDB struct{}

func (m *MongoDB) read() ([]byte, error) {
	return []byte("MongoDB数据\n"), nil
}

type Redis struct{}

func (r *Redis) read() ([]byte, error) {
	return []byte("Redis数据\n"), nil
}

type ExportService struct {
	exporter Exporter
}

func (e *ExportService) export(ds DataSource) error {
	return e.exporter.export(ds)
}

func TestCase1(t *testing.T) {
	es1 := ExportService{
		exporter: &Excel{
			w: os.Stdout,
		},
	}
	es1.export(&MySQL{})

	es2 := ExportService{
		exporter: &CSV{
			w: os.Stdout,
		},
	}
	es2.export(&MongoDB{})

	es3 := ExportService{
		exporter: &JSON{
			w: os.Stdout,
		},
	}
	es3.export(&Redis{})
}
