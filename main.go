package main

import "fmt"

type Ulasan struct {
	namaClient string
	rating     float64
	komentar   string
}

type Fasilitas struct {
	namaFasilitas string
}

type RuangKerja struct {
	id              int
	nama            string
	lokasi          string
	fasilitas       [5]Fasilitas
	hargaSewa       int
	ulasan          [10]Ulasan
	jumlahUlasan    int
	jumlahFasilitas int
}

const MAKS_RUANG = 10
const MAKS_FASILITAS = 5

var daftarRuangKerja [MAKS_RUANG]RuangKerja
var jumlahRuang int = 0

func tambahRuangKerja() {
	if jumlahRuang >= MAKS_RUANG {
		fmt.Println("Jumlah ruang kerja sudah mencapai batas maksimum.")
		return
	}

	var ruang RuangKerja
	ruang.id = jumlahRuang + 1

	fmt.Print("Masukkan nama ruang kerja: ")
	fmt.Scan(&ruang.nama)

	fmt.Print("Masukkan lokasi ruang kerja: ")
	fmt.Scan(&ruang.lokasi)

	fmt.Print("Masukkan harga sewa ruang kerja: ")
	fmt.Scan(&ruang.hargaSewa)
	fmt.Scan() 

	fmt.Println("Masukkan fasilitas ruang kerja satu per satu (max 5), ketik 'selesai' untuk berhenti:")
	count := 0
	for count < MAKS_FASILITAS {
		fmt.Printf("Fasilitas %d: ", count+1)
		var fasilitasInput string
		fmt.Scan(&fasilitasInput)
		if fasilitasInput == "selesai" {
			break
		}
		if fasilitasInput != "" {
			ruang.fasilitas[count] = Fasilitas{namaFasilitas: fasilitasInput}
			count++
		}
	}
	ruang.jumlahFasilitas = count

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
	fmt.Print("Masukkan ID ruang kerja yang ingin diubah: ")
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
	fmt.Print("Masukkan nama ruang kerja: ")
	fmt.Scan(&daftarRuangKerja[index].nama)

	fmt.Print("Masukkan lokasi ruang kerja: ")
	fmt.Scan(&daftarRuangKerja[index].lokasi)

	fmt.Print("Masukkan harga sewa ruang kerja: ")
	fmt.Scan(&daftarRuangKerja[index].hargaSewa)
	fmt.Scan()

	fmt.Println("Masukkan fasilitas baru untuk ruang kerja (maks 5), ketik 'selesai' untuk berhenti:")
	count := 0
	for count < MAKS_FASILITAS {
		fmt.Printf("Fasilitas %d: ", count+1)
		var fasilitasInput string
		fmt.Scan(&fasilitasInput)
		if fasilitasInput == "selesai" {
			break
		}
		if fasilitasInput != "" {
			daftarRuangKerja[index].fasilitas[count] = Fasilitas{namaFasilitas: fasilitasInput}
			count++
		}
	}
	daftarRuangKerja[index].jumlahFasilitas = count

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

		fasilitasStr := ""
		for j := 0; j < ruang.jumlahFasilitas; j++ {
			fasilitasStr += ruang.fasilitas[j].namaFasilitas
			if j < ruang.jumlahFasilitas-1 {
				fasilitasStr += ", "
			}
		}

		if ruang.jumlahUlasan == 0 {
			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: Belum ada rating\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa)
		} else {
			var totalRating float64 = 0
			for j := 0; j < ruang.jumlahUlasan; j++ {
				totalRating += ruang.ulasan[j].rating
			}
			avgRating := totalRating / float64(ruang.jumlahUlasan)

			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa, avgRating, ruang.jumlahUlasan)
		}
	}
}

func tampilkanDaftarRuangKerjaClient() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang terdaftar")
		return
	}

	fmt.Println("Daftar ruang kerja 🏢")
	for i := 0; i < jumlahRuang; i++ {
		ruang := daftarRuangKerja[i]

		fasilitasStr := ""
		for j := 0; j < ruang.jumlahFasilitas; j++ {
			fasilitasStr += ruang.fasilitas[j].namaFasilitas
			if j < ruang.jumlahFasilitas-1 {
				fasilitasStr += ", "
			}
		}

		if ruang.jumlahUlasan == 0 {
			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: Belum ada rating\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa)
		} else {
			var totalRating float64 = 0
			for j := 0; j < ruang.jumlahUlasan; j++ {
				totalRating += ruang.ulasan[j].rating
			}
			avgRating := totalRating / float64(ruang.jumlahUlasan)

			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa, avgRating, ruang.jumlahUlasan)
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

func tambahUlasan(id int) {
	idx := cariIndexRuang(id)
	if idx == -1 {
		fmt.Println("Ruang kerja tidak ditemukan")
		return
	}
	if daftarRuangKerja[idx].jumlahUlasan >= 10 {
		fmt.Println("Slot ulasan penuh untuk ruang ini.")
		return
	}

	var nama, komentar string
	var rating float64
	fmt.Print("Nama Anda: ")
	fmt.Scan(&nama)
	ruang := &daftarRuangKerja[idx]
	for i := 0; i < ruang.jumlahUlasan; i++ {
		if ruang.ulasan[i].namaClient == nama {
			fmt.Println("Nama tersebut sudah pernah memberikan ulasan. Gunakan nama lain.")
			return
		}
	}
	fmt.Print("Rating (1-5): ")
	fmt.Scan(&rating)
	fmt.Print("Komentar: ")
	fmt.Scan(&komentar)

	ul := Ulasan{namaClient: nama, rating: rating, komentar: komentar}
	ruang.ulasan[ruang.jumlahUlasan] = ul
	ruang.jumlahUlasan++
	fmt.Println("Ulasan ditambahkan ✅")
}

func cariIndexRuang(id int) int {
	for i := 0; i < jumlahRuang; i++ {
		if daftarRuangKerja[i].id == id {
			return i
		}
	}
	return -1
}

func rataRating(r RuangKerja) float64 {
	if r.jumlahUlasan == 0 {
		return 0
	}
	var total float64
	for i := 0; i < r.jumlahUlasan; i++ {
		total += r.ulasan[i].rating
	}
	return total / float64(r.jumlahUlasan)
}

func editUlasan(id int) {
	idx := cariIndexRuang(id)
	if idx == -1 {
		fmt.Println("Ruang kerja tidak ditemukan")
		return
	}

	var nama string
	fmt.Print("Masukkan nama Anda (ulasan yg mau diedit): ")
	fmt.Scan(&nama)

	found := false
	for i := 0; i < daftarRuangKerja[idx].jumlahUlasan; i++ {
		if daftarRuangKerja[idx].ulasan[i].namaClient == nama {
			found = true
			fmt.Print("Rating baru: ")
			fmt.Scan(&daftarRuangKerja[idx].ulasan[i].rating)
			fmt.Print("Komentar baru: ")
			fmt.Scan(&daftarRuangKerja[idx].ulasan[i].komentar)
			fmt.Println("Ulasan diupdate ✅")
			break
		}
	}
	if !found {
		fmt.Println("Ulasan dengan nama tersebut tidak ditemukan")
	}
}

func hapusUlasan(id int) {
	idx := cariIndexRuang(id)
	if idx == -1 {
		fmt.Println("Ruang kerja tidak ditemukan")
		return
	}

	var nama string
	fmt.Print("Masukkan nama Anda (ulasan yg mau dihapus): ")
	fmt.Scan(&nama)

	for i := 0; i < daftarRuangKerja[idx].jumlahUlasan; i++ {
		if daftarRuangKerja[idx].ulasan[i].namaClient == nama {
			for j := i; j < daftarRuangKerja[idx].jumlahUlasan-1; j++ {
				daftarRuangKerja[idx].ulasan[j] = daftarRuangKerja[idx].ulasan[j+1]
			}
			daftarRuangKerja[idx].jumlahUlasan--
			fmt.Println("Ulasan dihapus ✅")
			return
		}
	}
	fmt.Println("Ulasan dengan nama tersebut tidak ditemukan ❌")
}

func menuDetailRuang(id int) {
	idx := cariIndexRuang(id)
	if idx == -1 {
		fmt.Println("Ruang kerja tidak ditemukan ❌")
		return
	}
	r := daftarRuangKerja[idx]
	fmt.Printf("\n== Detail Ruang ID %d ==\nNama: %s\nLokasi: %s\nHarga: %d\n",
		r.id, r.nama, r.lokasi, r.hargaSewa)
	fmt.Printf("Rating: %.2f (%d ulasan)\n\n", rataRating(r), r.jumlahUlasan)

	for {
		fmt.Println("1. Tambah ulasan ➕")
		fmt.Println("2. Edit ulasan ✏️")
		fmt.Println("3. Hapus ulasan 🗑️")
		fmt.Println("4. Kembali 🔙")
		fmt.Print("Pilih opsi (1-4): ")

		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			tambahUlasan(id)
		case 2:
			editUlasan(id)
		case 3:
			hapusUlasan(id)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func searchRuang(keyword string) {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang terdaftar.")
		return
	}

	idx := binarySearchNama(keyword)

	if idx != -1 {
		ruang := daftarRuangKerja[idx]

		fasilitasStr := ""
		for i := 0; i < ruang.jumlahFasilitas; i++ {
			fasilitasStr += ruang.fasilitas[i].namaFasilitas
			if i < ruang.jumlahFasilitas-1 {
				fasilitasStr += ", "
			}
		}

		if ruang.jumlahUlasan == 0 {
			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: Belum ada rating\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa)
		} else {
			var totalRating float64 = 0
			for i := 0; i < ruang.jumlahUlasan; i++ {
				totalRating += ruang.ulasan[i].rating
			}
			avgRating := totalRating / float64(ruang.jumlahUlasan)

			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
				ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa, avgRating, ruang.jumlahUlasan)
		}

	} else {
		fmt.Println("Mencari berdasarkan lokasi...")

		ketemu := false
		for i := 0; i < jumlahRuang; i++ {
			if daftarRuangKerja[i].lokasi == keyword {
				ruang := daftarRuangKerja[i]

				fasilitasStr := ""
				for j := 0; j < ruang.jumlahFasilitas; j++ {
					fasilitasStr += ruang.fasilitas[j].namaFasilitas
					if j < ruang.jumlahFasilitas-1 {
						fasilitasStr += ", "
					}
				}

				if ruang.jumlahUlasan == 0 {
					fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: Belum ada rating\n",
						ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa)
				} else {
					var totalRating float64 = 0
					for j := 0; j < ruang.jumlahUlasan; j++ {
						totalRating += ruang.ulasan[j].rating
					}
					avgRating := totalRating / float64(ruang.jumlahUlasan)

					fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
						ruang.id, ruang.nama, ruang.lokasi, fasilitasStr, ruang.hargaSewa, avgRating, ruang.jumlahUlasan)
				}
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Tidak ada ruang kerja dengan nama atau lokasi tersebut ❌")
		}
	}
}


func binarySearchNama(keyword string) int {
	left := 0
	right := jumlahRuang - 1

	for left <= right {
		mid := (left + right) / 2
		if daftarRuangKerja[mid].nama == keyword {
			return mid
		} else if keyword < daftarRuangKerja[mid].nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func filterFasilitas() {
	if jumlahRuang == 0 {
		fmt.Println("Belum ada ruang kerja yang terdaftar.")
		return
	}

	var fas string
	fmt.Print("Masukkan nama fasilitas yang dicari: ")
	fmt.Scan(&fas)

	ditemukan := false
	for i := 0; i < jumlahRuang; i++ {
		r := daftarRuangKerja[i]
		adaFasilitas := false
		for j := 0; j < r.jumlahFasilitas; j++ {
			if r.fasilitas[j].namaFasilitas == fas {
				adaFasilitas = true
				break
			}
		}

		if adaFasilitas {
			fasilitasStr := ""
			for j := 0; j < r.jumlahFasilitas; j++ {
				fasilitasStr += r.fasilitas[j].namaFasilitas
				if j < r.jumlahFasilitas-1 {
					fasilitasStr += ", "
				}
			}

			if r.jumlahUlasan == 0 {
				fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: Belum ada rating\n",
					r.id, r.nama, r.lokasi, fasilitasStr, r.hargaSewa)
			} else {
				var totalRating float64 = 0
				for k := 0; k < r.jumlahUlasan; k++ {
					totalRating += r.ulasan[k].rating
				}
				avgRating := totalRating / float64(r.jumlahUlasan)

				fmt.Printf("ID: %d, Nama: %s, Lokasi: %s, Fasilitas: [%s], Harga Sewa: %d, Rating: %.2f (%d ulasan)\n",
					r.id, r.nama, r.lokasi, fasilitasStr, r.hargaSewa, avgRating, r.jumlahUlasan)
			}

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada ruang kerja dengan fasilitas tersebut ❌")
	}
}


func submenuDaftarClient() {
	for {
		fmt.Println("\n--- Pilih opsi berikut ---")
		fmt.Println("1. Search nama / lokasi 🔍")
		fmt.Println("2. filter berdasarkan fasilitas")
		fmt.Println("3. Pilih ruang kerja (input ID) 📄")
		fmt.Println("4. Kembali ke menu sebelumnya 🔙")
		fmt.Print("Pilih opsi (1-4): ")

		var op int
		fmt.Scan(&op)

		switch op {
		case 1:
			var key string
			fmt.Print("Masukkan nama atau lokasi: ")
			fmt.Scan(&key)
			searchRuang(key)
		case 2:
			filterFasilitas()
		case 3:
			var id int
			fmt.Print("Masukkan ID ruang: ")
			fmt.Scan(&id)
			menuDetailRuang(id)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
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
			tampilkanDaftarRuangKerjaClient()
			submenuDaftarClient()
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
