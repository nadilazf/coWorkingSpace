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




func main() {
	fmt.Println("")
}
