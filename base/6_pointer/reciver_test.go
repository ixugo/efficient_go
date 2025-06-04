
package pointer
// 模拟 MediaAPI 结构体
type MockAPI struct {
	name     string
	vodCover *string // 模拟延迟初始化的字段
}

// 值接收者方法 - 会复现问题
func (a MockAPI) GetNameValue() string {
	if a.vodCover == nil {
		return a.name + ": vodCover is nil"
	}
	return a.name + ": vodCover is " + *a.vodCover
}

// 指针接收者方法 - 正确的做法
func (a *MockAPI) GetNamePtr() string {
	if a.vodCover == nil {
		return a.name + ": vodCover is nil"
	}
	return a.name + ": vodCover is " + *a.vodCover
}

// 模拟路由注册函数
func registerRoutes(api *MockAPI) (func() string, func() string) {
	// 这里创建方法值，模拟 gin 路由注册时的情况
	valueMethod := api.GetNameValue // 值接收者方法值
	ptrMethod := api.GetNamePtr     // 指针接收者方法值

	return valueMethod, ptrMethod
}

// 测试方法值捕获时机问题
func TestMethodValueCapture(t *testing.T) {
	// 1. 创建 API 实例（此时 vodCover 为 nil）
	api := &MockAPI{
		name:     "TestAPI",
		vodCover: nil, // 初始为 nil
	}

	// 2. 注册路由（模拟 registerMedia 函数）
	// 此时创建方法值，捕获当前状态
	valueMethod, ptrMethod := registerRoutes(api)

	// 3. 延迟初始化（模拟 setupRouter 中的赋值）
	coverValue := "initialized_cover"
	api.vodCover = &coverValue

	// 4. 测试直接调用 - 都应该能看到新值
	t.Log("=== 直接调用测试 ===")
	directValue := api.GetNameValue()
	directPtr := api.GetNamePtr()
	t.Logf("直接调用值接收者: %s", directValue)
	t.Logf("直接调用指针接收者: %s", directPtr)

	// 5. 测试方法值调用 - 展示问题
	t.Log("=== 方法值调用测试 ===")
	methodValue := valueMethod() // 值接收者方法值
	methodPtr := ptrMethod()     // 指针接收者方法值
	t.Logf("方法值(值接收者): %s", methodValue)
	t.Logf("方法值(指针接收者): %s", methodPtr)

	// 6. 验证结果
	if directValue != "TestAPI: vodCover is initialized_cover" {
		t.Errorf("直接调用值接收者失败: got %s", directValue)
	}
	if directPtr != "TestAPI: vodCover is initialized_cover" {
		t.Errorf("直接调用指针接收者失败: got %s", directPtr)
	}

	// 关键测试：方法值的行为差异
	if methodValue != "TestAPI: vodCover is nil" {
		t.Errorf("值接收者方法值应该显示 nil，但得到: %s", methodValue)
	}
	if methodPtr != "TestAPI: vodCover is initialized_cover" {
		t.Errorf("指针接收者方法值应该显示新值，但得到: %s", methodPtr)
	}

	t.Log("=== 结论 ===")
	t.Log("值接收者的方法值在创建时就捕获了结构体的副本")
	t.Log("指针接收者的方法值捕获的是指针，能看到后续的修改")
}

// 测试多次修改的情况
func TestMultipleModifications(t *testing.T) {
	api := &MockAPI{name: "MultiTest", vodCover: nil}

	// 创建方法值
	valueMethod, ptrMethod := registerRoutes(api)

	// 第一次修改
	cover1 := "first_value"
	api.vodCover = &cover1

	t.Logf("第一次修改后 - 方法值(值接收者): %s", valueMethod())
	t.Logf("第一次修改后 - 方法值(指针接收者): %s", ptrMethod())

	// 第二次修改
	cover2 := "second_value"
	api.vodCover = &cover2

	t.Logf("第二次修改后 - 方法值(值接收者): %s", valueMethod())
	t.Logf("第二次修改后 - 方法值(指针接收者): %s", ptrMethod())

	// 验证：值接收者方法值始终显示 nil，指针接收者方法值显示最新值
	if valueMethod() != "MultiTest: vodCover is nil" {
		t.Error("值接收者方法值应该始终显示初始状态")
	}
	if ptrMethod() != "MultiTest: vodCover is second_value" {
		t.Error("指针接收者方法值应该显示最新状态")
	}
}
