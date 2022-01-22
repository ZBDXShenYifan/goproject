package main
import (
	"fmt"
	"time"
)


func main() {
	//1.获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//格式化日期时间
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// i := 0 
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100) //每0.1s打印一个数字
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	//Unix和UnixNano的使用
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) 
}