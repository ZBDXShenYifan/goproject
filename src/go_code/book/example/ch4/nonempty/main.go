package main

import "fmt"

func main(){
	data := []string{"one", "", "three"}
	data = nonempty(data)
	fmt.Printf("%q\n", data)

	data2 := []string{"one", "", "three"}
	data2 = nonempty2(data)
	fmt.Printf("%q\n", data2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}