package main

import (
	"fmt"
	"TGS_PEMJAR_1_223400020/product"
)

func main() {
	fmt.Println("=== Daftar Barang ===")

	// Buat data barang
	b1 := barang.Baru(1, "Buku", 25000)
	b2 := barang.Baru(2, "Pulpen", 5000)

	// Tampilkan info
	b1.Info()
	b2.Info()
}