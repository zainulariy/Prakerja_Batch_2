package main

import "fmt"

// Fungsi untuk menentukan apakah suatu bilangan merupakan bilangan prima atau bukan
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n, "merupakan bilangan prima")
	} else {
		fmt.Println(n, "bukan merupakan bilangan prima")
	}
}
