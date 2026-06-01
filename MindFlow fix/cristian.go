package main

import (
	"bufio"
	"fmt"
	"strings"
)

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

func ubahTask(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang ingin diubah (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) 
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

func hapusTask(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang ingin dihapus (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) 
	idx := cariIndeksTaskBinary(*t, target)

	if idx != -1 {
		for i := idx; i < t.N-1; i++ {
			t.Data[i] = t.Data[i+1]
		}
		t.N--
		fmt.Println("Data tugas berhasil dihapus!")
	} else {
		fmt.Println("Tugas pada tanggal tersebut tidak ditemukan.")
	}
}

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

func cariIndeksTaskBinary(t TabTask, target string) int {
	left := 0
	right := t.N - 1
	idx := -1

	for left <= right && idx == -1 {
		mid := (left + right) / 2
		if t.Data[mid].Tanggal == target {
			idx = mid 
		} else if t.Data[mid].Tanggal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func cariTaskBinary(t *TabTask, scanner *bufio.Scanner) {
	fmt.Print("Masukkan Tanggal Task yang dicari (YYYY-MM-DD): ")
	scanner.Scan()
	target := scanner.Text()

	urutTaskByTanggal(t) 
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