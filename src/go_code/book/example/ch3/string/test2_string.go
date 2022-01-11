package main

//判断某一个字符串是否为另一个的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//判断他是否为另一个字符串的后缀
func HasSuffix(s, substr string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//判断它是否为另一个串的子字符串
//Go包Contains函数使用了散列方法让搜索更高效
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true;
		}
	}
	return false
}