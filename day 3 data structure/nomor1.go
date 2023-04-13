package main

import "fmt"

func ArrayMerge(arrayA, arrayb []string) []string {
	var res []string
	var merge []string

	merge = append(arrayA, arrayb...)
	inRes := make(map[string]bool)
	for _, v := range merge {
		if _, y := inRes[v]; !y {
			inRes[v] = true
			res = append(res, v)
		}
	}

	return res
}

func main() {

	// test case

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []
}
