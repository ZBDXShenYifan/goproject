package main
import (
	"fmt"
	"net"
	"bufio"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	//功能
	//bufio.
}