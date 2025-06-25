package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Deskripsi string
	Done        bool
}

var tasks []Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n____ Program List Tugas ____")
		fmt.Println("1. Lihat Semua Tugas")
		fmt.Println("2. Tambah Tugas Baru")
		fmt.Println("3. Tandai Tugas Sudah Selesai")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5) : ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			listTasks()
		case "2":
			addTask(scanner)
		case "3":
			completeTask(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Kamu telah keluar dari program.")
			return
		default:
			fmt.Println("Pilihan enggak valid, coba lagi deh.")
		}
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("\n Masih Kosong ler")
		return
	}
	fmt.Println("\nDaftar Tugas:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Deskripsi)
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan deskripsi tugas: ")
	scanner.Scan()
	desc := scanner.Text()
	tasks = append(tasks, Task{Deskripsi: desc})
	fmt.Println("Tugas berhasil ditambahkan.")
}

func completeTask(scanner *bufio.Scanner) {
	listTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Masukkan nomor tugas yang selesai: ")
	scanner.Scan()
	input := scanner.Text()
	i, err := strconv.Atoi(input)
	if err != nil || i < 1 || i > len(tasks) {
		fmt.Println("Nomor tugas tidak valid.")
		return
	}
	tasks[i-1].Done = true
	fmt.Println("Tugas ditandai sebagai selesai.")
}

func deleteTask(scanner *bufio.Scanner) {
	listTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
	scanner.Scan()
	input := scanner.Text()
	i, err := strconv.Atoi(input)
	if err != nil || i < 1 || i > len(tasks) {
		fmt.Println("Nomor tugas engga valid.")
		return
	}
	tasks = append(tasks[:i-1], tasks[i:]...)
	fmt.Println("Tugas berhasil dihapus.")
}