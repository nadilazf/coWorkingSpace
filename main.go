package main

import "fmt"

type Ulasan struct {
	namaClient string
	rating     float64
	komentar   string
}

type RuangKerja struct {
	id        int
	nama      string
	lokasi    string
	hargaSewa int
}

const MAKS_RUANG = 10

var daftarRuangKerja [MAKS_RUANG]RuangKerja
var jumlahRuang int = 0

func tambahRuangKerja(){
	if jumlahRuang >= MAKS_RUANG{
		fmt.Println("Jumlah ruang kerja sudah mencapai batas maksimum.")
		return
	}

	var ruang RuangKerja
	fmt.Println("Masukkan ID ruang kerja:")
	fmt.Scan(&ruang.id)
	fmt.Println("Masukkan nama ruang kerja:")
	fmt.Scan(&ruang.nama)
	fmt.Println("Masukkan lokasi ruang kerja:")
	fmt.Scan(&ruang.lokasi)
	fmt.Println("Masukkan harga sewa ruang kerja:")
	fmt.Scan(&ruang.hargaSewa)

	daftarRuangKerja[jumlahRuang] = ruang
	jumlahRuang++
	
	fmt.Println("Ruang kerja berhasil ditambahkan.")
}

func tampilkanDaftarRuangKerja(){
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang terdaftar.")
		return
	}

	fmt.Println("Daftar ruang kerja:")
	for i := 0; i < jumlahRuang; i++ {
		ruang := daftarRuangKerja[i]
		fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Harga Sewa: %d\n", ruang.id, ruang.nama, ruang.lokasi, ruang.hargaSewa)
	}
}

func main() {
	tambahRuangKerja()
	tampilkanDaftarRuangKerja()
}
