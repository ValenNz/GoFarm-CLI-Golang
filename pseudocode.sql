CONSTANT Maks = 100

STRUCT Sapi:
    id: string
    nama: string
    berat: int
    status: string

STRUCT Pakan:
    nama: string
    harga: int
    stok: int

STRUCT Kesehatan:
    idSapi: string
    tanggal: string
    diagnosa: string

VAR dataSapi: array[Maks] of Sapi
VAR dataPakan: array[Maks] of Pakan
VAR dataKesehatan: array[Maks] of Kesehatan
VAR jumlahSapi, jumlahPakan, jumlahKesehatan, idTerakhir: int = 0

PROCEDURE menuUtama()
    LOOP
        TAMPILKAN menu utama
        INPUT pilihan
        IF pilihan == 1 THEN
            CALL menuManajemenSapi()
        ELSE IF pilihan == 2 THEN
            CALL menuManajemenPakan()
        ELSE IF pilihan == 3 THEN
            CALL menuManajemenKesehatan()
        ELSE IF pilihan == 0 THEN
            KELUAR dari program
        ELSE
            TAMPILKAN "Pilihan tidak valid"
        ENDIF
    ENDLOOP
ENDPROCEDURE

PROCEDURE menuManajemenSapi()
    LOOP
        TAMPILKAN submenu sapi
        INPUT pilihan
        SWITCH pilihan
            CASE 1: CALL tambahSapi()
            CASE 2: CALL tampilkanSapi()
            CASE 3: CALL urutkanSapiByBerat()
            CASE 4: CALL cariSapiByID()
            CASE 5: CALL cariSapiByStatus()
            CASE 6: CALL editSapi()
            CASE 7: CALL hapusSapi()
            CASE 0: RETURN
            DEFAULT: TAMPILKAN "Pilihan tidak valid"
        ENDSWITCH
    ENDLOOP
ENDPROCEDURE


PROCEDURE tambahSapi()
    IF jumlahSapi >= Maks THEN
        TAMPILKAN "Kapasitas penuh"
        RETURN
    ENDIF

    INCREMENT idTerakhir
    GENERATE id baru: "S" + 3 digit idTerakhir
    INPUT nama
    LOOP
        INPUT berat
        IF berat > 0 THEN
            BREAK
        ELSE
            TAMPILKAN "Berat tidak valid"
        ENDIF
    ENDLOOP

    TENTUKAN status berdasarkan berat:
        IF berat < 100 THEN "pedet"
        ELSE IF berat < 250 THEN "grower"
        ELSE "siap_potong"

    SIMPAN data ke array dataSapi
    INCREMENT jumlahSapi
    TAMPILKAN "Data berhasil ditambahkan"
ENDPROCEDURE

PROCEDURE tampilkanSapi()
    IF jumlahSapi == 0 THEN
        TAMPILKAN "Belum ada data"
        RETURN
    ENDIF

    TAMPILKAN seluruh isi array dataSapi
ENDPROCEDURE

PROCEDURE editSapi()
    IF jumlahSapi == 0 THEN
        TAMPILKAN "Belum ada data"
        RETURN
    ENDIF

    CALL urutkanSapiByID()   // insertion sort by ID

    INPUT id
    index ← binarySearchSapiByID(id)
    IF index == -1 THEN
        TAMPILKAN "Sapi tidak ditemukan"
        RETURN
    ENDIF

    TAMPILKAN data lama
    INPUT nama baru atau "-" jika tidak ingin mengubah
    IF input != "-" THEN update nama

    LOOP UNTIL berat valid atau "-" :
        INPUT berat baru atau "-"
        IF input == "-" THEN BREAK
        VALIDASI angka
        UPDATE berat dan status jika valid
    ENDLOOP
ENDPROCEDURE

PROCEDURE urutkanSapiByID()
    // Insertion sort berdasarkan ID (case-insensitive)
    FOR i FROM 1 TO jumlahSapi-1 DO
        key ← dataSapi[i]
        j ← i - 1
        WHILE j >= 0 AND id[j] > key.id DO
            geser dataSapi[j] ke dataSapi[j+1]
            j ← j - 1
        ENDWHILE
        tempatkan key di dataSapi[j+1]
    ENDFOR
ENDPROCEDURE

FUNCTION binarySearchSapiByID(id): int
    SET low = 0, high = jumlahSapi - 1
    WHILE low <= high DO
        SET mid = (low + high) / 2
        IF id[mid] == id THEN RETURN mid
        ELSE IF id[mid] < id THEN low = mid + 1
        ELSE high = mid - 1
    ENDWHILE
    RETURN -1
ENDFUNCTION

// ===== Submenu lain dapat dibuat dengan struktur serupa =====

PROCEDURE menuManajemenPakan()
    // Sama seperti menuManajemenSapi dengan fungsi:
    // tambahPakan, tampilkanPakan, urutkanPakanByHarga, updateStokPakan, hapusPakan
ENDPROCEDURE

PROCEDURE menuManajemenKesehatan()
    // Sama seperti lainnya, memanggil:
    // tambahKesehatan, tampilkanKesehatan, editKesehatan, hapusKesehatan, lihatRiwayatKesehatanSapi
ENDPROCEDURE

// Tambahkan prosedur tambahan sesuai modul (pakan dan kesehatan)

PROCEDURE MAIN()
    CALL menuUtama()
ENDPROCEDURE
