package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello,你好" //golang的编码统一为utf-8,（ascii的字符占用一个字节，汉字占用3个字节）
	fmt.Println(len(str))

	str2 := "hello北京"

	str3 := []rune(str2) //rune的切片

	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c\n", str3[i])
	}

	n, err := strconv.Atoi("123hello")
	if err != nil {
		fmt.Println("转换错误", err)
	}else {
		fmt.Println("转换成的结果是", n)
	}

	str4 := strconv.Itoa(12345)
	fmt.Println(str4)

	var bytes = []byte("hello go")
	fmt.Println("bytes=", bytes)

	str5 := string([]byte{97, 98, 99})
	fmt.Println(str5)

	//10进制转2进制
	str6 := strconv.FormatInt(123,2)
	fmt.Println(str6)

	b := strings.Contains("seafood", "foo")
	fmt.Println(b)

	num := strings.Count("ceheese", "e")
	fmt.Println(num)

	b = strings.EqualFold("abc", "Abc")
	fmt.Println(b)
	fmt.Println("abc" == "Abc")

	index :=  strings.Index("NUL_abc", "abc")
	fmt.Println(index)

	index = strings.LastIndex("go golang", "go")
	fmt.Println(index)

	str = strings.Replace("go go hello", "go", "go语言", 2)
	fmt.Println(str)

	//n=-1表示全部替换 str本身并没有变化
	str = strings.Replace("go go hello", "go", "go语言", -1)
	fmt.Println(str)

	strArr := strings.Split("hello,world,ok", ",")
	fmt.Println(strArr)

	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}

	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Println(str)
	str = strings.ToUpper(str)
	fmt.Println(str)

	//省去两头的空格
	str = strings.TrimSpace(" tn a lone gopher strn ")
	fmt.Println(str)

	//将字符串两边指定的字符去掉(去掉空格和!)
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("%q\n",str)

	//TrimLeft TrimRight

	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Println(b)
	//判断是否以指定的字符串结束
	b= strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Println(b)
}