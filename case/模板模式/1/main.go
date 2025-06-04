package main

import "fmt"

// 定义模板接口
type BeverageTemplate interface {
	BoilWater()           // 固定步骤1
	Brew()                // 可变步骤2
	PourInCup()           // 固定步骤3
	AddCondiments()       // 可变步骤4
	WantCondiments() bool // 钩子方法(可选步骤)
}

// 定义模板结构体
type BeverageMaker struct {
	beverage BeverageTemplate
}

// 模板方法 - 定义算法骨架
func (b *BeverageMaker) MakeBeverage() {
	b.beverage.BoilWater()
	b.beverage.Brew()
	b.beverage.PourInCup()
	if b.beverage.WantCondiments() {
		b.beverage.AddCondiments()
	}
}

// 具体实现 - 咖啡
type Coffee struct{}

func (c *Coffee) BoilWater() {
	fmt.Println("将水煮沸到100摄氏度")
}

func (c *Coffee) Brew() {
	fmt.Println("用沸水冲泡咖啡粉")
}

func (c *Coffee) PourInCup() {
	fmt.Println("把咖啡倒入杯子中")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("加入糖和牛奶")
}

func (c *Coffee) WantCondiments() bool {
	return true // 默认加调料
}

// 具体实现 - 茶
type Tea struct{}

func (t *Tea) BoilWater() {
	fmt.Println("将水煮沸到80摄氏度")
}

func (t *Tea) Brew() {
	fmt.Println("用80度热水浸泡茶叶")
}

func (t *Tea) PourInCup() {
	fmt.Println("把茶倒入杯子中")
}

func (t *Tea) AddCondiments() {
	fmt.Println("加入柠檬")
}

func (t *Tea) WantCondiments() bool {
	var answer string
	fmt.Print("茶要加柠檬吗？(y/n): ")
	fmt.Scanln(&answer)
	return answer == "y"
}

func main() {
	// 制作咖啡
	fmt.Println("====== 制作咖啡 ======")
	coffeeMaker := &BeverageMaker{&Coffee{}}
	coffeeMaker.MakeBeverage()

	// 制作茶
	fmt.Println("\n====== 制作茶 ======")
	teaMaker := &BeverageMaker{&Tea{}}
	teaMaker.MakeBeverage()
}
