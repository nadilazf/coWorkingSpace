package main

import "fmt"

type Ulasan struct {
	NamaClient string
	Rating float64
	Komentar string
}

type RuangKerja struct {
	ID int
	Nama string
	Lokasi string
	HargaSewa int
	DaftarUlasan [5]Ulasan
	jumlahUlasan int
}

const MAKS_RUANG = 5
const MAKS_ULASAN = 5

var daftarRuangKerja [MAKS_RUANG]RuangKerja
var jumlahRuang int = 0

func tambahUlasan(idRuang int, namaClient string, rating float64, komentar string) {
	indeksRuang := -1
	for i := 0; i < jumlahRuang; i++ {
		if daftarRuangKerja[i].ID == idRuang {
			indeksRuang = i
			break
		}
	}

	if indeksRuang == -1 {
		fmt.Println("Ruang kerja tidak ditemukan.")
		return
	}

	ruang := &daftarRuangKerja[indeksRuang]
	if ruang.jumlahUlasan < MAKS_ULASAN {
		ruang.DaftarUlasan[ruang.jumlahUlasan] = Ulasan{
			NamaClient: namaClient,
			Rating: rating,
			Komentar: komentar,
		}
		ruang.jumlahUlasan++
		fmt.Println("Ulasan berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah ulasan maksimal sudah tercapai.")
	}
}

func main(){

}