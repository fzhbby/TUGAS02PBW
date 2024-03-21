package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	// Menyimpan data mahasiswa dalam map
	dataMahasiswa := make(map[string]Mahasiswa)
	dataMahasiswa["4522210152"] = Mahasiswa{"Kaylatifa Zahabiy", "4522210152", "Teknik Informatika"}
	dataMahasiswa["4522210121"] = Mahasiswa{"Afiyah Nabilah", "4522210121", "Teknik Informatika"}
	dataMahasiswa["4522210150"] = Mahasiswa{"Apriani Simamora", "4522210150", "Teknik Informatika"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println("Daftar Nama Mahasiswa:")
	for npm, mhs := range dataMahasiswa {
		fmt.Println(npm, mhs.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM
	npm := "4522210152"
	mhs, ok := dataMahasiswa[npm]
	if ok {
		fmt.Println("Data Mahasiswa dengan NPM", npm, ":")
		fmt.Println(mhs)
	} else {
		fmt.Println("Data Mahasiswa dengan NPM", npm, "tidak ditemukan.")
	}
}
