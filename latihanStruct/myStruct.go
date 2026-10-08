package latihanstruct

import "fmt"

type Todo struct {
	ID     int
	Title  string
	IsDone bool
}

func MyStruct() {
	varBaru1 := Todo{1, "belajar go", false}

	varBaru2 := Todo{
		ID:     2,
		Title:  "istirahat",
		IsDone: false,
	}

	fmt.Println("varBaru1 =", varBaru1)
	fmt.Println("varBaru2 =", varBaru2)
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Println("varBaru1 =", "saya sedang ", varBaru1.Title, "sudah selesai ", varBaru1.IsDone)
}

type Alamat struct {
	Kota     string
	Provinsi string
}

type Siswa struct {
	Nama     string
	Nilai    int
	Domisili Alamat
}

func MyStruct2() {

	siswa1 := Siswa{
		Nama:  "Kuncoro",
		Nilai: 80,
		Domisili: Alamat{
			Kota:     "Klaten",
			Provinsi: "Jateng",
		},
	}

	fmt.Println(siswa1)

	daftarSiswa := []Siswa{
		{
			Nama:  "Ali",
			Nilai: 70,
		},
		{
			Nama:  "Budi",
			Nilai: 80,
		},
	}

	fmt.Println(len(daftarSiswa))

	i := 0
	for i < len(daftarSiswa) {

		fmt.Println("Siswa ", i, ":")
		fmt.Println("Nama :", daftarSiswa[i].Nama)
		fmt.Println("Nilai :", daftarSiswa[i].Nilai)

		lulus := "tidak lulus"

		if daftarSiswa[i].Nilai >= 75 {
			lulus = "lulus"
		}

		fmt.Println("Lulus :", lulus)

		fmt.Println("==========================")

		i++
	}

}
