package main

import "fmt"

func main() {
	var usia int

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	if usia < 18 {
		fmt.Println("Anak-anak")
	} else {
		fmt.Println("Dewasa")
	}
}
