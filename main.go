package main

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
	RevCount int
}

const MAKS_RUANG = 5

var ruangKerjaList [MAKS_RUANG]RuangKerja
var jumlahRuang int = 0

func main(){
	
}