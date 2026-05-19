package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Struktur Data
type MoodRecord struct {
	Tanggal   string // Format: YYYY-MM-DD [cite: 12]
	SkorEmosi int    // Skala 1-10 [cite: 11]
	Deskripsi string // [cite: 11]
}

type Task struct {
	Nama      string // [cite: 11]
	Durasi    int    // dalam menit [cite: 11]
	Prioritas int    // 1 (Tinggi), 2 (Sedang), 3 (Rendah) [cite: 13]
	Selesai   bool   // [cite: 14]
	Tanggal   string // Format: YYYY-MM-DD [cite: 12]
}

var listMood []MoodRecord
var listTask []Task
var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Println("================================================================")
	fmt.Println("   Halo! Selamat datang di RuangKesehatan - Asisten Virtual Anda.  ")
	fmt.Println("    Saya di sini untuk menjaga kesehatan mental & produktivitasmu. ")
	fmt.Println("================================================================")

	// Fitur Interaksi Awal Asisten Virtual (Check-in Mood Pengguna) [cite: 6]
	asistenCheckIn()

	for {
		fmt.Println("\n>>> [MENU UTAMA RUANGKESEHATAN] <<<")
		fmt.Println("1. Curhat / Perbarui Catatan Suasana Hati (Mood)")
		fmt.Println("2. Cari Catatan Mood atau Tugas (Pencarian)")
		fmt.Println("3. Urutkan Daftar Tugas Sesuai Prioritas/Durasi")
		fmt.Println("4. Lihat Analisis Tren Mental & Produktivitas")
		fmt.Println("5. Istirahat / Keluar Aplikasi")
		pilihan := readInput("Pilih menu (1-5): ")

		switch pilihan {
		case "1":
			menuMood()
		case "2":
			menuTask()
		case "3":
			menuUrut()
		case "4":
			tampilkanStatistik()
		case "5":
			fmt.Println("\n[RuangKesehatan]: Terima kasih sudah berproses hari ini. Jangan lupa istirahat, ya!")
			return
		default:
			fmt.Println("\n[RuangKesehatan]: Pilihan tidak tersedia, coba masukkan angka 1-5 ya.")
		}
	}
}

// ================= INTERAKSI ASISTEN VIRTUAL =================

func asistenCheckIn() {
	tglHariIni := time.Now().Format("2006-01-02")
	fmt.Printf("\n[RuangKesehatan]: Hari ini tanggal %s. Bagaimana keadaan hatimu sekarang?\n", tglHariIni)
	skorStr := readInput("Berikan skor emosimu hari ini (Skala 1-10): ")
	skor, _ := strconv.Atoi(skorStr)
	deskripsi := readInput("Ceritakan sedikit apa yang kamu rasakan hari ini: ")

	// Menambahkan ke riwayat percakapan/catatan mood [cite: 7, 10]
	listMood = append(listMood, MoodRecord{Tanggal: tglHariIni, SkorEmosi: skor, Deskripsi: deskripsi})

	// Respons Asisten yang personal berdasarkan Skor Emosi [cite: 6]
	fmt.Println("\n[RuangKesehatan]: Terima kasih sudah berbagi cerita.")
	if skor <= 4 {
		fmt.Println(">> Respon Asisten:  Tampaknya hari ini cukup berat buatmu. Ingat, tidak apa-apa untuk merasa lelah. Jangan terlalu memaksakan diri dalam bekerja hari ini, . Kesehatan mentalmu adalah yang utama. <<")
	} else if skor <= 7 {
		fmt.Println(">> Respon Asisten: Hari yang lumayan stabil. Tetap jaga ritme kerjamu dan jangan lupa ambil jeda istirahat di sela-sela kesibukan. <<")
	} else {
		fmt.Println(">> Respon Asisten: luar biasa! Energinya sangat positif hari ini. manfaatkan momentum baik ini untuk menyelesaikan target-targetmu! <<")
	}
}

// ================= KELOLA DATA (CRUD) =================

func menuMood() {
	fmt.Println("\n--- KELOLA RIWAYAT SUASANA HATI ---")
	fmt.Println("1. Tambah Catatan Baru")
	fmt.Println("2. Ubah Cerita/Skor Mood")
	fmt.Println("3. Hapus Catatan Mood")
	pilihan := readInput("Pilih aksi: ")

	switch pilihan {
	case "1": // [cite: 10]
		tgl := readInput("Masukkan Tanggal (YYYY-MM-DD): ")
		skorStr := readInput("Skor Emosi (1-10): ")
		skor, _ := strconv.Atoi(skorStr)
		deskripsi := readInput("Deskripsi Perasaan: ")
		listMood = append(listMood, MoodRecord{Tanggal: tgl, SkorEmosi: skor, Deskripsi: deskripsi})
		fmt.Println("[RuangKesehatan]: Catatan emosi berhasil disimpan.")
	case "2": // [cite: 10]
		tgl := readInput("Masukkan Tanggal yang ingin diubah (YYYY-MM-DD): ")
		idx := -1
		for i, m := range listMood {
			if m.Tanggal == tgl {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("[RuangKesehatan]: Maaf, data pada tanggal tersebut tidak ditemukan.")
			return
		}
		skorStr := readInput("Skor Emosi Baru (1-10): ")
		skor, _ := strconv.Atoi(skorStr)
		deskripsi := readInput("Deskripsi Perasaan Baru: ")
		listMood[idx].SkorEmosi = skor
		listMood[idx].Deskripsi = deskripsi
		fmt.Println("[RuangKesehatan]: Catatan emosimu berhasil diperbarui.")
	case "3": // [cite: 10]
		tgl := readInput("Masukkan Tanggal yang ingin dihapus (YYYY-MM-DD): ")
		idx := -1
		for i, m := range listMood {
			if m.Tanggal == tgl {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("[RuangKesehatan]: Data tidak ditemukan.")
			return
		}
		listMood = append(listMood[:idx], listMood[idx+1:]...)
		fmt.Println("[RuangKesehatan]: Catatan emosi berhasil dihapus.")
	}
}

func menuTask() {
	fmt.Println("\n--- DAFTAR TUGAS DAN PRODUKTIVITAS ---")
	fmt.Println("1. Tambah Tugas Baru")
	fmt.Println("2. Perbarui Status Tugas (Selesai/Belum)")
	fmt.Println("3. Hapus Tugas dari Daftar")
	pilihan := readInput("Pilih aksi: ")

	switch pilihan {
	case "1": // [cite: 10]
		nama := readInput("Nama Tugas / Aktivitas: ")
		durasiStr := readInput("Estimasi Durasi Pengerjaan (menit): ")
		durasi, _ := strconv.Atoi(durasiStr)
		prioStr := readInput("Prioritas Tugas (1=Tinggi, 2=Sedang, 3=Rendah): ")
		prio, _ := strconv.Atoi(prioStr)
		tgl := readInput("Target Tanggal Pengerjaan (YYYY-MM-DD): ")

		listTask = append(listTask, Task{Nama: nama, Durasi: durasi, Prioritas: prio, Tanggal: tgl, Selesai: false})
		fmt.Println("[RuangKesehatan]: Tugas baru berhasil ditambahkan ke daftar tokomu.")
	case "2": // [cite: 10]
		nama := readInput("Masukkan Nama Tugas yang ingin diubah statusnya: ")
		idx := -1
		for i, t := range listTask {
			if strings.EqualFold(t.Nama, nama) {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("[RuangKesehatan]: Tugas tidak ditemukan.")
			return
		}
		statusStr := readInput("Apakah tugas ini sudah kamu selesaikan? (y/n): ")
		if strings.ToLower(statusStr) == "y" {
			listTask[idx].Selesai = true
			fmt.Println("[RuangKesehatan]: nice! Satu beban tugas sudah diselesaikan. Good job!")
		} else {
			listTask[idx].Selesai = false
			fmt.Println("[RuangKesehatan]: Status tugas dikembalikan ke belum selesai.")
		}
	case "3": // [cite: 10]
		nama := readInput("Masukkan Nama Tugas yang ingin dihapus: ")
		idx := -1
		for i, t := range listTask {
			if strings.EqualFold(t.Nama, nama) {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("[RuangKesehatan]: Tugas tidak ditemukan.")
			return
		}
		listTask = append(listTask[:idx], listTask[idx+1:]...)
		fmt.Println("[RuangKesehatan]: Tugas berhasil dihapus dari daftar.")
	}
}

// ================= PENCARIAN (SEARCHING) =================

func menuCari() {
	fmt.Println("\n--- FITUR PENCARIAN ASISTEN ---")
	fmt.Println("1. Cari Tugas berdasarkan Kata Kunci (Sequential Search)")
	fmt.Println("2. Cari Catatan Mood berdasarkan Tanggal (Binary Search)")
	pilihan := readInput("Pilih opsi pencarian: ")

	if pilihan == "1" { // Sequential Search [cite: 12]
		keyword := readInput("Masukkan kata kunci tugas: ")
		fmt.Println("\n[RuangKesehatan]: Berikut hasil pencarian tugasmu:")
		ketemu := false
		for _, t := range listTask {
			if strings.Contains(strings.ToLower(t.Nama), strings.ToLower(keyword)) {
				fmt.Printf("- %s | Durasi: %d mnt | Prioritas: %d | Status Selesai: %t (%s)\n", t.Nama, t.Durasi, t.Prioritas, t.Selesai, t.Tanggal)
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("[RuangKesehatan]: Tidak ada nama tugas yang cocok dengan kata kunci tersebut.")
		}
	} else if pilihan == "2" { // Binary Search [cite: 12]
		if len(listMood) == 0 {
			fmt.Println("[RuangKesehatan]: Riwayat suasana hatimu masih kosong. Belum bisa mencari.")
			return
		}
		tgl := readInput("Masukkan tanggal yang ingin dicari (YYYY-MM-DD): ")

		// Binary search butuh data terurut berdasarkan key pencarian
		sort.Slice(listMood, func(i, j int) bool {
			return listMood[i].Tanggal < listMood[j].Tanggal
		})

		low, high := 0, len(listMood)-1
		idx := -1

		for low <= high {
			mid := (low + high) / 2
			if listMood[mid].Tanggal == tgl {
				idx = mid
				break
			} else if listMood[mid].Tanggal < tgl {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if idx != -1 {
			fmt.Printf("\n[RuangKesehatan]: Ketemu! Pada %s, skor emosimu adalah %d/10. Catatanmu: \"%s\"\n", listMood[idx].Tanggal, listMood[idx].SkorEmosi, listMood[idx].Deskripsi)
		} else {
			fmt.Println("[RuangKesehatan]: Saya tidak menemukan catatan mood untuk tanggal tersebut.")
		}
	}
}

// ================= PENGURUTAN (SORTING) =================

func menuUrut() {
	fmt.Println("\n--- PENGURUTAN TUGAS PRODUKTIVITAS ---")
	fmt.Println("1. Urutkan berdasarkan Tingkat Prioritas (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Durasi Singkat ke Lama (Insertion Sort)")
	pilihan := readInput("Pilih metode pengurutan: ")

	if pilihan == "1" { // Selection Sort [cite: 13]
		n := len(listTask)
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				if listTask[j].Prioritas < listTask[minIdx].Prioritas {
					minIdx = j
				}
			}
			listTask[i], listTask[minIdx] = listTask[minIdx], listTask[i]
		}
		fmt.Println("[RuangKesehatan]: Berhasil mengurutkan tugas dari prioritas tertinggi!")
		tampilkanSemuaTugas()

	} else if pilihan == "2" { // Insertion Sort [cite: 13]
		n := len(listTask)
		for i := 1; i < n; i++ {
			key := listTask[i]
			j := i - 1
			for j >= 0 && listTask[j].Durasi > key.Durasi {
				listTask[j+1] = listTask[j]
				j = j - 1
			}
			listTask[j+1] = key
		}
		fmt.Println("[RuangKesehatan]: Berhasil mengurutkan tugas dari durasi tersingkat!")
		tampilkanSemuaTugas()
	}
}

func tampilkanSemuaTugas() {
	for _, t := range listTask {
		prioText := "Rendah"
		if t.Prioritas == 1 {
			prioText = "Tinggi"
		} else if t.Prioritas == 2 {
			prioText = "Sedang"
		}
		fmt.Printf("- %s | Prioritas: %s | Durasi: %d mnt | Selesai: %t\n", t.Nama, prioText, t.Durasi, t.Selesai)
	}
}

// ================= STATISTIK TREN (KESEHATAN MENTAL & PRODUKTIVITAS) =================

func tampilkanStatistik() {
	fmt.Println("\n--- LAPORAN KESEHATAN MENTAL & PRODUKTIVITAS ---")

	// 1. Analisis Tren Suasana Hati [cite: 14]
	if len(listMood) == 0 {
		fmt.Println("[RuangKesehatan]: Belum ada data suasana hati yang cukup untuk dianalisis.")
	} else {
		totalSkor := 0
		for _, m := range listMood {
			totalSkor += m.SkorEmosi
		}
		rataRata := float64(totalSkor) / float64(len(listMood))
		fmt.Printf(" Tingkat Kondisi Emosimu: %.2f / 10.0\n", rataRata)

		if rataRata <= 5.0 {
			fmt.Println("  Catatan Asisten: Tren emosimu belakangan ini agak menurun. Mohon luangkan waktu untuk shoping .")
		} else {
			fmt.Println("  Catatan Asisten: Tren emosimu berada di grafik yang baik! Pertahankan kondisi kenyamanan pikiranmu.")
		}
	}

	// 2. Analisis Tingkat Penyelesaian Tugas Harian [cite: 14]
	tglHariIni := time.Now().Format("2006-01-02")
	totalTaskHariIni := 0
	taskSelesaiHariIni := 0

	for _, t := range listTask {
		if t.Tanggal == tglHariIni {
			totalTaskHariIni++
			if t.Selesai {
				taskSelesaiHariIni++
			}
		}
	}

	fmt.Println()
	if totalTaskHariIni == 0 {
		fmt.Printf(" Tingkat Penyelesaian Tugas Hari Ini (%s): Belum ada rencana tugas untuk hari ini.\n", tglHariIni)
	} else {
		persentase := (float64(taskSelesaiHariIni) / float64(totalTaskHariIni)) * 100
		fmt.Printf(" Tingkat Penyelesaian Tugas Hari Ini (%s): %.2f%% selesai.\n", tglHariIni, persentase)
		fmt.Printf("  (%d dari %d tugas berhasil kamu selesaikan)\n", taskSelesaiHariIni, totalTaskHariIni)
	}
}
