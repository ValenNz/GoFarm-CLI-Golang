```markdown
# Sistem Manajemen Peternakan Sapi Potong (Go CLI)

![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Aplikasi CLI untuk manajemen peternakan sapi potong berbasis Go dengan fitur lengkap untuk mengelola data sapi, pakan, dan kesehatan ternak.

## 🚀 Fitur Utama

### 📋 Modul Manajemen Sapi
- CRUD data sapi (ID, nama, berat, status)
- Sorting berdasarkan berat (Selection Sort)
- Pencarian by ID/status
- Validasi input terintegrasi

### 🌾 Modul Manajemen Pakan
- Catat stok pakan (nama, harga, jumlah)
- Urutkan by harga (Insertion Sort)
- Update stok dengan validasi

### 💊 Modul Kesehatan
- Riwayat kesehatan per sapi
- Auto-cleanup data terkait
- Tampilan kronologis

## 🛠️ Teknologi
- **Bahasa**: Go 1.18+
- **Struktur Data**: Array statis
- **Algoritma**:
  - Selection Sort (data sapi)
  - Insertion Sort (data pakan)
  - Sequential Search

## 📥 Instalasi
1. Pastikan Go terinstall:
   ```bash
   go version
   ```
2. Clone repo:
   ```bash
   git clone https://github.com/username/peternakan-sapi.git
   ```
3. Jalankan program:
   ```bash
   cd peternakan-sapi
   go run main.go
   ```

## 🖥️ Penggunaan
```
Menu Utama:
1. Manajemen Sapi
2. Manajemen Pakan  
3. Manajemen Kesehatan
0. Keluar
```

## 📂 Struktur File
```
├── main.go          # Main program
└── README.md        # Dokumentasi
```

## 🧪 Testing
Program sudah dilengkapi data dummy:
- 3 data sapi
- 3 data pakan  
- 2 riwayat kesehatan

## 🤝 Kontribusi
Pull request dipersilakan. Untuk perubahan besar, buka issue terlebih dahulu.

## 📜 Lisensi
MIT © 2025 ValenNz
```