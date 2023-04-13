package main

import "fmt"

func main() {
	data := []string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}

	var i map[string]int = map[string]int{}

	for _, v := range data {
		i[v] = i[v] + 1
	}

	fmt.Println(i)
}
