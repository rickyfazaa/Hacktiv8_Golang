package main

import (
	"fmt"
	"os"
	"strconv"
)

// Membuat Wadah Biodata
type Tabel struct {
	Nomor     string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Pengisian Tabel
var biodata = []Tabel{
	{
		Nomor:                        "1",
		Nama:                         "Ricky Khairul Faza",
		Alamat:                       "Jln. Buntu",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Certifikasi is number one slur!",
	},
	{
		Nomor:                        "2",
		Nama:                         "Agus Pro Klitih",
		Alamat:                       "Jln. Singapore KW",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Hanya saya yang tau!",
	},
	{
		Nomor:                        "3",
		Nama:                         "Rey Joice",
		Alamat:                       "Jln. Iklan",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Ingin jualan shampoo",
	},
	{
		Nomor:                        "4",
		Nama:                         "Van Persie",
		Alamat:                       "Jln. Belanda Raya",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Saya kira Go disini adalah Gol",
	},
	{
		Nomor:                        "5",
		Nama:                         "Bento",
		Alamat:                       "Jln. Peradaban",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Mencari angin segar",
	},
	{
		Nomor:                        "6",
		Nama:                         "Leonardo Samsudin",
		Alamat:                       "Jln. Merah",
		Pekerjaan:                    "Mahasiswa",
		Alasan /* memilih golang */ : "Ho-oh Tenann",
	},
}

func main() {
	check, _ := strconv.Atoi(os.Args[1])
	check = int(check)
	if check <= 0 || check >= 7 {
		fmt.Printf("Inputan minimal 1 dan mempunyai maximal 6")
		return
	} else {
		for _, cari := range biodata {
			if os.Args[1] == cari.Nomor {
				fmt.Printf("[ BIODATA NOMOR %v ]\nNama 		: %v\nAlamat		: %v\nPekerjaan	: %v\nAlasan		: %v\n", cari.Nomor, cari.Nama, cari.Alamat, cari.Pekerjaan, cari.Alasan)
			}
		}
	}
}
