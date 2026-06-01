package main

import (
	"bufio"
	"fmt"
	"os"
)

// Konstanta batas maksimum array
const MAX = 100

// Tipe bentukan untuk entitas
type Mood struct {
	Tanggal   string 
	Skor      int    
	Deskripsi string
}

type Task struct {
	Tanggal   string 
	Nama      string
	Durasi    int 
	Prioritas int 
	Selesai   bool
}

type TabMood struct {
	Data [MAX]Mood
	N    int
}

type TabTask struct {
	Data [MAX]Task
	N    int
}

// ============================================================================
// ALUR MENU UTAMA
// ============================================================================
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var tMood TabMood
	var tTask TabTask
	tMood.N = 0
	tTask.N = 0

	var menu int
	berjalan := true 

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
		scanner.Scan() 

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
			berjalan = false 
		} else {
			fmt.Println("Pilihan menu tidak valid!")
		}
	}
}

// ============================================================================
// FUNGSI MODUL MOOD (SEQUENTIAL)
// ============================================================================

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

func hapusMood(t *TabMood, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Mood yang ingin dihapus (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	idx := cariIndeksMoodSequential(*t, target)

	if idx != -1 {
		for i := idx; i < t.N-1; i++ {
			t.Data[i] = t.Data[i+1]
		}
		t.N--
		fmt.Println("Data mood berhasil dihapus!")
	} else {
		fmt.Println("Data mood pada tanggal tersebut tidak ditemukan.")
	}
}

func cariIndeksMoodSequential(t TabMood, target string) int {
	idx := -1
	i := 0
	for i < t.N && idx == -1 {
		if t.Data[i].Tanggal == target {
			idx = i 
		}
		i++
	}
	return idx
}