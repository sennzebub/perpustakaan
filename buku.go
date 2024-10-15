package main

import (
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Amount int
}

type User struct {
	Username string
	Password string
}

type BorrowingHistory struct {
	History []string
}

var (
	books = []Book{
		{"Pemrograman", 10},
		{"Film", 5},
		{"Printing", 20},
	}

	user = User{"Akhdan", "2406499071"}

	history = BorrowingHistory{}
)

func main() {
	fmt.Println("Selamat datang di Perpustakaan Vokasi!")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	var password string
	fmt.Scanln(&password)

	if username != user.Username || password != user.Password {
		fmt.Println("Username atau password salah.")
		return
	}

	for {
		fmt.Println("\nMenu Program:")
		fmt.Println("1. Lihat informasi pengguna program")
		fmt.Println("2. Lihat daftar buku")
		fmt.Println("3. Tambah daftar buku")
		fmt.Println("4. Tambah peminjaman buku")
		fmt.Println("5. Histori peminjaman buku")
		fmt.Println("6. Keluar dari program")

		fmt.Print("Masukkan pilihan: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			LihatInformasiPenggunaProgram()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			KeluarDariProgram()
		default:
			fmt.Println("Pilihan tidak ada.")
		}
	}
}

func LihatInformasiPenggunaProgram() {
	fmt.Println(" ")
	fmt.Println("Informasi pengguna program:")
	fmt.Println(" ")
	fmt.Println("Username:", user.Username)
	fmt.Println("NPM:", user.Password)
	fmt.Println("Jenis kelamin: Laki-laki")
	fmt.Println("Makanan favorit: Nasi Kebuli")
	fmt.Println("Minuman favorit: Jus Terong Belanda")
}

func LihatDaftarBuku() {
	fmt.Println("Daftar buku:")
	for i, book := range books {
		fmt.Printf("%d. Nama Buku: %s\n   Jumlah: %d\n", i+1, book.Title, book.Amount)
	}
}

func TambahDaftarBuku() {
	fmt.Println("Tambah Daftar Buku")
	fmt.Print("Nama Buku: ")
	var title string
	fmt.Scanln(&title)
	fmt.Print("Jumlah: ")
	var amount int
	fmt.Scanln(&amount)

	books = append(books, Book{title, amount})
	fmt.Println("Berhasil ditambahkan!")
}

func TambahPeminjamanBuku() {
	fmt.Println("Daftar Buku")
	for i, book := range books {
		if book.Amount > 0 {
			fmt.Printf("%d. %s\n", i+1, book.Title)
			fmt.Printf("   Jumlah: %d buku tersedia\n", book.Amount)
		}
	}

	fmt.Print("Masukkan judul buku yang ingin dipinjam: ")
	var title string
	fmt.Scanln(&title)

	for i, book := range books {
		if book.Title == title {
			if book.Amount <= 0 {
				fmt.Println("Buku tidak tersedia untuk dipinjam.")
				return
			}

			fmt.Print("Masukkan jumlah buku yang ingin dipinjam: ")
			var borrowAmount int
			fmt.Scanln(&borrowAmount)

			if borrowAmount <= 0 {
				fmt.Println("Jumlah peminjaman tidak valid.")
				return
			}

			if borrowAmount > book.Amount {
				fmt.Println("Jumlah peminjaman melebihi stok buku.")
				return
			}

			book.Amount -= borrowAmount
			history.History = append(history.History, fmt.Sprintf("Peminjaman %s sebanyak %d", title, borrowAmount))
			fmt.Println("Peminjaman buku berhasil!")
			books[i] = book
			return
		}
	}

	fmt.Println("Buku tidak ditemukan.")
}

func HistoriPeminjamanBuku() {
	fmt.Println("Histori peminjaman buku:")
	for _, history := range history.History {
		fmt.Println(history)
	}
}

func KeluarDariProgram() {
	fmt.Println("Terima kasih telah menggunakan program peminjaman buku Vokasi!")
	os.Exit(0)
}
