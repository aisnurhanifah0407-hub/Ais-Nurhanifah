package main

import (
	"bufio"
	"fmt"
	"strings"
)

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

func urutTaskSelection(t *TabTask) {
	fmt.Print("Pilih urutan (1: Ascending (Naik), 2: Descending (Turun)): ")
	var pilihan int
	fmt.Scan(&pilihan)

	for i := 0; i < t.N-1; i++ {
		idxSasar := i
		for j := i + 1; j < t.N; j++ {
			if pilihan == 1 { 
				if t.Data[j].Prioritas < t.Data[idxSasar].Prioritas {
					idxSasar = j
				}
			} else { 
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

func urutTaskInsertion(t *TabTask) {
	fmt.Print("Pilih urutan (1: Ascending (Naik), 2: Descending (Turun)): ")
	var pilihan int
	fmt.Scan(&pilihan)

	for i := 1; i < t.N; i++ {
		key := t.Data[i]
		j := i - 1

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

func tampilSemuaTask(t TabTask) {
	if t.N == 0 {
		fmt.Println("Data task kosong.")
	}
	for i := 0; i < t.N; i++ {
		fmt.Printf("- %s (Tanggal: %s | Durasi: %d | Prioritas: %d)\n", t.Data[i].Nama, t.Data[i].Tanggal, t.Data[i].Durasi, t.Data[i].Prioritas)
	}
}

func tampilStatistik(tMood TabMood, tTask TabTask) {
	fmt.Println("\n+++ MindFlow +++")
	fmt.Println("++++ Statistik Tren Suasana Hati & Productivity ++++")

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