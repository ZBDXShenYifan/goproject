package main

func main() {
	a bool = s != "" && s[0] == 'x' //这是安全的
	b bool = s[0] == 'x' && s != "" //如果作用于空字符串，这是不安全的

	//布尔值无法转换为数值,需要显式用到if
	i := 0
	if b {
		i = 1
	}

	func btoi(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	//反向操作只需判断number != 0
}