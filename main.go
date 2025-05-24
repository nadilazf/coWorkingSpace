package main

import "fmt"

type Ulasan struct {
	namaClient string
	rating     float64
	komentar   string
}

type RuangKerja struct {
	id           int
	nama         string
	lokasi       string
	hargaSewa    int
	ulasan       [10]Ulasan
	jumlahUlasan int
}

const MAKS_RUANG = 10

var daftarRuangKerja [MAKS_RUANG]RuangKerja
var jumlahRuang int = 0

func tambahRuangKerja() {
	if jumlahRuang >= MAKS_RUANG {
		fmt.Println("Jumlah ruang kerja sudah mencapai batas maksimum.")
		return
	}

	var ruang RuangKerja
	ruang.id = jumlahRuang + 1

	fmt.Println("Masukkan nama ruang kerja:")
	fmt.Scan(&ruang.nama)
	fmt.Println("Masukkan lokasi ruang kerja:")
	fmt.Scan(&ruang.lokasi)
	fmt.Println("Masukkan harga sewa ruang kerja:")
	fmt.Scan(&ruang.hargaSewa)

	daftarRuangKerja[jumlahRuang] = ruang
	jumlahRuang++

	fmt.Println("Ruang kerja berhasil ditambahkan ✅")
}

func editRuangKerja() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang bisa diedit.")
		return
	}

	var id int
	fmt.Print("Masukkan ID ruang kerja yang ingin diubah:")
	fmt.Scan(&id)

	index := -1
	for i := 0; i < jumlahRuang; i++ {
		if daftarRuangKerja[i].id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Ruang kerja dengan ID tersebut tidak ditemukan ❌")
		return
	}

	fmt.Println("Masukkan data baru:")
	fmt.Print("Nama baru: ")
	fmt.Scan(&daftarRuangKerja[index].nama)
	fmt.Print("Lokasi baru: ")
	fmt.Scan(&daftarRuangKerja[index].lokasi)
	fmt.Print("Harga sewa baru: ")
	fmt.Scan(&daftarRuangKerja[index].hargaSewa)

	fmt.Println("Data ruang kerja berhasil diperbarui ✅")
}

func hapusRuangKerja() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang bisa dihapus.")
		return
	}

	var id int
	fmt.Println("Masukkan ID ruang kerja yang ingin dihapus: ")
	fmt.Scan(&id)

	index := -1
	for i := 0; i < jumlahRuang; i++ {
		if daftarRuangKerja[i].id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Ruang kerja dengan ID tersebut tidak ditemukan ❌")
	}

	for i := index; i < jumlahRuang-1; i++ {
		daftarRuangKerja[i] = daftarRuangKerja[i+1]
	}

	jumlahRuang--

	fmt.Println("Ruang kerja berhasil dihapus ✅")
}

func tampilkanDaftarRuangKerja() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang terdaftar")
		return
	}

	fmt.Println("Daftar ruang kerja 🏢")
	for i := 0; i < jumlahRuang; i++ {
		ruang := daftarRuangKerja[i]

		if ruang.jumlahUlasan == 0 {
			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Harga Sewa: %d, Rating: Belum ada rating\n",
				ruang.id, ruang.nama, ruang.lokasi, ruang.hargaSewa)
		} else {
			var totalRating float64 = 0
			for j := 0; j < ruang.jumlahUlasan; j++ {
				totalRating += ruang.ulasan[j].rating
			}
			avgRating := totalRating / float64(ruang.jumlahUlasan)

			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
				ruang.id, ruang.nama, ruang.lokasi, ruang.hargaSewa, avgRating, ruang.jumlahUlasan)
		}
	}
}

func tampilkanUlasanRuangKerja() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang tersedia.")
		return
	}

	var id int
	fmt.Print("Masukkan ID ruang kerja yang ingin dilihat ulasannya: ")
	fmt.Scan(&id)

	index := -1
	for i := 0; i < jumlahRuang; i++ {
		if daftarRuangKerja[i].id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Ruang kerja tidak ditemukan ❌")
		return
	}

	ruang := daftarRuangKerja[index]

	if ruang.jumlahUlasan == 0 {
		fmt.Printf("Belum ada ulasan untuk ruang kerja '%s'.\n", ruang.nama)
		return
	}

	fmt.Printf("Ulasan untuk ruang kerja '%s':\n", ruang.nama)
	for i := 0; i < ruang.jumlahUlasan; i++ {
		ulasan := ruang.ulasan[i]
		fmt.Printf("- %s beri rating %.1f ⭐\n Komentar: %s\n", ulasan.namaClient, ulasan.rating, ulasan.komentar)
	}
}

func menuAdmin() {
	for {
		fmt.Println("\n===== MENU ADMIN 👤 =====")
		fmt.Println("1. Tambah ruang kerja ➕")
		fmt.Println("2. Edit ruang kerja ✏️")
		fmt.Println("3. Hapus ruang kerja 🗑️")
		fmt.Println("4. Tampilkan semua ruang kerja 🏢")
		fmt.Println("5. Lihat ulasan ruang kerja 📝")
		fmt.Println("6. Kembali ke menu utama 🔙")
		fmt.Println("7. Keluar 🚪")
		fmt.Print("Pilih opsi (1-7): ")

		var opsi int
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			tambahRuangKerja()
		case 2:
			editRuangKerja()
		case 3:
			hapusRuangKerja()
		case 4:
			tampilkanDaftarRuangKerja()
		case 5:
			tampilkanUlasanRuangKerja()
		case 6:
			return
		case 7:
			fmt.Println("Program selesai")
			exit()
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func menuClient() {
	for {
		fmt.Println("=== MENU CLIENT 👥 ===")
		fmt.Println("1. Tampilkan semua ruang kerja 🏢")
		fmt.Println("2. Kembali ke menu utama 🔙")
		fmt.Println("3. Keluar 🚪")
		fmt.Print("Pilih opsi (1-3): ")

		var opsiClient int
		fmt.Scan(&opsiClient)

		switch opsiClient {
		case 1:
			tampilkanDaftarRuangKerja()
		case 2:
			return
		case 3:
			fmt.Println("Program selesai")
			exit()
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func main() {
	for {
		fmt.Println("=== SELAMAT DATANG DI COWORKING APP 🏢📱 ===")
		fmt.Println("Pilih peran Anda:")
		fmt.Println("1. Admin 👤")
		fmt.Println("2. Client 👥")
		fmt.Println("3. Keluar 🚪")
		fmt.Print("Masukkan pilihan (1-3): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuAdmin()
		case 2:
			menuClient()
		case 3:
			fmt.Println("Sampai Jumpa 👋")
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func exit() {
	fmt.Println("Terima Kasih dan Sampai Jumpa 👋")
	panic("Exit 🚪")
}
