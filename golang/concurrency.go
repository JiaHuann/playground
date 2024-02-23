// 练习golang并发相关内容
package main

// import (
// 	"fmt"
// )

// func f(msg string) {
// 	fmt.Println(msg)
// }

// func main() {
// 	go f("hi,this is go routine")
// }

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 假设出现了错误
			continue
		}
		handleConn(conn) // 处理连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 连接关闭，则停止执行
		}
		time.Sleep(1 * time.Second)
	}
}
