package main

import (
	"fmt"
	"math"
)

type admin struct {
	Username string
	Password string
}

type Pinjaman struct {
	Username string
	Password string
	Nama string
	JumlahPinjaman int
	Tenor int
	Status string
	Bunga float64
	BayaranPerbulan float64
	sisaPinjam int
	sudahBayar bool
}

var dataAdmin [3]admin
var dataPeminjam [1000]Pinjaman
var counterAdmin int = 0
var counterPeminjam int = 0

func ClearScreen() {
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func menuAdmin(user string) {
	fmt.Printf("TANGGAL: 19-04-2026 \n")
	fmt.Println("-----------------------------------------------")
	fmt.Println("============= ADMIN MENU ======================")
	fmt.Println("	1.Lihat data nasabah	")
	fmt.Println("	2.Lihat ajuan peminjaman	")
	fmt.Println()
}

func menu(user string) {
	fmt.Printf("TANGGAL: 19-04-2026 \n")
	fmt.Println("================================================================================")
	fmt.Printf("| Selamat datang, %-60s |\n", user)
	fmt.Printf("| %-76s |\n", "Apa yang bisa kami bantu hari ini ? ")
	fmt.Println("================================================================================")
}

// UNTUK MENDAFTARKAN ADMIN DENGAN MAKSIMAL ADMIN ADALAH 3
func regisAdmin(data *[3]admin, i int) {
	fmt.Println()
	fmt.Printf("=%-10s Buat Akun =%-10s", " ", " ")
	fmt.Print("Nama : ")
	fmt.Scan(&data[i].Username)
	fmt.Print("Password : ")
	fmt.Scan(&data[i].Password)
	fmt.Printf("=====================================")
	fmt.Print("YEAYYY!DATA KAMU BERHASIL DIBUAT!!")
}

// CEK APAKAH ADMIN ADA ? 
// func cekDataAdmin(usn string, pw string) bool {
// 	for _,v := range dataAdmin {
// 		if usn == v.Username && pw == v.Password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func cekDataNasabah(usn string,pw string) bool {
// 	for _,v := range dataPeminjam {
// 		if usn == v.Nama && pw == v.Password {
// 			return true
// 		}
// 	}
// 	return false
// }

func loginAdmin() (string, bool) {
	var username, password string

	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	for _, v := range dataAdmin {
		if username == v.Username && password == v.Password {
			fmt.Print("Password dan username anda benar!")
			menu(v.Username)

			return v.Username, true
		}
	}
	return "username atau password kamu salah!", false
}

func loginPeminjam() (string, bool) {
	var username, password string

	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	for _, v := range dataPeminjam {
		if username == v.Username && password == v.Password {
			fmt.Print("Password anda benar!")
			menu(v.Username)
			return v.Username, true
		}
	}
	return "username atau password anda salah!", false
}

func Register() {
	ClearScreen()
	var pilihan string
	fmt.Println("--------- DAFTAR AKUN BARU ---------")
	fmt.Println("1.Daftar Sebagai admin")
	fmt.Println("2.Register sebagai peminjam")
	fmt.Scan(&pilihan)

	var usn string
	fmt.Print("Harap masukan username anda : ")
	fmt.Scan(&usn)

	switch pilihan {
	case "1":
		found := false
		for _, v := range dataAdmin {
			if v.Username == usn {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("\nUsername : %s sudah terdaftar!!", usn)
			loginAdmin()
		} else {
			var pw string
			fmt.Print("Harap masukan password anda : ")
			fmt.Scan(&pw)
			dataAdmin[counterAdmin] = admin{Username: usn, Password: pw}
			counterAdmin++
			fmt.Println("registrasi admin berhasil")
		}

	case "2":
		found := false
		for _, v := range dataPeminjam {
			if v.Nama == usn {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("\nUsername : %s sudah terdaftar!!", usn)
			loginPeminjam()
		} else {
			nasabahBaru := Pinjaman{
				Nama:   usn,
				Status: "Calon Peminjam",
				// Masih mikir wkwkkwk
			}
			dataPeminjam[counterPeminjam] = nasabahBaru
			counterPeminjam++
			fmt.Println("Registrasi peminjaman berhasil")
		}
	}
}

func main() {
	fmt.Println("Program Sistem Pinjaman")
	// Tambahkan logika utama di sini
}