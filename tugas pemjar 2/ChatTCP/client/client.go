package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	alamat := "localhost:3000"
	koneksi, err := net.Dial("tcp", alamat)
	if err != nil {
		fmt.Println("Gagal terhubung ke server:", err)
		return
	}
	defer koneksi.Close()
	fmt.Println("Terhubung ke server di", alamat)

	// Baca dari server
	go func() {
		pembaca := bufio.NewScanner(koneksi)
		for pembaca.Scan() {
			fmt.Println(pembaca.Text())
		}
	}()

	// Kirim dari stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		pesan := strings.TrimSpace(scanner.Text())
		if pesan == "keluar" {
			fmt.Println("Keluar dari chat...")
			return
		}
		fmt.Fprintln(koneksi, pesan)
	}
}
