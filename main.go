package main

import (
	"fmt"
	"strings"
)

const Maks = 100

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

var dataSapi [Maks]Sapi
var jumlahSapi int

var dataPakan [Maks]Pakan
var jumlahPakan int

var dataKesehatan [Maks]Kesehatan
var jumlahKesehatan int

// ========== MENU UTAMA DAN SUBMENU ==========
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
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahKesehatan()
		} else if pilihan == 2 {
			tampilkanKesehatan()
		} else if pilihan == 3 {
			lihatRiwayatKesehatanSapi()
		} else if pilihan == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ========== FUNGSI MANAJEMEN SAPI ==========
func tambahSapi() {
	if jumlahSapi >= Maks {
		fmt.Println("Kapasitas data sapi penuh.")
		return
	}

	var sapi Sapi
	fmt.Print("ID: ")
	fmt.Scan(&sapi.id)

	// Cek apakah ID sudah ada
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == sapi.id {
			fmt.Println("ID sapi sudah ada.")
			return
		}
	}

	fmt.Print("Nama: ")
	fmt.Scan(&sapi.nama)
	fmt.Print("Berat (kg): ")
	fmt.Scan(&sapi.berat)

	// Validasi status
	for {
		fmt.Print("Status (Pedet/Grower/Siap Potong): ")
		fmt.Scan(&sapi.status)
		sapi.status = strings.ToLower(sapi.status)
		if sapi.status == "pedet" || sapi.status == "grower" || sapi.status == "siap potong" {
			break
		}
		fmt.Println("Status tidak valid. Harus Pedet, Grower, atau Siap Potong.")
	}

	dataSapi[jumlahSapi] = sapi
	jumlahSapi++
	fmt.Println("Data sapi berhasil ditambahkan.")
}

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

func urutkanSapiByBerat() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var asc int
	fmt.Print("Urutkan Ascending (1) atau Descending (0)? ")
	fmt.Scan(&asc)

	for i := 0; i < jumlahSapi-1; i++ {
		extremeIdx := i
		for j := i + 1; j < jumlahSapi; j++ {
			if (asc == 1 && dataSapi[j].berat < dataSapi[extremeIdx].berat) ||
				(asc == 0 && dataSapi[j].berat > dataSapi[extremeIdx].berat) {
				extremeIdx = j
			}
		}
		dataSapi[i], dataSapi[extremeIdx] = dataSapi[extremeIdx], dataSapi[i]
	}

	fmt.Println("Data sapi telah diurutkan.")
	tampilkanSapi()
}

func cariSapiByID() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var id string
	fmt.Print("Masukkan ID sapi: ")
	fmt.Scan(&id)

	found := false
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == id {
			fmt.Println("\nData ditemukan:")
			fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")
			fmt.Printf("%-8s %-15s %-8d %-15s\n",
				dataSapi[i].id, dataSapi[i].nama, dataSapi[i].berat, dataSapi[i].status)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Sapi tidak ditemukan.")
	}
}

func cariSapiByStatus() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var status string
	fmt.Print("Masukkan status (Pedet/Grower/Siap Potong): ")
	fmt.Scan(&status)
	status = strings.ToLower(status)

	fmt.Printf("\nDaftar Sapi dengan status '%s':\n", status)
	fmt.Printf("%-8s %-15s %-8s\n", "ID", "Nama", "Berat")
	found := false
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

func editSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var id string
	fmt.Print("Masukkan ID sapi yang akan diedit: ")
	fmt.Scan(&id)

	idx := -1
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Sapi tidak ditemukan.")
		return
	}

	fmt.Println("\nData saat ini:")
	fmt.Printf("%-8s %-15s %-8s %-15s\n", "ID", "Nama", "Berat", "Status")
	fmt.Printf("%-8s %-15s %-8d %-15s\n",
		dataSapi[idx].id, dataSapi[idx].nama, dataSapi[idx].berat, dataSapi[idx].status)

	fmt.Println("\nMasukkan data baru (kosongkan jika tidak ingin mengubah):")

	var input string

	fmt.Printf("Nama [%s]: ", dataSapi[idx].nama)
	fmt.Scan(&input)
	if input != "" {
		dataSapi[idx].nama = input
	}

	fmt.Printf("Berat [%d]: ", dataSapi[idx].berat)
	fmt.Scan(&input)
	if input != "" {
		var newBerat int
		fmt.Sscanf(input, "%d", &newBerat)
		dataSapi[idx].berat = newBerat
	}

	for {
		fmt.Printf("Status [%s]: ", dataSapi[idx].status)
		fmt.Scan(&input)
		if input == "" {
			break
		}
		input = strings.ToLower(input)
		if input == "pedet" || input == "grower" || input == "siap potong" {
			dataSapi[idx].status = input
			break
		}
		fmt.Println("Status tidak valid. Harus Pedet, Grower, atau Siap Potong.")
	}

	fmt.Println("Data sapi berhasil diperbarui.")
}

func hapusSapi() {
	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi.")
		return
	}

	var id string
	fmt.Print("Masukkan ID sapi yang akan dihapus: ")
	fmt.Scan(&id)

	idx := -1
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == id {
			idx = i
			break
		}
	}

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

// ========== FUNGSI MANAJEMEN PAKAN ==========
func tambahPakan() {
	if jumlahPakan >= Maks {
		fmt.Println("Kapasitas data pakan penuh.")
		return
	}

	var pakan Pakan
	fmt.Print("Nama Pakan: ")
	fmt.Scan(&pakan.nama)
	fmt.Print("Harga: ")
	fmt.Scan(&pakan.harga)
	fmt.Print("Stok: ")
	fmt.Scan(&pakan.stok)

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

func updateStokPakan() {
	if jumlahPakan == 0 {
		fmt.Println("Belum ada data pakan.")
		return
	}

	tampilkanPakan()

	var namaPakan string
	fmt.Print("\nMasukkan nama pakan yang akan diupdate: ")
	fmt.Scan(&namaPakan)

	idx := -1
	for i := 0; i < jumlahPakan; i++ {
		if strings.EqualFold(dataPakan[i].nama, namaPakan) {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Pakan tidak ditemukan.")
		return
	}

	fmt.Printf("\nData pakan %s:\n", dataPakan[idx].nama)
	fmt.Printf("Stok saat ini: %d\n", dataPakan[idx].stok)

	var stokBaru int
	fmt.Print("Masukkan stok baru: ")
	fmt.Scan(&stokBaru)

	if stokBaru < 0 {
		fmt.Println("Stok tidak boleh negatif.")
		return
	}

	dataPakan[idx].stok = stokBaru
	fmt.Println("Stok pakan berhasil diperbarui.")
}

// ========== FUNGSI MANAJEMEN KESEHATAN ==========
func tambahKesehatan() {
	if jumlahKesehatan >= Maks {
		fmt.Println("Kapasitas data kesehatan penuh.")
		return
	}

	if jumlahSapi == 0 {
		fmt.Println("Belum ada data sapi. Silakan tambah sapi terlebih dahulu.")
		return
	}

	tampilkanSapi()

	var kesehatan Kesehatan
	fmt.Print("\nID Sapi: ")
	fmt.Scan(&kesehatan.idSapi)

	// Cek apakah ID sapi ada
	sapiAda := false
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == kesehatan.idSapi {
			sapiAda = true
			break
		}
	}

	if !sapiAda {
		fmt.Println("ID sapi tidak ditemukan.")
		return
	}

	fmt.Print("Tanggal (dd/mm/yyyy): ")
	fmt.Scan(&kesehatan.tanggal)
	fmt.Print("Diagnosa: ")
	fmt.Scan(&kesehatan.diagnosa)

	dataKesehatan[jumlahKesehatan] = kesehatan
	jumlahKesehatan++
	fmt.Println("Catatan kesehatan berhasil ditambahkan.")
}

func tampilkanKesehatan() {
	if jumlahKesehatan == 0 {
		fmt.Println("Belum ada data kesehatan.")
		return
	}

	fmt.Println("\nDaftar Catatan Kesehatan:")
	fmt.Printf("%-8s %-12s %-30s\n", "ID Sapi", "Tanggal", "Diagnosa")
	for i := 0; i < jumlahKesehatan; i++ {
		fmt.Printf("%-8s %-12s %-30s\n",
			dataKesehatan[i].idSapi, dataKesehatan[i].tanggal, dataKesehatan[i].diagnosa)
	}
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

	// Cek apakah ID sapi ada
	sapiAda := false
	for i := 0; i < jumlahSapi; i++ {
		if dataSapi[i].id == idSapi {
			sapiAda = true
			break
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
		if dataKesehatan[i].idSapi == idSapi {
			fmt.Printf("%-12s %-30s\n",
				dataKesehatan[i].tanggal, dataKesehatan[i].diagnosa)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada riwayat kesehatan untuk sapi ini.")
	}
}

// ========== MAIN ==========
func main() {
	// Data dummy untuk testing
	dataSapi[0] = Sapi{"S001", "Budi", 250, "Siap Potong"}
	dataSapi[1] = Sapi{"S002", "Lani", 150, "Grower"}
	dataSapi[2] = Sapi{"S003", "Momo", 80, "Pedet"}
	jumlahSapi = 3

	dataPakan[0] = Pakan{"Rumput Gajah", 5000, 100}
	dataPakan[1] = Pakan{"Konsentrat", 12000, 50}
	dataPakan[2] = Pakan{"Jerami", 3000, 200}
	jumlahPakan = 3

	dataKesehatan[0] = Kesehatan{"S001", "01/01/2023", "Sehat"}
	dataKesehatan[1] = Kesehatan{"S002", "15/01/2023", "Flu ringan"}
	jumlahKesehatan = 2

	menuUtama()
}
