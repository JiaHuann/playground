package main

import (
	"fmt"

	"rsc.io/quote"
)

// 所有小知识点被xxxTEST包装，最后在主函数通过TESTFUNC
// arg1:知识点
// arg2:知识点对应的实现函数（测试函数）
func TESTFUNC(Q string, f func()) {
	fmt.Println("------------", Q, "------------------")
	f()
}

func printString(a, b, c string) (string, string, string) {
	return a, b, c
}

func deferTEST() {
	defer fmt.Println("[defer] I am defered,actually i am up to you.")
	print("Q[defer]: whereis my friend? ")
}

func defineFuncTEST() {
	a, b, c := printString("aaa", "bbb", "ccc")
	println("Q[Function define and call]:", a, b, c)
}

func ifLoopTEST() {
	if1Sum := 0
	if if1Sum == 0 {
		print("Q[if]:if1Sum is 0\n")
	}

	if if1Sum++; if1Sum == 1 {
		print("Q[more-with-if]:if1Sum is 1\n")
	}

	loop1Sum := 0
	for i := 0; i < 10; i++ {
		loop1Sum++
	}
	println("Q[for-loop]:", loop1Sum)

	loop2Sum := 0
	for loop2Sum < 100 {
		loop2Sum++
	}
	println("Q[while-loop]:", loop2Sum)
}

func sliceArrayTEST() {
	KeyValue := []string{
		"key is following:",
		"key1",
		"key2",
		"value is following",
		"1111",
		"22222",
	}
	fmt.Println("raw:", KeyValue)
	Key := KeyValue[1:3]
	Value := KeyValue[4:]
	fmt.Println("Key Slice:", Key)
	fmt.Println("Value Slice:", Value)
}

func main() {
	fmt.Println(quote.Go())
	//--------------练习：FUNCTION---------------------//
	TESTFUNC("Q[define and call function]", defineFuncTEST)
	//--------------练习：判断和循环--------------------//
	TESTFUNC("Q[if and loop]", ifLoopTEST)
	//----------------练习：匿名变量-----------------------//
	_, retSecond, _ := printString("a", "b", "c")
	println("Q[annomoyous var]:", retSecond)
	//---------------练习： defer TEST--------------------//
	TESTFUNC("[Q]defer", deferTEST)
	//---------------练习：切片--------------------------------//
	TESTFUNC("Q[Slice Array]", sliceArrayTEST)

}
