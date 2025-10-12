package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	alamat := ":4000"
	server, err := net.ListenPacket("udp", alamat)
	if err != nil {
		fmt.Println("Gagal membuat server UDP:", err)
		return
	}
	defer server.Close()
	fmt.Println("Server UDP berjalan di", alamat)

	buffer := make([]byte, 1024)
	for {
		n, alamatKlien, err := server.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Kesalahan membaca:", err)
			continue
		}
		pesan := strings.TrimSpace(string(buffer[:n]))
		fmt.Printf("Dari %v: %s\n", alamatKlien, pesan)
		server.WriteTo([]byte("Pesan diterima: "+pesan), alamatKlien)
	}
}
