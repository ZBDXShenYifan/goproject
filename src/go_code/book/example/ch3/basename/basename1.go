package main

import "fmt"

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a.go"))
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a/b.c.go"))
}

func basename(s string) string {
	
	for i := len(s) - 1; i >= 0; i-- {
		//将最后一个'/'之前的全部内容舍弃掉
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

		//保留最后一个'.'之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
