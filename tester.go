package main

import "fmt"

func mapping(slice []string) map[string]int {
	var i map[string]int = map[string]int{}

	for _, v := range data {
		i[v] = i[v] + 1
	}
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}
