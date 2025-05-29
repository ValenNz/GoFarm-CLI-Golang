package main

import (
	"fmt"
	"strings"
	"time"
)

const Maks = 100

/* Struct */
type Sapi struct {
	id     string
	nama   string
	berat  int
	status string
}
type Pakan struct {
	nama  string
	harga int
	stok  int
}
type Kesehatan struct {
	idSapi   string
	tanggal  string
	diagnosa string
}

/* Variabel Global */
var dataSapi [Maks]Sapi
var jumlahSapi int
var idTerakhir int

var dataPakan [Maks]Pakan
var jumlahPakan int

var dataKesehatan [Maks]Kesehatan
var jumlahKesehatan int

/* ========== MENU UTAMA DAN SUBMENU ========== */
func menuUtama() {
	var pilihan int
	for {
		fmt.Println("\n========= SISTEM PETERNAKAN SAPI POTONG =========")
		fmt.Println("[1] Manajemen Sapi")
		fmt.Println("[2] Manajemen Pakan")
		fmt.Println("[3] Manajemen Kesehatan")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu utama: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuManajemenSapi()
		} else if pilihan == 2 {
			menuManajemenPakan()
		} else if pilihan == 3 {
			menuManajemenKesehatan()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan sistem ini.")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuManajemenSapi() {
	var pilihan int
	for {
		fmt.Println("\n--- Manajemen Sapi ---")
		fmt.Println("[1] Tambah Data Sapi")
		fmt.Println("[2] Tampilkan Semua Data Sapi")
		fmt.Println("[3] Urutkan Sapi berdasarkan Berat")
		fmt.Println("[4] Cari Sapi berdasarkan ID")
		fmt.Println("[5] Cari Sapi berdasarkan Status")
		fmt.Println("[6] Edit Data Sapi")
		fmt.Println("[7] Hapus Data Sapi")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahSapi()
		} else if pilihan == 2 {
			tampilkanSapi()
		} else if pilihan == 3 {
			urutkanSapiByBerat()
		} else if pilihan == 4 {
			cariSapiByID()
		} else if pilihan == 5 {
			cariSapiByStatus()
		} else if pilihan == 6 {
			editSapi()
		} else if pilihan == 7 {
			hapusSapi()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuManajemenPakan() {
	var pilihan int
	for {
		fmt.Println("\n--- Manajemen Pakan ---")
		fmt.Println("[1] Tambah Data Pakan")
		fmt.Println("[2] Tampilkan Semua Pakan")
		fmt.Println("[3] Urutkan Pakan berdasarkan Harga")
		fmt.Println("[4] Update Stok Pakan")
		fmt.Println("[5] Hapus Pakan")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahPakan()
		} else if pilihan == 2 {
			tampilkanPakan()
		} else if pilihan == 3 {
			urutkanPakanByHarga()
		} else if pilihan == 4 {
			updateStokPakan()
		} else if pilihan == 5 {
			hapusPakan()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuManajemenKesehatan() {
	var pilihan int
	for {
		fmt.Println("\n--- Manajemen Kesehatan ---")
		fmt.Println("[1] Tambah Catatan Kesehatan")
		fmt.Println("[2] Tampilkan Semua Catatan Kesehatan")
		fmt.Println("[3] Lihat Riwayat Kesehatan Sapi")
		fmt.Println("[4] Lihat Riwayat Kesehatan Sapi")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahKesehatan()
		} else if pilihan == 2 {
			tampilkanKesehatan()
		} else if pilihan == 3 {
			editKesehatan()
		} else if pilihan == 4 {
			hapusKesehatan()
		} else if pilihan == 5 {
			lihatRiwayatKesehatanSapi()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

/* ========== FUNGSI MANAJEMEN SAPI ========== */

/* Create Data Sapi*/
func tambahSapi() {
	if jumlahSapi >= Maks {
		fmt.Println("Kapasitas data sapi penuh.")
		return
	}

	var sapi Sapi
	idTerakhir++
	sapi.id = fmt.Sprintf("S%03d", idTerakhir)

	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == sapi.id {
			fmt.Println("Terjadi duplikasi ID, coba lagi.")
			return
		}
	}

	fmt.Printf("ID otomatis: %s\n", sapi.id)
	fmt.Print("Nama: ")
	fmt.Scan(&sapi.nama)

	for {
		fmt.Print("Berat (kg): ")
		fmt.Scan(&sapi.berat)
		if sapi.berat <= 0 {
			fmt.Println("Berat tidak valid")
		} else {
			break
		}
	}

	if sapi.berat < 100 {
		sapi.status = "pedet"
	} else if sapi.berat < 250 {
		sapi.status = "grower"
	} else {
		sapi.status = "siap_potong"
	}

	dataSapi[jumlahSapi] = sapi
	jumlahSapi++
	fmt.Println("Data sapi berhasil ditambahkan.")
}

/* Read Data Sapi*/
func tampilkanSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	fmt.Println("\nDaftar Sapi:")
	fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")
	for i := 0; i < jumlahSapi; i++ {
		fmt.Printf("%-8s %-15s %-8d %-15s\n",
			dataSapi[i].id, dataSapi[i].nama, dataSapi[i].berat, dataSapi[i].status)
	}
}

/* Update Data Sapi*/
func editSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	/* Program Insertion Sort */
	/*
		membesar
		Bandingkan key dengan elemen sebelumnya, lalu geser elemen yang lebih besar ke kanan,
		hingga menemukan posisi yang tepat untuk menyisipkan key.
	*/
	urutkanSapiByID := func() {
		for i := 1; i < jumlahSapi; i++ {
			key := dataSapi[i] // Menyimpan data sapi yang akan dibandingkan (disisipkan ke tempat yang sesuai).
			j := i - 1         // inisialisasi variabel pembanding yang menunjuk ke elemen sebelum key.
			for j >= 0 && strings.ToUpper(dataSapi[j].id) > strings.ToUpper(key.id) {
				/*
					Periksa apakah ID sapi sebelumnya lebih besar dari key.id.
					Gunakan strings.ToUpper() agar pencocokan huruf besar kecil tidak masalah (case-insensitive).
				*/
				dataSapi[j+1] = dataSapi[j] // Geser elemen yang lebih besar ke kanan untuk memberi ruang bagi key.
				j--                         // Bergerak ke elemen sebelumnya.
			}
			dataSapi[j+1] = key // Tempatkan key di posisi yang sesuai setelah pergeseran selesai.
		}
	}

	/* Program Binary Search */
	binarySearchSapiByID := func(id string) int {
		low := 0
		high := jumlahSapi - 1   // Menentukan batas bawah (low) dan atas (high) dari pencarian.
		id = strings.ToUpper(id) //  Konversi ID input ke huruf besar supaya cocok dengan data yang juga dibandingkan dalam uppercase.

		for low <= high { // Loop selama batas bawah tidak melebihi batas atas.
			mid := (low + high) / 2                    //  Tentukan indeks tengah dari area pencarian.
			midID := strings.ToUpper(dataSapi[mid].id) // Ambil ID sapi di indeks tengah dan ubah ke huruf besar.

			if midID == id { //  Jika ID cocok → kembalikan indeks mid (berhasil ditemukan).
				return mid
			} else if midID < id { //  Jika ID tengah lebih kecil dari yang dicari, cari ke kanan (low = mid + 1).
				low = mid + 1
			} else { // Jika ID tengah lebih besar, cari ke kiri (high = mid - 1).
				high = mid - 1
			}
		}
		return -1 // Jika keluar dari loop dan tidak ditemukan, kembalikan -1.
	}

	// konversi string ke int manual, return -1 jika ada karakter bukan digit
	toInt := func(s string) int {
		result := 0
		if len(s) == 0 {
			return -1
		}
		for i := 0; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return -1
			}
			result = result*10 + int(s[i]-'0')
		}
		return result
	}

	urutkanSapiByID()

	var id string
	fmt.Print("Masukkan ID sapi yang akan diedit: ")
	fmt.Scan(&id)
	id = strings.ToUpper(strings.TrimSpace(id))

	idx := binarySearchSapiByID(id)

	if idx == -1 {
		fmt.Println("Sapi tidak ditemukan.")
		return
	}

	fmt.Println("\nData saat ini:")
	fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")
	fmt.Printf("%-8s %-15s %-8d %-15s\n",
		dataSapi[idx].id, dataSapi[idx].nama, dataSapi[idx].berat, dataSapi[idx].status)

	fmt.Println("\nMasukkan data baru (ketik '-' jika tidak ingin mengubah):")

	var input string

	// Edit Nama
	fmt.Printf("Nama [%s]: ", dataSapi[idx].nama)
	fmt.Scan(&input)
	input = strings.TrimSpace(input)
	if input != "-" && input != "" {
		dataSapi[idx].nama = input
	}

	// Edit Berat dengan validasi manual
	for {
		fmt.Printf("Berat [%d]: ", dataSapi[idx].berat)
		fmt.Scan(&input)
		input = strings.TrimSpace(input)

		if input == "-" || input == "" {
			// batal ubah berat atau tidak diubah
			break
		}

		beratBaru := toInt(input)
		if beratBaru == -1 || beratBaru <= 0 {
			fmt.Println("Berat tidak valid. Masukkan angka positif atau '-' untuk batal.")
			continue
		}

		dataSapi[idx].berat = beratBaru

		if beratBaru < 100 {
			dataSapi[idx].status = "pedet"
		} else if beratBaru < 200 {
			dataSapi[idx].status = "grower"
		} else {
			dataSapi[idx].status = "siap_potong"
		}
		break
	}

	fmt.Println("Data sapi berhasil diperbarui.")
}

/* Delete Data Sapi */
func hapusSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	urutkanSapiByID := func() {
		for i := 1; i < jumlahSapi; i++ {
			key := dataSapi[i]
			j := i - 1
			for j >= 0 && strings.ToUpper(dataSapi[j].id) > strings.ToUpper(key.id) {
				dataSapi[j+1] = dataSapi[j]
				j--
			}
			dataSapi[j+1] = key
		}
	}

	binarySearchSapiByID := func(id string) int {
		low := 0
		high := jumlahSapi - 1
		id = strings.ToUpper(id)

		for low <= high {
			mid := (low + high) / 2
			midID := strings.ToUpper(dataSapi[mid].id)

			if midID == id {
				return mid
			} else if midID < id {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return -1
	}

	// Urutkan data sapi dulu
	urutkanSapiByID()

	// Input ID yang ingin dihapus
	var id string
	fmt.Print("Masukkan ID sapi yang akan dihapus: ")
	fmt.Scan(&id)
	id = strings.ToUpper(strings.TrimSpace(id))

	// Cari dengan binary search
	idx := binarySearchSapiByID(id)

	if idx == -1 {
		fmt.Println("Sapi tidak ditemukan.")
		return
	}

	// Tampilkan data yang akan dihapus
	fmt.Println("\nData sapi yang akan dihapus:")
	fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")
	fmt.Printf("%-8s %-15s %-8d %-15s\n",
		dataSapi[idx].id, dataSapi[idx].nama, dataSapi[idx].berat, dataSapi[idx].status)

	var konfirmasi string
	fmt.Print("\nApakah Anda yakin ingin menghapus data ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if strings.ToLower(konfirmasi) == "y" {
		// Hapus data kesehatan terkait sapi ini terlebih dahulu
		newIndex := 0
		for i := 0; i < jumlahKesehatan; i++ {
			if dataKesehatan[i].idSapi != id {
				dataKesehatan[newIndex] = dataKesehatan[i]
				newIndex++
			}
		}
		jumlahKesehatan = newIndex

		// Geser semua elemen setelah idx ke kiri
		for i := idx; i < jumlahSapi-1; i++ {
			dataSapi[i] = dataSapi[i+1]
		}
		jumlahSapi--
		fmt.Println("Data sapi berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

/* Urutkan Sapi By Berat*/
func urutkanSapiByBerat() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var asc int
	fmt.Print("Urutkan Ascending (1) atau Descending (0)? ")
	fmt.Scan(&asc)

	/* Selection Sort */
	for i := 0; i < jumlahSapi-1; i++ { // Loop dari indeks pertama sampai sebelum terakhir
		extremeIdx := i // Asumsikan posisi i sebagai indeks nilai terkecil/terbesar

		for j := i + 1; j < jumlahSapi; j++ { // Cari nilai terkecil/terbesar di sisa array
			if (asc == 1 && dataSapi[j].berat < dataSapi[extremeIdx].berat) || // Jika ascending dan data j lebih kecil
				(asc == 0 && dataSapi[j].berat > dataSapi[extremeIdx].berat) { // Jika descending dan data j lebih besar
				extremeIdx = j // Update indeks nilai terkecil/terbesar
			}
		}
		dataSapi[i], dataSapi[extremeIdx] = dataSapi[extremeIdx], dataSapi[i] // Tukar posisi elemen di i dan extremeIdx
	}

	fmt.Println("Data sapi telah diurutkan.")
	tampilkanSapi()
}

/*  Cari Sapi By Id */
func cariSapiByID() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var id string
	fmt.Print("Masukkan ID sapi: ")
	fmt.Scan(&id)
	id = strings.ToUpper(id)

	found := false
	fmt.Println("\nHasil pencarian:")
	fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")

	/* Program Sequental Search */
	for i := 0; i < jumlahSapi; i++ { //Loop dari awal hingga akhir array dataSapi.
		if strings.ToUpper(dataSapi[i].id) == id { // Setiap dataSapi[i].id diubah ke huruf besar, lalu dibandingkan dengan input id.
			fmt.Printf("%-8s %-15s %-8d %-15s\n",
				dataSapi[i].id, dataSapi[i].nama, dataSapi[i].berat, dataSapi[i].status)
			found = true // Jika cocok, maka data sapi akan ditampilkan, dan found di-set ke true.
		}
	}

	if !found {
		fmt.Println("Sapi tidak ditemukan.")
	}
}

/* Sequential Search */
func cariSapiByStatus() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var status string
	fmt.Print("Masukkan status (Pedet/Grower/Siap_Potong): ")
	fmt.Scan(&status)
	status = strings.ToLower(status)

	fmt.Printf("\nDaftar Sapi dengan status '%s':\n", status)
	fmt.Printf("%-8s %-15s %-8s\n", "ID", "Nama", "Berat")
	found := false
	/* Sequental Search */
	for i := 0; i < jumlahSapi; i++ {
		if strings.ToLower(dataSapi[i].status) == status {
			fmt.Printf("%-8s %-15s %-8d\n",
				dataSapi[i].id, dataSapi[i].nama, dataSapi[i].berat)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada sapi dengan status tersebut.")
	}
}

// ========== FUNGSI MANAJEMEN PAKAN ==========
func tambahPakan() {
	// Helper untuk cek apakah string hanya angka 0-9
	isDigitOnly := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
		return true
	}

	// Helper untuk konversi string angka ke int secara manual
	toInt := func(s string) int {
		result := 0
		for i := 0; i < len(s); i++ {
			result = result*10 + int(s[i]-'0')
		}
		return result
	}

	if jumlahPakan >= Maks {
		fmt.Println("Kapasitas data pakan penuh.")
		return
	}

	var pakan Pakan
	fmt.Print("Nama Pakan: ")
	fmt.Scan(&pakan.nama)

	for {
		fmt.Print("Harga: ")
		var input string
		fmt.Scan(&input)
		if isDigitOnly(input) {
			pakan.harga = toInt(input)
			break
		} else {
			fmt.Println("Harga tidak valid. Masukkan angka positif atau nol.")
		}
	}

	for {
		fmt.Print("Stok: ")
		var input string
		fmt.Scan(&input)
		if isDigitOnly(input) {
			pakan.stok = toInt(input)
			break
		} else {
			fmt.Println("Stok tidak valid. Masukkan angka positif atau nol.")
		}
	}

	dataPakan[jumlahPakan] = pakan
	jumlahPakan++
	fmt.Println("Data pakan berhasil ditambahkan.")
}

func tampilkanPakan() {
	if jumlahPakan == 0 {
		fmt.Println("Belum ada data pakan.")
		return
	}

	fmt.Println("\nDaftar Pakan:")
	fmt.Printf("%-15s %-10s %-8s\n", "Nama", "Harga", "Stok")
	for i := 0; i < jumlahPakan; i++ {
		fmt.Printf("%-15s %-10d %-8d\n",
			dataPakan[i].nama, dataPakan[i].harga, dataPakan[i].stok)
	}
}

/* Insertion Sort Dan Binary Search*/
func updateStokPakan() {
	if jumlahPakan == 0 {
		fmt.Println("Belum ada data pakan.")
		return
	}

	// Fungsi internal untuk mengurutkan dataPakan berdasarkan nama pakan (Insertion Sort)
	urutkanPakanByNama := func() {
		for i := 1; i < jumlahPakan; i++ {
			key := dataPakan[i]
			j := i - 1
			for j >= 0 && strings.ToUpper(dataPakan[j].nama) > strings.ToUpper(key.nama) {
				dataPakan[j+1] = dataPakan[j]
				j--
			}
			dataPakan[j+1] = key
		}
	}

	// Fungsi internal binary search berdasarkan nama pakan (case-insensitive)
	binarySearchPakanByNama := func(nama string) int {
		low := 0
		high := jumlahPakan - 1
		nama = strings.ToUpper(nama)

		for low <= high {
			mid := (low + high) / 2
			midNama := strings.ToUpper(dataPakan[mid].nama)

			if midNama == nama {
				return mid
			} else if midNama < nama {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return -1
	}

	// Helper untuk konversi string ke int manual, mengembalikan -1 jika ada karakter non-digit
	toInt := func(s string) int {
		result := 0
		for i := 0; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return -1 // invalid input
			}
			result = result*10 + int(s[i]-'0')
		}
		return result
	}

	// Urutkan dulu data pakan
	urutkanPakanByNama()

	// Tampilkan data pakan
	tampilkanPakan()

	// Input nama pakan yang akan diupdate
	var namaPakan string
	fmt.Print("\nMasukkan nama pakan yang akan diupdate: ")
	fmt.Scan(&namaPakan)
	namaPakan = strings.ToUpper(strings.TrimSpace(namaPakan))

	// Cari index pakan dengan binary search
	idx := binarySearchPakanByNama(namaPakan)

	if idx == -1 {
		fmt.Println("Pakan tidak ditemukan.")
		return
	}

	// Tampilkan stok saat ini
	fmt.Printf("\nData pakan %s:\n", dataPakan[idx].nama)
	fmt.Printf("Stok saat ini: %d\n", dataPakan[idx].stok)

	// Loop input stok baru sampai valid (non-negatif integer)
	for {
		fmt.Print("Masukkan stok baru: ")
		var input string
		fmt.Scan(&input)
		stokBaru := toInt(input)
		if stokBaru == -1 {
			fmt.Println("Input tidak valid. Masukkan angka bulat positif atau nol.")
			continue
		}
		if stokBaru < 0 {
			fmt.Println("Stok tidak boleh negatif.")
			continue
		}
		// Update stok
		dataPakan[idx].stok = stokBaru
		fmt.Println("Stok pakan berhasil diperbarui.")
		break
	}
}

func hapusPakan() {
	if jumlahPakan == 0 {
		fmt.Println("Belum ada data pakan.")
		return
	}

	// Fungsi internal untuk mengurutkan dataPakan berdasarkan nama pakan (Insertion Sort)
	urutkanPakanByNama := func() {
		for i := 1; i < jumlahPakan; i++ {
			key := dataPakan[i]
			j := i - 1
			for j >= 0 && strings.ToUpper(dataPakan[j].nama) > strings.ToUpper(key.nama) {
				dataPakan[j+1] = dataPakan[j]
				j--
			}
			dataPakan[j+1] = key
		}
	}

	// Fungsi internal binary search berdasarkan nama pakan (case-insensitive)
	binarySearchPakanByNama := func(nama string) int {
		low := 0
		high := jumlahPakan - 1
		nama = strings.ToUpper(nama)

		for low <= high {
			mid := (low + high) / 2
			midNama := strings.ToUpper(dataPakan[mid].nama)

			if midNama == nama {
				return mid
			} else if midNama < nama {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return -1
	}

	// Urutkan data pakan dulu
	urutkanPakanByNama()

	// Tampilkan data pakan
	tampilkanPakan()

	// Input nama pakan yang akan dihapus
	var namaPakan string
	fmt.Print("\nMasukkan nama pakan yang akan dihapus: ")
	fmt.Scan(&namaPakan)
	namaPakan = strings.ToUpper(strings.TrimSpace(namaPakan))

	// Cari index pakan dengan binary search
	idx := binarySearchPakanByNama(namaPakan)

	if idx == -1 {
		fmt.Println("Pakan tidak ditemukan.")
		return
	}

	// Tampilkan data pakan yang akan dihapus
	fmt.Printf("\nData pakan yang akan dihapus:\n")
	fmt.Printf("%-15s %-10s\n", "Nama", "Stok")
	fmt.Printf("%-15s %-10d\n", dataPakan[idx].nama, dataPakan[idx].stok)

	// Konfirmasi hapus
	var konfirmasi string
	fmt.Print("\nApakah Anda yakin ingin menghapus data ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if strings.ToLower(konfirmasi) == "y" {
		// Geser data pakan setelah idx ke kiri (hapus elemen idx)
		for i := idx; i < jumlahPakan-1; i++ {
			dataPakan[i] = dataPakan[i+1]
		}
		jumlahPakan--
		fmt.Println("Data pakan berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

/* Selection Sort */
func urutkanPakanByHarga() {
	if jumlahPakan == 0 {
		fmt.Println("Belum ada data pakan.")
		return
	}

	var asc int
	fmt.Print("Urutkan Ascending (1) atau Descending (0)? ")
	fmt.Scan(&asc)

	for i := 1; i < jumlahPakan; i++ {
		key := dataPakan[i]
		j := i - 1
		for j >= 0 && ((asc == 1 && dataPakan[j].harga > key.harga) ||
			(asc == 0 && dataPakan[j].harga < key.harga)) {
			dataPakan[j+1] = dataPakan[j]
			j--
		}
		dataPakan[j+1] = key
	}

	fmt.Println("Data pakan telah diurutkan.")
	tampilkanPakan()
}

// ========== FUNGSI MANAJEMEN KESEHATAN ==========
func tambahKesehatan() {
	if jumlahKesehatan >= Maks {
		fmt.Println("Kapasitas data kesehatan penuh.")
		return
	}

	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi. Silakan tambah data sapi terlebih dahulu.")
		return
	}

	tampilkanSapi()

	var kesehatan Kesehatan
	fmt.Print("\nMasukkan ID sapi: ")
	fmt.Scan(&kesehatan.idSapi)
	kesehatan.idSapi = strings.ToUpper(strings.TrimSpace(kesehatan.idSapi))

	// Urutkan data sapi berdasarkan ID terlebih dahulu (Insertion Sort)
	for i := 1; i < jumlahSapi; i++ {
		key := dataSapi[i]
		j := i - 1
		for j >= 0 && strings.ToUpper(dataSapi[j].id) > strings.ToUpper(key.id) {
			dataSapi[j+1] = dataSapi[j]
			j--
		}
		dataSapi[j+1] = key
	}

	// Binary search ID sapi
	idx := -1
	low := 0
	high := jumlahSapi - 1
	for low <= high {
		mid := (low + high) / 2
		midID := strings.ToUpper(dataSapi[mid].id)
		if midID == kesehatan.idSapi {
			idx = mid
			low = high + 1 // keluar loop tanpa break
		} else if midID < kesehatan.idSapi {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx == -1 {
		fmt.Println("ID sapi tidak ditemukan.")
		return
	}

	// Tanggal otomatis (format: dd/mm/yyyy)
	kesehatan.tanggal = time.Now().Format("02/01/2006")
	fmt.Println("Tanggal otomatis diisi:", kesehatan.tanggal)

	fmt.Print("Diagnosa: ")
	fmt.Scan(&kesehatan.diagnosa)

	dataKesehatan[jumlahKesehatan] = kesehatan
	jumlahKesehatan++
	fmt.Println("Catatan kesehatan berhasil ditambahkan.")
}

func tampilkanKesehatan() {
	if jumlahKesehatan == 0 {
		fmt.Println("Belum ada data catatan kesehatan.")
		return
	}

	fmt.Println("\nDaftar Catatan Kesehatan:")
	fmt.Printf("%-8s %-12s %-30s\n", "ID Sapi", "Tanggal", "Diagnosa")
	for i := 0; i < jumlahKesehatan; i++ {
		fmt.Printf("%-8s %-12s %-30s\n",
			dataKesehatan[i].idSapi, dataKesehatan[i].tanggal, dataKesehatan[i].diagnosa)
	}
}

func editKesehatan() {
	if jumlahKesehatan == 0 {
		fmt.Println("Belum ada data kesehatan.")
		return
	}

	// Insertion Sort berdasarkan idSapi
	for i := 1; i < jumlahKesehatan; i++ {
		key := dataKesehatan[i]
		j := i - 1
		for j >= 0 && strings.ToUpper(dataKesehatan[j].idSapi) > strings.ToUpper(key.idSapi) {
			dataKesehatan[j+1] = dataKesehatan[j]
			j--
		}
		dataKesehatan[j+1] = key
	}

	var id string
	fmt.Print("Masukkan ID sapi yang catatan kesehatannya akan diedit: ")
	fmt.Scan(&id)
	id = strings.ToUpper(strings.TrimSpace(id))

	// Binary Search
	idx := -1
	low := 0
	high := jumlahKesehatan - 1
	for low <= high {
		mid := (low + high) / 2
		midID := strings.ToUpper(dataKesehatan[mid].idSapi)
		if midID == id {
			idx = mid
			low = high + 1
		} else if midID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx == -1 {
		fmt.Println("Catatan kesehatan tidak ditemukan.")
		return
	}

	fmt.Println("Data saat ini:")
	fmt.Printf("ID Sapi: %s\n", dataKesehatan[idx].idSapi)
	fmt.Printf("Tanggal: %s\n", dataKesehatan[idx].tanggal)
	fmt.Printf("Diagnosa: %s\n", dataKesehatan[idx].diagnosa)

	var input string

	fmt.Print("Tanggal baru (ketik '-' jika tidak diubah): ")
	fmt.Scan(&input)
	input = strings.TrimSpace(input)
	if input != "-" && input != "" {
		dataKesehatan[idx].tanggal = input
	}

	fmt.Print("Diagnosa baru (ketik '-' jika tidak diubah): ")
	fmt.Scan(&input)
	input = strings.TrimSpace(input)
	if input != "-" && input != "" {
		dataKesehatan[idx].diagnosa = input
	}

	fmt.Println("Data kesehatan berhasil diperbarui.")
}

func hapusKesehatan() {
	if jumlahKesehatan == 0 {
		fmt.Println("Belum ada data kesehatan.")
		return
	}

	// Insertion Sort berdasarkan idSapi
	for i := 1; i < jumlahKesehatan; i++ {
		key := dataKesehatan[i]
		j := i - 1
		for j >= 0 && strings.ToUpper(dataKesehatan[j].idSapi) > strings.ToUpper(key.idSapi) {
			dataKesehatan[j+1] = dataKesehatan[j]
			j--
		}
		dataKesehatan[j+1] = key
	}

	var id string
	fmt.Print("Masukkan ID sapi yang catatan kesehatannya akan dihapus: ")
	fmt.Scan(&id)
	id = strings.ToUpper(strings.TrimSpace(id))

	// Binary Search
	idx := -1
	low := 0
	high := jumlahKesehatan - 1
	for low <= high {
		mid := (low + high) / 2
		midID := strings.ToUpper(dataKesehatan[mid].idSapi)
		if midID == id {
			idx = mid
			low = high + 1
		} else if midID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx == -1 {
		fmt.Println("Catatan kesehatan tidak ditemukan.")
		return
	}

	// Geser data ke kiri
	for i := idx; i < jumlahKesehatan-1; i++ {
		dataKesehatan[i] = dataKesehatan[i+1]
	}
	jumlahKesehatan--

	fmt.Println("Data kesehatan berhasil dihapus.")
}

func lihatRiwayatKesehatanSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	tampilkanSapi()

	var idSapi string
	fmt.Print("\nMasukkan ID sapi: ")
	fmt.Scan(&idSapi)
	idSapi = strings.ToUpper(strings.TrimSpace(idSapi))

	// Cek apakah ID sapi ada
	sapiAda := false
	for i := 0; i < jumlahSapi; i++ {
		if strings.EqualFold(dataSapi[i].id, idSapi) {
			sapiAda = true
		}
	}

	if !sapiAda {
		fmt.Println("ID sapi tidak ditemukan.")
		return
	}

	fmt.Printf("\nRiwayat kesehatan sapi %s:\n", idSapi)
	fmt.Printf("%-12s %-30s\n", "Tanggal", "Diagnosa")
	found := false
	for i := 0; i < jumlahKesehatan; i++ {
		if strings.EqualFold(dataKesehatan[i].idSapi, idSapi) {
			fmt.Printf("%-12s %-30s\n", dataKesehatan[i].tanggal, dataKesehatan[i].diagnosa)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada riwayat kesehatan untuk sapi ini.")
	}
}

func buatDataDummy() {
	// Data sapi dummy
	dataSapi[0] = Sapi{"S001", "Limousin", 120, "grower"}
	dataSapi[1] = Sapi{"S002", "Putih", 80, "pedet"}
	dataSapi[2] = Sapi{"S003", "Hitam", 210, "Siap_Potong"}
	jumlahSapi = 3

	// Data kesehatan dummy
	dataKesehatan[0] = Kesehatan{"S001", "01/05/2025", "Sehat"}
	dataKesehatan[1] = Kesehatan{"S001", "15/05/2025", "Demam ringan"}
	dataKesehatan[2] = Kesehatan{"S003", "10/05/2025", "Cacingan"}
	jumlahKesehatan = 3

	// Data pakan dummy
	dataPakan[0] = Pakan{"Rumput", 20000, 50}
	dataPakan[1] = Pakan{"Jagung", 10000, 30}
	dataPakan[2] = Pakan{"Dedak", 30000, 20}
	jumlahPakan = 3
}

/* Menu Utama */
func main() {
	// buatDataDummy()
	menuUtama()
}
