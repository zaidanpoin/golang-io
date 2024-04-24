package main

import (
	"fmt"
	"os"
)

func main() {
	// Mengecek apakah file test.txt sudah ada
	_, err := os.Stat("test.txt")
	fmt.Println(err)
	if os.IsNotExist(err) {
		// Jika file test.txt tidak ada, maka buat file baru
		file, err := os.Create("test.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Printf("Success create file, %s\n", file.Name())
	} else {
		fmt.Println("File already exists")
	}

	// Membuat direktori test
	// err = os.Mkdir("test", os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Success create directory")

	// Membaca isi file test.txt
	fileContent, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))

	// Menulis data ke file test.txt
	data := "aku ganteng"
	err = os.WriteFile("test.txt", []byte(data), 0644)
	if err != nil {
		panic(err)
	}

	// Membuat file baru write-data.txt dan menulis data ke dalamnya
	file, err := os.Create("write-data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Menulis data ke file write-data.txt
	n, err := file.WriteString("ini hasil write data dari Golang")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success wrote %d bytes data\n", n)
}
