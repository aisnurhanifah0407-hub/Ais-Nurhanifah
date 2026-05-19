package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Kapasitas maksimum array statis sesuai spesifikasi
const MAX int = 100

// ================= TIPE BENTUKAN (STRUKTUR DATA) =================

type MoodRecord struct {
	Tanggal   string // Format: YYYY-MM-DD
	SkorEmosi int    // Skala 1-10
	Deskripsi string //
}

type Task struct {
	Nama      string //
	Durasi    int    // Dalam menit
	Prioritas int    // 1 (Tinggi), 2 (Sedang), 3 (Rendah)
	Selesai   bool
	Tanggal   string // Format: YYYY-MM-DD
}

// SPESIFIKASI: Variabel global HANYA boleh untuk array utama dan ukurannya (N)
var listMood [MAX]MoodRecord
var nMood int = 0

var listTask [MAX]Task
var nTask int = 0

// ================= SUBPROGRAM PEMBANTU =================

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi bantu untuk mengecek apakah sebuah tanggal berada dalam 7 hari terakhir dari hari ini
func apakahMingguIni(tanggalTarget string) bool {
	tglFormat := "2006-01-02"
	t, err := time.Parse(tglFormat, tanggalTarget)
	if err != nil {
		return false
	}

	hariIni := time.Now()
	tujuhHariLalu := hariIni.AddDate(0, 0, -7)

	// Cek apakah 't' berada di antara 7 hari lalu hingga hari ini
	return t.After(tujuhHariLalu) && t.Before(hariIni.AddDate(0, 0, 1))
}

// ================= MAIN PROGRAM =================

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======================================================================")
	fmt.Println("   Halo! Selamat datang di CekKesehatanMental - Asisten Virtual Anda ")
	fmt.Println("             dan Produktivitas Harian Anda               ")
	fmt.Println("======================================================================")

	// Meminta input Nama Pengguna di awal program
	namaUser := readInput(reader, "Sebelum kita mulai, boleh tahu nama Anda? ")
	if namaUser == "" {
		namaUser = "Pengguna" // Nilai default jika nama dikosongkan
	}

	// Interaksi Awal Asisten Virtual (Mengirimkan variabel namaUser)
	asistenCheckIn(reader, namaUser)

	jalankanMenu := true
	for jalankanMenu {
		fmt.Printf("\n>>> [MENU UTAMA CEKKESEHATANMENTAL - Halo, %s!] <<<\n", namaUser)
		fmt.Println("1. Kelola Catatan Suasana Hati (Mood)")
		fmt.Println("2. Kelola Daftar Tugas Harian (Task)")
		fmt.Println("3. Cari Data (Tugas / Catatan Emosi)")
		fmt.Println("4. Urutkan Daftar Tugas")
		fmt.Println("5. Lihat Statistik & Tren Mingguan")
		fmt.Println("6. Istirahat / Keluar Aplikasi")
		pilihan := readInput(reader, "Pilih menu (1-6): ")

		switch pilihan {
		case "1":
			menuMood(reader)
		case "2":
			menuTask(reader)
		case "3":
			menuCari(reader)
		case "4":
			menuUrut(reader)
		case "5":
			tampilkanStatistik()
		case "6":
			fmt.Printf("\n[CekKesehatanMental]: Terima kasih sudah berproses hari ini, %s. Jangan lupa menjaga kesehatan mentalmu!\n", namaUser)
			jalankanMenu = false
		default:
			fmt.Println("\n[CekKesehatanMental]: Pilihan tidak tersedia, coba masukkan angka 1-6 ya.")
		}
	}
}

// ================= INTERAKSI ASISTEN VIRTUAL =================

func asistenCheckIn(reader *bufio.Reader, namaUser string) {
	tglHariIni := time.Now().Format("2006-01-02")
	fmt.Printf("\n[CekKesehatanMental]: Halo %s! Hari ini tanggal %s. Bagaimana keadaan hatimu sekarang?\n", namaUser, tglHariIni)
	skorStr := readInput(reader, "Berikan skor emosimu hari ini (Skala 1-10): ")
	skor, _ := strconv.Atoi(skorStr)
	deskripsi := readInput(reader, "Ceritakan sedikit apa yang kamu rasakan hari ini: ")

	if nMood < MAX {
		listMood[nMood] = MoodRecord{Tanggal: tglHariIni, SkorEmosi: skor, Deskripsi: deskripsi}
		nMood++
	}

	fmt.Printf("\n[CekKesehatanMental]: Terima kasih sudah berbagi cerita, %s.\n", namaUser)
	if skor <= 4 {
		fmt.Println(">> Respon Asisten: Tampaknya hari ini cukup berat buatmu. Jangan terlalu memaksakan diri ya. <<")
	} else if skor <= 7 {
		fmt.Println(">> Respon Asisten: Hari yang lumayan stabil. Tetap jaga ritme kerjamu! <<")
	} else {
		fmt.Println(">> Respon Asisten: Luar biasa! Energinya sangat positif hari ini. <<")
	}
}

// ================= KELOLA DATA MOOD (CRUD) =================

func menuMood(reader *bufio.Reader) {
	fmt.Println("\n--- KELOLA RIWAYAT SUASANA HATI ---")
	fmt.Println("1. Tambah Catatan Baru")
	fmt.Println("2. Ubah Cerita/Skor Mood")
	fmt.Println("3. Hapus Catatan Mood")
	pilihan := readInput(reader, "Pilih aksi: ")

	switch pilihan {
	case "1":
		if nMood >= MAX {
			fmt.Println("[CekKesehatanMental]: Penyimpanan penuh!")
			return
		}
		tgl := readInput(reader, "Masukkan Tanggal (YYYY-MM-DD): ")
		skorStr := readInput(reader, "Skor Emosi (1-10): ")
		skor, _ := strconv.Atoi(skorStr)
		deskripsi := readInput(reader, "Deskripsi Perasaan: ")

		listMood[nMood] = MoodRecord{Tanggal: tgl, SkorEmosi: skor, Deskripsi: deskripsi}
		nMood++
		fmt.Println("[CekKesehatanMental]: Catatan emosi berhasil disimpan.")
	case "2":
		tgl := readInput(reader, "Masukkan Tanggal yang ingin diubah (YYYY-MM-DD): ")
		idx := -1
		// SPESIFIKASI: Sequential search tanpa break/continue
		for i := 0; i < nMood && idx == -1; i++ {
			if listMood[i].Tanggal == tgl {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("[CekKesehatanMental]: Maaf, data pada tanggal tersebut tidak ditemukan.")
			return
		}
		skorStr := readInput(reader, "Skor Emosi Baru (1-10): ")
		skor, _ := strconv.Atoi(skorStr)
		deskripsi := readInput(reader, "Deskripsi Perasaan Baru: ")
		listMood[idx].SkorEmosi = skor
		listMood[idx].Deskripsi = deskripsi
		fmt.Println("[CekKesehatanMental]: Catatan emosimu berhasil diperbarui.")
	case "3":
		tgl := readInput(reader, "Masukkan Tanggal yang ingin dihapus (YYYY-MM-DD): ")
		idx := -1
		// SPESIFIKASI: Sequential search tanpa break/continue
		for i := 0; i < nMood && idx == -1; i++ {
			if listMood[i].Tanggal == tgl {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("[CekKesehatanMental]: Data tidak ditemukan.")
			return
		}
		// Menggeser elemen array statis ke kiri untuk menghapus
		for i := idx; i < nMood-1; i++ {
			listMood[i] = listMood[i+1]
		}
		nMood--
		fmt.Println("[CekKesehatanMental]: Catatan emosi berhasil dihapus.")
	}
}

// ================= KELOLA DATA TUGAS (CRUD) =================

func menuTask(reader *bufio.Reader) {
	fmt.Println("\n--- DAFTAR TUGAS DAN PRODUKTIVITAS ---")
	fmt.Println("1. Tambah Tugas Baru")
	fmt.Println("2. Perbarui Status Tugas (Selesai/Belum)")
	fmt.Println("3. Hapus Tugas dari Daftar")
	pilihan := readInput(reader, "Pilih aksi: ")

	switch pilihan {
	case "1":
		if nTask >= MAX {
			fmt.Println("[CekKesehatanMental]: Daftar tugas penuh!")
			return
		}
		nama := readInput(reader, "Nama Tugas / Aktivitas: ")
		durasiStr := readInput(reader, "Estimasi Durasi Pengerjaan (menit): ")
		durasi, _ := strconv.Atoi(durasiStr)
		prioStr := readInput(reader, "Prioritas Tugas (1=Tinggi, 2=Sedang, 3=Rendah): ")
		prio, _ := strconv.Atoi(prioStr)
		tgl := readInput(reader, "Target Tanggal Pengerjaan (YYYY-MM-DD): ")

		listTask[nTask] = Task{Nama: nama, Durasi: durasi, Prioritas: prio, Tanggal: tgl, Selesai: false}
		nTask++
		fmt.Println("[CekKesehatanMental]: Tugas baru berhasil ditambahkan.")
	case "2":
		nama := readInput(reader, "Masukkan Nama Tugas yang ingin diubah statusnya: ")
		idx := -1
		// SPESIFIKASI: Sequential search tanpa break/continue
		for i := 0; i < nTask && idx == -1; i++ {
			if strings.EqualFold(listTask[i].Nama, nama) {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("[CekKesehatanMental]: Tugas tidak ditemukan.")
			return
		}
		statusStr := readInput(reader, "Apakah tugas ini sudah kamu selesaikan? (y/n): ")
		if strings.ToLower(statusStr) == "y" {
			listTask[idx].Selesai = true
			fmt.Println("[CekKesehatanMental]: Status tugas berhasil diperbarui menjadi Selesai.")
		} else {
			listTask[idx].Selesai = false
			fmt.Println("[CekKesehatanMental]: Status tugas dikembalikan ke belum selesai.")
		}
	case "3":
		nama := readInput(reader, "Masukkan Nama Tugas yang ingin dihapus: ")
		idx := -1
		// SPESIFIKASI: Sequential search tanpa break/continue
		for i := 0; i < nTask && idx == -1; i++ {
			if strings.EqualFold(listTask[i].Nama, nama) {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("[CekKesehatanMental]: Tugas tidak ditemukan.")
			return
		}
		// Menggeser elemen array statis ke kiri untuk menghapus
		for i := idx; i < nTask-1; i++ {
			listTask[i] = listTask[i+1]
		}
		nTask--
		fmt.Println("[CekKesehatanMental]: Tugas berhasil dihapus dari daftar.")
	}
}

// ================= PENCARIAN (SEARCHING) =================

func menuCari(reader *bufio.Reader) {
	fmt.Println("\n--- FITUR PENCARIAN CEKKESEHATANMENTAL ---")
	fmt.Println("1. Cari Tugas berdasarkan Kata Kunci Nama (Sequential Search)")
	fmt.Println("2. Cari Catatan Mood berdasarkan Tanggal (Binary Search)")
	pilihan := readInput(reader, "Pilih opsi pencarian: ")

	if pilihan == "1" {
		// SPESIFIKASI: Sequential Search berdasarkan Kata Kunci
		keyword := readInput(reader, "Masukkan kata kunci nama tugas: ")
		fmt.Println("\n[CekKesehatanMental]: Berikut hasil pencarian tugasmu:")
		ketemu := false

		for i := 0; i < nTask; i++ {
			if strings.Contains(strings.ToLower(listTask[i].Nama), strings.ToLower(keyword)) {
				fmt.Printf("- %s | Durasi: %d mnt | Prioritas: %d | Tanggal: %s\n",
					listTask[i].Nama, listTask[i].Durasi, listTask[i].Prioritas, listTask[i].Tanggal)
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("[CekKesehatanMental]: Tidak ada nama tugas yang cocok.")
		}
	} else if pilihan == "2" {
		// SPESIFIKASI: Binary Search berdasarkan Tanggal
		if nMood == 0 {
			fmt.Println("[CekKesehatanMental]: Riwayat suasana hatimu masih kosong.")
			return
		}
		tglTarget := readInput(reader, "Masukkan tanggal catatan mood (YYYY-MM-DD): ")

		// Urutkan dulu datanya berdasarkan tanggal agar bisa diproses Binary Search
		for i := 0; i < nMood-1; i++ {
			minIdx := i
			for j := i + 1; j < nMood; j++ {
				if listMood[j].Tanggal < listMood[minIdx].Tanggal {
					minIdx = j
				}
			}
			listMood[i], listMood[minIdx] = listMood[minIdx], listMood[i]
		}

		// Implementasi Binary Search (Tanpa kata kunci break)
		low, high := 0, nMood-1
		idx := -1

		for low <= high && idx == -1 {
			mid := (low + high) / 2
			if listMood[mid].Tanggal == tglTarget {
				idx = mid
			} else if listMood[mid].Tanggal < tglTarget {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if idx != -1 {
			fmt.Printf("\n[CekKesehatanMental]: Data Ditemukan!\nTanggal: %s\nSkor Emosi: %d/10\nCatatan: \"%s\"\n",
				listMood[idx].Tanggal, listMood[idx].SkorEmosi, listMood[idx].Deskripsi)
		} else {
			fmt.Println("[CekKesehatanMental]: Catatan mood untuk tanggal tersebut tidak ditemukan.")
		}
	}
}

// ================= PENGURUTAN (SORTING) =================

func menuUrut(reader *bufio.Reader) {
	if nTask == 0 {
		fmt.Println("[CekKesehatanMental]: Daftar tugas kosong, tidak ada data untuk diurutkan.")
		return
	}

	fmt.Println("\n--- PENGURUTAN DAFTAR TUGAS ---")
	fmt.Println("1. Urutkan berdasarkan Tingkat Prioritas (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Waktu Pengerjaan/Durasi (Insertion Sort)")
	pilihan := readInput(reader, "Pilih metode pengurutan (1-2): ")

	// SPESIFIKASI: Harus bisa urutan naik (ascending) maupun turun (descending)
	fmt.Println("Pilih Arah Urutan:")
	fmt.Println("1. Urutan Naik (Ascending)")
	fmt.Println("2. Urutan Turun (Descending)")
	arah := readInput(reader, "Pilih arah (1-2): ")

	if pilihan == "1" {
		// Selection Sort (Prioritas)
		for i := 0; i < nTask-1; i++ {
			targetIdx := i
			for j := i + 1; j < nTask; j++ {
				if arah == "1" { // Ascending
					if listTask[j].Prioritas < listTask[targetIdx].Prioritas {
						targetIdx = j
					}
				} else { // Descending
					if listTask[j].Prioritas > listTask[targetIdx].Prioritas {
						targetIdx = j
					}
				}
			}
			listTask[i], listTask[targetIdx] = listTask[targetIdx], listTask[i]
		}
		fmt.Println("[CekKesehatanMental]: Berhasil mengurutkan tugas berdasarkan prioritas!")
		tampilkanSemuaTugas()

	} else if pilihan == "2" {
		// Insertion Sort (Durasi)
		for i := 1; i < nTask; i++ {
			key := listTask[i]
			j := i - 1

			if arah == "1" { // Ascending
				for j >= 0 && listTask[j].Durasi > key.Durasi {
					listTask[j+1] = listTask[j]
					j = j - 1
				}
			} else { // Descending
				for j >= 0 && listTask[j].Durasi < key.Durasi {
					listTask[j+1] = listTask[j]
					j = j - 1
				}
			}
			listTask[j+1] = key
		}
		fmt.Println("[CekKesehatanMental]: Berhasil mengurutkan tugas berdasarkan durasi!")
		tampilkanSemuaTugas()
	}
}

func tampilkanSemuaTugas() {
	for i := 0; i < nTask; i++ {
		t := listTask[i]
		prioText := "Rendah"
		if t.Prioritas == 1 {
			prioText = "Tinggi"
		} else if t.Prioritas == 2 {
			prioText = "Sedang"
		}
		fmt.Printf("- %s | Prioritas: %s | Durasi: %d mnt | Selesai: %t (%s)\n", t.Nama, prioText, t.Durasi, t.Selesai, t.Tanggal)
	}
}

// ================= STATISTIK TREN =================

func tampilkanStatistik() {
	fmt.Println("\n--- LAPORAN STATISTIK TREN CEKKESEHATANMENTAL ---")

	// 1. Analisis Tren Suasana Hati Mingguan (7 Hari Terakhir)
	totalSkorMingguan := 0
	jumlahDataMingguan := 0

	for i := 0; i < nMood; i++ {
		if apakahMingguIni(listMood[i].Tanggal) {
			totalSkorMingguan += listMood[i].SkorEmosi
			jumlahDataMingguan++
		}
	}

	if jumlahDataMingguan == 0 {
		fmt.Println(" [Mood] : Belum ada catatan suasana hati dalam 7 hari terakhir ini.")
	} else {
		rataRata := float64(totalSkorMingguan) / float64(jumlahDataMingguan)
		fmt.Printf(" [Mood] : Rata-rata Skor Emosimu Minggu Ini: %.2f / 10.0\n", rataRata)
		if rataRata <= 5.0 {
			fmt.Println("          Pesan Asisten: Grafik emosimu minggu ini agak menurun. Luangkan waktu istirahat ya!")
		} else {
			fmt.Println("          Pesan Asisten: Grafik suasana hatimu stabil dan baik. Pertahankan!")
		}
	}

	// 2. Analisis Tingkat Penyelesaian Tugas Harian
	tglHariIni := time.Now().Format("2006-01-02")
	totalTaskHariIni := 0
	taskSelesaiHariIni := 0

	for i := 0; i < nTask; i++ {
		if listTask[i].Tanggal == tglHariIni {
			totalTaskHariIni++
			if listTask[i].Selesai {
				taskSelesaiHariIni++
			}
		}
	}

	fmt.Println()
	if totalTaskHariIni == 0 {
		fmt.Printf(" [Task] : Tingkat Penyelesaian Tugas Hari Ini (%s): Belum ada target tugas.\n", tglHariIni)
	} else {
		persentase := (float64(taskSelesaiHariIni) / float64(totalTaskHariIni)) * 100
		fmt.Printf(" [Task] : Tingkat Penyelesaian Tugas Hari Ini (%s): %.2f%% selesai.\n", tglHariIni, persentase)
		fmt.Printf("          (%d dari %d tugas berhasil kamu selesaikan)\n", taskSelesaiHariIni, totalTaskHariIni)
	}
}
