package main

import (
	"bufio"
	"fmt"
	"net"
	//"os"
	"strings"
)

func main() {
	alamat := ":3000"
	pendengar, err := net.Listen("tcp", alamat)
	if err != nil {
		fmt.Println("Gagal membuat server:", err)
		return
	}
	defer pendengar.Close()
	fmt.Println("Server TCP berjalan di", alamat)

	for {
		koneksi, err := pendengar.Accept()
		if err != nil {
			fmt.Println("Gagal menerima koneksi:", err)
			continue
		}
		fmt.Println("Klien terhubung:", koneksi.RemoteAddr())
		go tanganiKoneksi(koneksi)
	}
}

func tanganiKoneksi(koneksi net.Conn) {
	defer koneksi.Close()
	pembaca := bufio.NewReader(koneksi)

	for {
		pesan, err := pembaca.ReadString('\n')
		if err != nil {
			fmt.Println("Klien keluar:", koneksi.RemoteAddr())
			return
		}
		pesan = strings.TrimSpace(pesan)
		fmt.Println("Dari", koneksi.RemoteAddr(), ":", pesan)
		koneksi.Write([]byte("Pesan diterima: " + pesan + "\n"))
	}
}
