package main

// Aplikasi Asisten Virtual Kesehatan Mental dan Produktivitas (MindFlow)
// Proyek Kelompok: Ais, Cristian, Rian

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Konstanta batas maksimum array
const MAX = 100

// Tipe bentukan untuk entitas
type Mood struct {
	Tanggal   string // Format: YYYY-MM-DD
	Skor      int    // Skala 1-10
	Deskripsi string
}

type Task struct {
	Tanggal   string // Format: YYYY-MM-DD
	Nama      string
	Durasi    int // Waktu pengerjaan dalam menit
	Prioritas int // 1 (Tinggi), 2 (Sedang), 3 (Rendah)
	Selesai   bool
}

// Tipe bentukan untuk membungkus Array Statis dan Counternya
type TabMood struct {
	Data [MAX]Mood
	N    int
}

type TabTask struct {
	Data [MAX]Task
	N    int
}

// ============================================================================
// PROGRAM UTAMA
// Penanggung Jawab: Ais
// ============================================================================
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var tMood TabMood
	var tTask TabTask
	tMood.N = 0
	tTask.N = 0

	var menu int
	berjalan := true // Pengganti statement 'break' untuk loop utama

	for berjalan {
		fmt.Println("\n=== MindFlow: Asisten Kesehatan Mental & Produktivitas ===")
		fmt.Println("1. Tambah Mood")
		fmt.Println("2. Ubah Mood (via Sequential Search)")
		fmt.Println("3. Hapus Mood (via Sequential Search)")
		fmt.Println("4. Tambah Task")
		fmt.Println("5. Ubah Task (via Binary Search)")
		fmt.Println("6. Hapus Task (via Binary Search)")
		fmt.Println("7. Cari Mood (Sequential Search - By Keyword)")
		fmt.Println("8. Cari Task (Binary Search - By Tanggal)")
		fmt.Println("9. Urutkan Task (Selection Sort - By Prioritas Asc/Desc)")
		fmt.Println("10. Urutkan Task (Insertion Sort - By Durasi Asc/Desc)")
		fmt.Println("11. Tampilkan Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&menu)
		scanner.Scan() // Clear buffer

		if menu == 1 {
			tambahMood(&tMood, scanner)
		} else if menu == 2 {
			ubahMood(&tMood, scanner)
		} else if menu == 3 {
			hapusMood(&tMood, scanner)
		} else if menu == 4 {
			tambahTask(&tTask, scanner)
		} else if menu == 5 {
			ubahTask(&tTask, scanner)
		} else if menu == 6 {
			hapusTask(&tTask, scanner)
		} else if menu == 7 {
			cariMoodSequential(tMood, scanner)
		} else if menu == 8 {
			cariTaskBinary(&tTask, scanner)
		} else if menu == 9 {
			urutTaskSelection(&tTask)
		} else if menu == 10 {
			urutTaskInsertion(&tTask)
		} else if menu == 11 {
			tampilStatistik(tMood, tTask)
		} else if menu == 0 {
			fmt.Println("Terima kasih telah menggunakan MindFlow!")
			berjalan = false // Keluar loop tanpa break
		} else {
			fmt.Println("Pilihan menu tidak valid!")
		}
	}
}

// ============================================================================
// SUBPROGRAM: MODUL MOOD (SEQUENTIAL FLOW)
// Penanggung Jawab: Ais
// ============================================================================

// I.S. : tMood terdefinisi, mungkin kosong
// F.S. : Data mood baru bertambah jika kapasitas belum penuh
func tambahMood(t *TabMood, scanner *bufio.Scanner) {
	if t.N >= MAX {
		fmt.Println("Kapasitas penyimpanan mood penuh!")
	} else {
		fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
		scanner.Scan()
		tgl := scanner.Text()

		fmt.Print("Masukkan Skor Emosi (1-10): ")
		var skor int
		fmt.Scan(&skor)
		scanner.Scan()

		fmt.Print("Masukkan Deskripsi Perasaan: ")
		scanner.Scan()
		deskripsi := scanner.Text()

		t.Data[t.N] = Mood{Tanggal: tgl, Skor: skor, Deskripsi: deskripsi}
		t.N++
		fmt.Println("Data mood berhasil ditambahkan!")
	}
}

// I.S. : tMood terdefinisi
// F.S. : Mengubah data mood yang dicari menggunakan Sequential Search
func ubahMood(t *TabMood, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Mood yang ingin diubah (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	idx := cariIndeksMoodSequential(*t, target)

	if idx != -1 {
		fmt.Println("Data ditemukan. Masukkan data baru:")
		fmt.Print("Masukkan Skor Emosi baru (1-10): ")
		var skor int
		fmt.Scan(&skor)
		scanner.Scan()

		fmt.Print("Masukkan Deskripsi Perasaan baru: ")
		scanner.Scan()
		deskripsi := scanner.Text()

		t.Data[idx].Skor = skor
		t.Data[idx].Deskripsi = deskripsi
		fmt.Println("Data mood berhasil diubah!")
	} else {
		fmt.Println("Data mood pada tanggal tersebut tidak ditemukan.")
	}
}

// I.S. : tMood terdefinisi
// F.S. : Menghapus data mood yang dicari menggunakan Sequential Search
func hapusMood(t *TabMood, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Mood yang ingin dihapus (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	idx := cariIndeksMoodSequential(*t, target)

	if idx != -1 {
		// Geser array ke kiri untuk menimpa data yang dihapus
		for i := idx; i < t.N-1; i++ {
			t.Data[i] = t.Data[i+1]
		}
		t.N--
		fmt.Println("Data mood berhasil dihapus!")
	} else {
		fmt.Println("Data mood pada tanggal tersebut tidak ditemukan.")
	}
}

// I.S. : tMood terdefinisi, target berisi tanggal yang dicari
// F.S. : Mengembalikan indeks dari data yang dicari, atau -1 jika tidak ditemukan
func cariIndeksMoodSequential(t TabMood, target string) int {
	idx := -1
	i := 0
	for i < t.N && idx == -1 {
		if t.Data[i].Tanggal == target {
			idx = i // Pengganti break
		}
		i++
	}
	return idx
}

// ============================================================================
// SUBPROGRAM: MODUL TASK & BINARY SEARCH
// Penanggung Jawab: Cristian
// ============================================================================

// I.S. : tTask terdefinisi, mungkin kosong
// F.S. : Data task baru bertambah jika kapasitas belum penuh
func tambahTask(t *TabTask, scanner *bufio.Scanner) {
	if t.N >= MAX {
		fmt.Println("Kapasitas penyimpanan tugas penuh!")
	} else {
		fmt.Print("Masukkan Tanggal Tugas (YYYY-MM-DD): ")
		scanner.Scan()
		tgl := scanner.Text()

		fmt.Print("Masukkan Nama Tugas: ")
		scanner.Scan()
		nama := scanner.Text()

		fmt.Print("Masukkan Durasi (menit): ")
		var durasi int
		fmt.Scan(&durasi)

		fmt.Print("Masukkan Prioritas (1: Tinggi, 2: Sedang, 3: Rendah): ")
		var prioritas int
		fmt.Scan(&prioritas)
		scanner.Scan()

		t.Data[t.N] = Task{Tanggal: tgl, Nama: nama, Durasi: durasi, Prioritas: prioritas, Selesai: false}
		t.N++
		fmt.Println("Data tugas berhasil ditambahkan!")
	}
}

// I.S. : tTask terdefinisi
// F.S. : Mengubah data task yang dicari menggunakan Binary Search
func ubahTask(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang ingin diubah (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) // Wajib diurutkan sebelum Binary Search
	idx := cariIndeksTaskBinary(*t, target)

	if idx != -1 {
		fmt.Println("Data ditemukan. Masukkan data baru:")
		fmt.Print("Masukkan Nama Tugas baru: ")
		scanner.Scan()
		nama := scanner.Text()

		fmt.Print("Masukkan Durasi baru (menit): ")
		var durasi int
		fmt.Scan(&durasi)

		fmt.Print("Masukkan Prioritas baru (1: Tinggi, 2: Sedang, 3: Rendah): ")
		var prioritas int
		fmt.Scan(&prioritas)
		scanner.Scan()

		fmt.Print("Apakah tugas sudah selesai? (y/n): ")
		scanner.Scan()
		selesai := strings.ToLower(scanner.Text()) == "y"

		t.Data[idx].Nama = nama
		t.Data[idx].Durasi = durasi
		t.Data[idx].Prioritas = prioritas
		t.Data[idx].Selesai = selesai
		fmt.Println("Data tugas berhasil diubah!")
	} else {
		fmt.Println("Tugas pada tanggal tersebut tidak ditemukan.")
	}
}

// I.S. : tTask terdefinisi
// F.S. : Menghapus data task yang dicari menggunakan Binary Search
func hapusTask(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang ingin dihapus (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) // Wajib diurutkan sebelum Binary Search
	idx := cariIndeksTaskBinary(*t, target)

	if idx != -1 {
		// Geser elemen array
		for i := idx; i < t.N-1; i++ {
			t.Data[i] = t.Data[i+1]
		}
		t.N--
		fmt.Println("Data tugas berhasil dihapus!")
	} else {
		fmt.Println("Tugas pada tanggal tersebut tidak ditemukan.")
	}
}

// I.S. : tTask terdefinisi, mungkin belum terurut berdasarkan tanggal
// F.S. : Elemen di dalam tTask terurut secara ascending berdasarkan tanggal (Untuk syarat Binary Search)
func urutTaskByTanggal(t *TabTask) {
	for i := 1; i < t.N; i++ {
		key := t.Data[i]
		j := i - 1
		for j >= 0 && t.Data[j].Tanggal > key.Tanggal {
			t.Data[j+1] = t.Data[j]
			j = j - 1
		}
		t.Data[j+1] = key
	}
}

// I.S. : tTask terdefinisi dan sudah terurut berdasarkan tanggal, target berisi tanggal yang dicari
// F.S. : Mengembalikan indeks dari data task yang dicari, atau -1 jika tidak ditemukan
func cariIndeksTaskBinary(t TabTask, target string) int {
	left := 0
	right := t.N - 1
	idx := -1

	for left <= right && idx == -1 {
		mid := (left + right) / 2
		if t.Data[mid].Tanggal == target {
			idx = mid // Kondisi ini akan menghentikan perulangan (pengganti break)
		} else if t.Data[mid].Tanggal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

// I.S. : tTask terdefinisi
// F.S. : Menampilkan hasil pencarian berdasarkan tanggal menggunakan Binary Search
func cariTaskBinary(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang dicari (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) // Wajib urut untuk Binary Search
	idx := cariIndeksTaskBinary(*t, target)

	fmt.Println("\n--- Hasil Pencarian Task ---")
	if idx != -1 {
		status := "Belum"
		if t.Data[idx].Selesai {
			status = "Selesai"
		}
		fmt.Printf("Ditemukan: %s | Durasi: %d mnt | Prioritas: %d | Status: %s\n", t.Data[idx].Nama, t.Data[idx].Durasi, t.Data[idx].Prioritas, status)
	} else {
		fmt.Println("Tidak ada tugas pada tanggal tersebut.")
	}
}

// ============================================================================
// SUBPROGRAM: MODUL SORTING & ANALISIS DATA
// Penanggung Jawab: Rian
// ============================================================================

// I.S. : tMood terdefinisi
// F.S. : Menampilkan hasil pencarian berdasarkan keyword deskripsi (Sequential)
func cariMoodSequential(t TabMood, scanner *bufio.Scanner) {
	fmt.Print("Masukkan kata kunci deskripsi (contoh: sedih, senang): ")
	scanner.Scan()
	keyword := scanner.Text()

	fmt.Println("\n--- Hasil Pencarian Mood ---")
	ketemu := false

	for i := 0; i < t.N; i++ {
		if strings.Contains(strings.ToLower(t.Data[i].Deskripsi), strings.ToLower(keyword)) {
			fmt.Printf("Tanggal: %s | Skor: %d | Deskripsi: %s\n", t.Data[i].Tanggal, t.Data[i].Skor, t.Data[i].Deskripsi)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Catatan mood dengan kata kunci tersebut tidak ditemukan.")
	}
}

// I.S. : tTask terdefinisi
// F.S. : Mengurutkan Task berdasarkan Prioritas (Mendukung Ascending & Descending via Selection Sort)
func urutTaskSelection(t *TabTask) {
	fmt.Print("Pilih urutan (1: Ascending (Naik), 2: Descending (Turun)): ")
	var pilihan int
	fmt.Scan(&pilihan)

	for i := 0; i < t.N-1; i++ {
		idxSasar := i
		for j := i + 1; j < t.N; j++ {
			if pilihan == 1 { // Ascending
				if t.Data[j].Prioritas < t.Data[idxSasar].Prioritas {
					idxSasar = j
				}
			} else { // Descending
				if t.Data[j].Prioritas > t.Data[idxSasar].Prioritas {
					idxSasar = j
				}
			}
		}
		temp := t.Data[i]
		t.Data[i] = t.Data[idxSasar]
		t.Data[idxSasar] = temp
	}

	fmt.Println("\nDaftar tugas berhasil diurutkan berdasarkan Prioritas!")
	tampilSemuaTask(*t)
}

// I.S. : tTask terdefinisi
// F.S. : Mengurutkan Task berdasarkan Durasi (Mendukung Ascending & Descending via Insertion Sort)
func urutTaskInsertion(t *TabTask) {
	fmt.Print("Pilih urutan (1: Ascending (Naik), 2: Descending (Turun)): ")
	var pilihan int
	fmt.Scan(&pilihan)

	for i := 1; i < t.N; i++ {
		key := t.Data[i]
		j := i - 1

		// Menggunakan flag boolean agar tidak pakai break/continue
		geser := true
		for j >= 0 && geser {
			if pilihan == 1 && t.Data[j].Durasi > key.Durasi {
				t.Data[j+1] = t.Data[j]
				j = j - 1
			} else if pilihan == 2 && t.Data[j].Durasi < key.Durasi {
				t.Data[j+1] = t.Data[j]
				j = j - 1
			} else {
				geser = false
			}
		}
		t.Data[j+1] = key
	}

	fmt.Println("\nDaftar tugas berhasil diurutkan berdasarkan Durasi!")
	tampilSemuaTask(*t)
}

// I.S. : tTask terdefinisi
// F.S. : Menampilkan seluruh daftar tugas ke layar
func tampilSemuaTask(t TabTask) {
	if t.N == 0 {
		fmt.Println("Data task kosong.")
	}
	for i := 0; i < t.N; i++ {
		fmt.Printf("- %s (Tanggal: %s | Durasi: %d | Prioritas: %d)\n", t.Data[i].Nama, t.Data[i].Tanggal, t.Data[i].Durasi, t.Data[i].Prioritas)
	}
}

// I.S. : tMood dan tTask terdefinisi
// F.S. : Menampilkan ringkasan statistik kontribusi tren data ke layar
func tampilStatistik(tMood TabMood, tTask TabTask) {
	fmt.Println("\n+++ MindFlow +++")
	fmt.Println("++++ Statistik Tren Suasana Hati & Produktivitas ++++")

	var totalSkor int = 0
	for i := 0; i < tMood.N; i++ {
		totalSkor += tMood.Data[i].Skor
	}

	if tMood.N > 0 {
		rataRata := float64(totalSkor) / float64(tMood.N)
		fmt.Printf("\n1. Rata-rata Skor Suasana Hati: %.2f / 10\n", rataRata)
	} else {
		fmt.Println("\n1. Rata-rata Skor Suasana Hati: Belum ada data.")
	}

	var taskSelesai int = 0
	for i := 0; i < tTask.N; i++ {
		if tTask.Data[i].Selesai {
			taskSelesai++
		}
	}

	if tTask.N > 0 {
		persentase := (float64(taskSelesai) / float64(tTask.N)) * 100
		fmt.Printf("2. Tingkat Penyelesaian Task: %.2f%%\n", persentase)
	} else {
		fmt.Println("2. Tingkat Penyelesaian Task: Belum ada data.")
	}

	fmt.Println("\n+++ MindFlow +++\n")
}