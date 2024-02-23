package main

import (
	"fmt"
)

type apexLegends struct {
	legends string
	skill   string
}

type legendsActor interface {
	jump()
	run()
	shoot()
}

type Ash struct{}
type Lifeline struct{}

func (hero apexLegends) getHeroNames_value() {
	hero.legends = "mirage"
	fmt.Println("非指针接受者方法 修改英雄名字为：", hero.legends)
}

func (hero *apexLegends) getHeroNames_pointer() {
	hero.legends = "mirage"
	fmt.Println("指针接受者方法 修改英雄名字为：", hero.legends)
}

func (hero Ash) jump() {
	fmt.Println("Ash jumped!")
}

func (hero Lifeline) jump() {
	fmt.Println("lifeline jumped!")
}

// 测试函数入口
func TESTFUNC(Q string, f func()) {
	fmt.Println("------------", Q, "------------------")
	f()
}

func structTest() {
	// 1.使用var初始化
	var hero_1 apexLegends
	hero_1.legends = "Ash"
	hero_1.skill = "flash across"
	fmt.Println(hero_1.legends)

	// 2.使用new初始化
	hero_2 := new(apexLegends)
	hero_2.legends = "lifeline"
	fmt.Println(hero_2.legends)

	// 3.使用字面量初始化
	hero_3 := &apexLegends{"pathfinder", ""}
	fmt.Println(hero_3.legends)
}

func structFuncTest() {
	hero := apexLegends{"wraith", ""}
	// 1.方法-非指针接受者 （传入的是一个拷贝）
	hero.getHeroNames_value()
	fmt.Println("方法返回后会发现没有修改：", hero.legends)
	// 2.方法-指针接受者
	hero_1 := &apexLegends{"wraith", ""}
	hero_1.getHeroNames_pointer()
	fmt.Println("方法返回后会发现原值被修改：", hero_1.legends)
}

func interfaceTest() {
	var hero Ash
	hero.jump()
	var hero1 Lifeline
	hero1.jump()
}

func main() {
	TESTFUNC("Lesson1:结构体定义及初始化", structTest)
	TESTFUNC("Lesson2:方法定义及调用", structFuncTest)
	TESTFUNC("Lesson3:接口定义及实现", interfaceTest)
}
