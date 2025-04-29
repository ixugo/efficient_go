package sitonggo

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// 桥是抽象的，两岸是具体的实现，通过抽象让两岸任意组合。

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

// 如何写出桥接模式??
// 1. 定义两个接口，左岸和右岸，左岸稳定，右岸变化。
// 2. 定义 service 桥，连接两岸，创建对象时传入左岸，执行函数时传入右岸
// 3. 各自实现前面定义的两个接口，即可

// 导出接口 - 抽象层
type Exporter interface {
	Export(data interface{}) error
}

// 数据源接口 - 实现层
type DataSource interface {
	GetData() (interface{}, error)
}

// 具体导出格式实现
type ExcelExporter struct {
	filePath string
}

func (e *ExcelExporter) Export(data interface{}) error {
	fmt.Printf("导出Excel到文件: %s\n", e.filePath)
	// 实际项目中这里会使用excel库来写入数据
	return nil
}

type CSVExporter struct {
	filePath string
}

func (c *CSVExporter) Export(data interface{}) error {
	fmt.Printf("导出CSV到文件: %s\n", c.filePath)
	// 实际项目中这里会使用csv库来写入数据
	return nil
}

type JSONExporter struct {
	filePath string
}

func (j *JSONExporter) Export(data interface{}) error {
	fmt.Printf("导出JSON到文件: %s\n", j.filePath)
	// 实际项目中这里会使用json库来写入数据
	return json.NewEncoder(os.Stdout).Encode(data)
}

// 具体数据源实现
type MySQLDataSource struct {
	tableName string
}

func (m *MySQLDataSource) GetData() (interface{}, error) {
	// 模拟从MySQL获取数据
	return map[string]interface{}{
		"source": "MySQL",
		"table":  m.tableName,
		"data":   []string{"row1", "row2", "row3"},
	}, nil
}

type MongoDBDataSource struct {
	collection string
}

func (m *MongoDBDataSource) GetData() (interface{}, error) {
	// 模拟从MongoDB获取数据
	return map[string]interface{}{
		"source":     "MongoDB",
		"collection": m.collection,
		"data":       []string{"doc1", "doc2", "doc3"},
	}, nil
}

type RedisDataSource struct {
	key string
}

func (r *RedisDataSource) GetData() (interface{}, error) {
	// 模拟从Redis获取数据
	return map[string]interface{}{
		"source": "Redis",
		"key":    r.key,
		"data":   "value",
	}, nil
}

// 导出服务 - 桥接器
type ExportService struct {
	exporter Exporter
}

func NewExportService(exporter Exporter) *ExportService {
	return &ExportService{exporter: exporter}
}

func (s *ExportService) ExportData(source DataSource) error {
	data, err := source.GetData()
	if err != nil {
		return fmt.Errorf("获取数据失败: %v", err)
	}

	if err := s.exporter.Export(data); err != nil {
		return fmt.Errorf("导出数据失败: %v", err)
	}

	return nil
}

func TestCase1(t *testing.T) {
	// 测试Excel导出
	excelService := NewExportService(&ExcelExporter{filePath: "data.xlsx"})
	if err := excelService.ExportData(&MySQLDataSource{tableName: "users"}); err != nil {
		t.Errorf("Excel导出失败: %v", err)
	}

	// 测试CSV导出
	csvService := NewExportService(&CSVExporter{filePath: "data.csv"})
	if err := csvService.ExportData(&MongoDBDataSource{collection: "orders"}); err != nil {
		t.Errorf("CSV导出失败: %v", err)
	}

	// 测试JSON导出
	jsonService := NewExportService(&JSONExporter{filePath: "data.json"})
	if err := jsonService.ExportData(&RedisDataSource{key: "cache:user:1"}); err != nil {
		t.Errorf("JSON导出失败: %v", err)
	}
}
