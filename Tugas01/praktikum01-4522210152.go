package main

import (
	"fmt"
)

func main() {
	// Meminta input nama dari pengguna
	fmt.Print("Masukkan nama Anda: ")
	var nama string
	fmt.Scanln(&nama)

	// Meminta input usia dari pengguna
	fmt.Print("Masukkan usia Anda: ")
	var usia int
	fmt.Scanln(&usia)

	// Menampilkan informasi yang dimasukkan pengguna
	fmt.Printf("Nama: %s\nUsia: %d tahun\n", nama, usia)
}
