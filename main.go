package main

import "fmt"

type admin struct {
	Username string
	Password string
}

type peminjaman struct {
	Nama string
	Jumlah int
	Tenor int
	Status string
	Bunga float64
	BayaranPerbulan float64
}

var dataAdmin [3]admin
var dataPeminjam [1000]peminjaman

//UNTUK MENDAFTARKAN ADMIN DENGAN MAKSIMAL ADMIN ADALAH 3
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

//CEK APAKAH ADMIN ADA ? 
func cekDataAdmin(usn string, password string) bool {
	for _,v := range dataAdmin {
		if usn == v.Username && password == v.Password {
			return true
		}
	}
	return false
}

//MENGHITUNG BUNGA BERDASARKAN STATUS
func hitungBunga(indeks int) {
		if dataPeminjam[indeks].Status == "Lunas" || dataPeminjam[indeks].Status == "Dalam Tenor" {
			dataPeminjam[indeks].Bunga = float64(dataPeminjam[indeks].Jumlah) * 0.02
		}else if dataPeminjam[indeks].Status == "Lunas Tenor" {
			dataPeminjam[indeks].Bunga == float64(dataPeminjam[indeks].Jumlah * 0.03)
		}
}

func hitungBayaranPerbulan(indeks int) {
	if dataPeminjam[indeks].Tenor > 0 {
		dataPeminjam[indeks].BayaranPerbulan = (float64(dataPeminjam[indeks].Jumlah) + dataPeminjam[indeks]/Bunga)
	}
}
func main() {

}