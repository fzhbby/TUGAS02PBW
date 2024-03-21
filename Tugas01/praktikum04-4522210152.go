package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Menggunakan bufio untuk input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Input nama
	fmt.Print("Masukkan nama Anda: ")
	nama, _ := reader.ReadString('\n')

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	usiaStr, _ := reader.ReadString('\n')
	usia, _ := strconv.Atoi(usiaStr)

	// Menentukan kategori usia
	var kategori string
	if usia < 18 {
		kategori = "anak-anak"
	} else {
		kategori = "dewasa"
	}

	// Menampilkan pesan selamat datang dan kategori usia
	fmt.Printf("Selamat datang, %s Anda termasuk kategori %s.\n", nama, kategori)

	// Menggunakan os untuk berinteraksi dengan sistem operasi
	// Misalnya, membuat direktori dengan nama pengguna
	err := os.Mkdir(nama, 0755)
	if err != nil {
		fmt.Println("Gagal membuat direktori:", err)
	} else {
		fmt.Println("Direktori berhasil dibuat.")
	}
}
