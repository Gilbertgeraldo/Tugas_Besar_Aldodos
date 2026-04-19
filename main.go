package main

import (
	"fmt"
	"math"
)

type admin struct {
	Username string
	Password string
}

type nasabah struct {
	username string
	password string
	NamaLengkap string
}

type Pinjaman struct {
	Nama string
	JumlahPinjaman float64
	Tenor int
	Status string
	Bunga float64
	BayaranPerbulan float64
	sisaPinjam int
	sudahBayar int
	SchemaBunga string
}

var dataAdmin [3]admin
var akunNasabah [1000]nasabah
var dataPeminjam [1000]Pinjaman
var counterAdmin int = 0
var counterPeminjam int = 0

func ClearScreen() {
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func garis1() {
	fmt.Print("=================================================================================")
}

func garis2() {
	fmt.Println("-------------------------------------------------------------------------------")
}

func Adminmenu(admin string) {
	fmt.Printf("TANGGAL : 19-04-2026\n")
	garis1()
	fmt.Printf("|			ADMIN MENU			")
	fmt.Printf("| %-76s |\n","1.Lihat daftar nasabah")
	fmt.Printf("| %-76s |\n","2.Lihat daftar pengajuan pinjaman")
	fmt.Printf("| %-76s |\n","3.Hapus data nasabah")
	fmt.Printf("| %-76s |\n","4.Hapus data nasabah")
	garis1()
}

func Usermenu(user string) {
	fmt.Printf("TANGGAL: 19-04-2026 \n")
	garis1()
	fmt.Printf("| Selamat datang, %-60s |\n", user)
	fmt.Printf("| %-76s |\n", "Apa yang bisa kami bantu hari ini ? ")
	fmt.Printf("| %-76s | \n","Ajukan peminjaman")
	fmt.Printf("| %-76s |\n","Lihat total pinjaman")
	fmt.Printf("| %-76s |\n","")
	fmt.Print("")
	garis1()
}

// UNTUK MENDAFTARKAN ADMIN DENGAN MAKSIMAL ADMIN ADALAH 3
// func regisAdmin(data *[3]admin, i int) {
// 	fmt.Println()
// 	fmt.Printf("=%-10s Buat Akun =%-10s", " ", " ")
// 	fmt.Print("Nama : ")	
// 	fmt.Scan(&data[i].Username)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&data[i].Password)
// 	fmt.Printf("=====================================")
// 	fmt.Print("YEAYYY!DATA KAMU BERHASIL DIBUAT!!")
// }

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

func loginAdmin() bool {
	var username, password string

	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	for _, v := range dataAdmin {
		if username == v.Username && password == v.Password {
			fmt.Print("Password dan username anda benar!")
			Adminmenu(v.Username)
			return true
		}
	}
	return false
}

func loginPeminjam() bool {
	var usn,pwd string

	fmt.Print("Username : ")
	fmt.Scan(&usn)
	fmt.Print("Password : ")
	fmt.Scan(&pwd)

	for _,v := range akunNasabah {
		if v.username == usn && v.password == pwd {
			return true
		}
	}
	return false
}

func Register() {
	ClearScreen()
	var pilihan string
	garis1()
	fmt.Println("--------- DAFTAR AKUN BARU ---------")
	fmt.Println("1.Daftar Sebagai admin")
	fmt.Println("2.Register sebagai peminjam")
	garis1()
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
				Status: "Menunggu",
				Tenor: 0,
				Bunga: 0,
				sudahBayar: 0,
				BayaranPerbulan: 0,
				sisaPinjam: 0,
				JumlahPinjaman: 0,
			}
			dataPeminjam[counterPeminjam] = nasabahBaru
			counterPeminjam++
			fmt.Println("Registrasi peminjaman berhasil")
		}
	}
}

// INI RUMUS UNTUK KALKULASI BUNGA

//Bunga variabel anuitas
// C = p * r(1 + r)^n / ((1 + r)^n - 1)
func HitungAnuitas(pokok,bunga float64,tenor int) (bayaran,totalBunga,totalBayar float64) {
	r := (bunga / 100.0) / 12.0
	if r == 0 {
		bayaran = math.Round((pokok/float64(tenor)) * 100) / 100
		return
	}
	rn := math.Pow(1 + r, float64(tenor))
	bayaran = math.Round(pokok * (r * rn/(rn-1))*100)/100
	totalBayar = math.Round()
	totalBunga = math.Round()
}

func tambahData(data *[1000]Pinjaman,nama string, jumlah,tenor int,status string,skema string,bunga float64, i int) {
	data[i].Nama = nama
	data[i].JumlahPinjaman = float64(jumlah)
	data[i].Tenor = tenor
	data[i].Status = status
	data[i].SchemaBunga = skema
	data[i].Bunga = bunga
}


func main() {
	fmt.Println("Program Sistem Pinjaman")
	// Tambahkan logika utama di sini
}