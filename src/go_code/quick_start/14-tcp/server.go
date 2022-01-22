package main

import (
	"fmt"
	"net" //做网络开发时，net包含有我们需要所有的方法和函数
)

func main() {
	fmt.Println("服务器开始监听...")
	//1. tcp表示使用网络协议是tcp, 在本地监听 8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close() //延时关闭listen

	//循环等待客户端来连接
	for {
		//等待客户端来连接
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accpet() suc conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr())
		}

		
	}
	fmt.Printf("listen=%v", listen)

}