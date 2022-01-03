package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一个指针传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	fmt.Println("------")

	delete(cityMap, "China")

	printMap(cityMap)

	fmt.Println("------")

	cityMap["USA"] = "DC"

	printMap(cityMap)

	fmt.Println("------")

	ChangeValue(cityMap)

	printMap(cityMap)
}