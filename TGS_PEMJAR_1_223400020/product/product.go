package barang

import "fmt"

type Barang struct {
	ID    int
	Nama  string
	Harga float64
}

func Baru(id int, nama string, harga float64) Barang {
	return Barang{
		ID:    id,
		Nama:  nama,
		Harga: harga,
	}
}

func (b Barang) Info() {
	fmt.Printf("[%d] %s - Rp%.2f\n", b.ID, b.Nama, b.Harga)
}
