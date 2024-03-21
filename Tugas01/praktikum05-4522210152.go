package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// Menampilkan informasi yang dimasukkan pengguna
	fmt.Printf("Nama: %s\nUsia: %d tahun\n", nama, usia)
}
