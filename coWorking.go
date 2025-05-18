package main

import "fmt"

/* ---------- TIPE DATA ---------- */

type Review struct {
	name    string
	comment string
	rating  float64
}

type CoWorkingSpace struct {
	id            int
	name          string
	location      string
	price         float64
	facilityCount int        // jumlah fasilitas yang terisi
	facilities    [10]string // maks 10 fasilitas per space
}

/* ---------- KONFIGURASI ---------- */

const maxSpaces = 100 // kapasitas maksimum co‑working space 100

var (
	coWorkingSpaces [maxSpaces]CoWorkingSpace
	totalSpaces     int = 0 // total slot array yang terisi
	nextID          int = 1 // ID unik
)

/* ---------- MAIN ---------- */

func main() {
	for {
		fmt.Println("===== Co‑Working Space App =====")
		fmt.Println("1. Masuk sebagai Admin")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			adminMenu()
		case 2:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

/* ---------- MENU ADMIN ---------- */

func adminMenu() {
	for {
		fmt.Println("\n===== Menu Admin =====")
		fmt.Println("1. Tambah Co‑Working Space")
		fmt.Println("2. Lihat Daftar Co‑Working Space")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			addCoWorkingSpace()
		case 2:
			listCoWorkingSpaces()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak validddd.")
		}
	}
}

/* ---------- TAMBAH DATA ---------- */

func addCoWorkingSpace() {
	if totalSpaces >= maxSpaces {
		fmt.Println("Data sudah penuh.")
		return
	}

	var nama, lokasi string
	var harga float64
	var nFac int

	fmt.Print("\nNama: ")
	fmt.Scan(&nama)
	fmt.Print("Lokasi: ")
	fmt.Scan(&lokasi)
	fmt.Print("Harga sewa: ")
	fmt.Scan(&harga)

	fmt.Print("Jumlah fasilitas (maks 10): ")
	fmt.Scan(&nFac)
	if nFac > 10 {
		nFac = 10
		fmt.Println("Dibatasi maksimal 10 fasilitas.")
	}

	var facArr [10]string
	for i := 0; i < nFac; i++ {
		fmt.Printf("Fasilitas #%d: ", i+1)
		fmt.Scan(&facArr[i])
	}

	coWorkingSpaces[totalSpaces] = CoWorkingSpace{
		id:            nextID,
		name:          nama,
		location:      lokasi,
		price:         harga,
		facilityCount: nFac,
		facilities:    facArr,
	}
	totalSpaces++
	nextID++

	fmt.Println("Data berhasil ditambahkan.")
}

/* ---------- LIST DATA ---------- */

func listCoWorkingSpaces() {
	if totalSpaces == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Println("\n=== Daftar Co‑Working Space ===")
	for i := 0; i < totalSpaces; i++ {
		s := coWorkingSpaces[i]
		fmt.Printf("ID: %d | Nama: %s | Lokasi: %s | Harga: %.2f\n",
			s.id, s.name, s.location, s.price)

		fmt.Print("  Fasilitas: ")
		for j := 0; j < s.facilityCount; j++ {
			fmt.Print(s.facilities[j])
			if j < s.facilityCount-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
