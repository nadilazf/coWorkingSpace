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
	ulasan [10]Ulasan
	jumlahUlasan int
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

func editRuangKerja(){
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

func hapusRuangKerja(){
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


func menuAdmin() {
	for {
		fmt.Println("\n===== MENU ADMIN 👤 =====")
		fmt.Println("1. Tambah ruang kerja ➕")
		fmt.Println("2. Tampilkan semua ruang kerja 🏢")
		fmt.Println("3. Edit ruang kerja ✏️")
		fmt.Println("4. Hapus ruang kerja 🗑️")
		fmt.Println("4. Kembali ke menu utama 🔙")
		fmt.Println("5. Keluar 🚪")
		fmt.Print("Pilih opsi (1-6): ")

		var opsi int
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			tambahRuangKerja()
		case 2:
			tampilkanDaftarRuangKerja()
		case 3:
			editRuangKerja()
		case 4:
			hapusRuangKerja()
		case 5:
			return
		case 6:
			fmt.Println("Program selesai")
			exit()
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func menuClient() {
	fmt.Println("=== MENU CLIENT 👥 ===")
	fmt.Println("(Masih dalam pengembangan ⚒️)")
	fmt.Printf("Kembali ke menu utama...\n")
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