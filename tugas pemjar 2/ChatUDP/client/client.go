package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	alamat := "localhost:4000"
	koneksi, err := net.Dial("udp", alamat)
	if err != nil {
		fmt.Println("Gagal terhubung ke server UDP:", err)
		return
	}
	defer koneksi.Close()
	fmt.Println("Terhubung ke server UDP di", alamat)

	go func() {
		pembaca := bufio.NewReader(koneksi)
		for {
			pesan, _ := pembaca.ReadString('\n')
			fmt.Print(pesan)
		}
	}()

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
