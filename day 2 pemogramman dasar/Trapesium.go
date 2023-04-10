package main

import "fmt"

func main() {

	var sa, sb, tinggi float64

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scan(&sa)

	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scan(&sb)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	luas := 0.5 * (sa + sb) * tinggi

	fmt.Println("Luas trapesium:", luas)
}
