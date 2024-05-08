package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + "蛋糕 " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//如果两个信道都关闭了，说明制作完成，结束程序
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("等待新蛋糕 ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println(" ... 草莓信道关闭")
			} else {
				fmt.Println("在草莓信道中收到一个新蛋糕。名为：", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... 巧克力信道关闭")
			} else {
				fmt.Println("在巧克力信道中收到一个新蛋糕。名为：", cakeName)
			}
		}
	}
}

func main() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend(choco_cs, "巧克力", 3) //制作3个巧克力蛋糕，然后发送
	go makeCakeAndSend(strbry_cs, "草莓", 3) //制作3个草莓蛋糕，然后发送

	//one cake receiver and packer
	go receiveCakeAndPack(strbry_cs, choco_cs) //收获

	//查看结果
	time.Sleep(2 * 1e9)
}
