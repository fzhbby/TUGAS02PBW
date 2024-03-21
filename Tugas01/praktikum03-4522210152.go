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

	// Menentukan kategori usia
	var kategori string
	if usia < 18 {
		kategori = "anak-anak"
	} else {
		kategori = "dewasa"
	}

	// Menampilkan pesan selamat datang dan kategori usia
	fmt.Printf("Selamat datang, %s Anda termasuk kategori %s.\n", nama, kategori)
}
