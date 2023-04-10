package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if n%7 == 0 {
		fmt.Println(n, "adalah kelipatan 7")
	} else {
		fmt.Println(n, "bukan kelipatan 7")
	}
}
