package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DataPolusi struct {
	Kota    string
	AQI     int
	Sumber  string
	Tanggal string
}

var data []DataPolusi

func main() {
	data = []DataPolusi{
		{"Purwokerto", 100, "Kendaraan", "17-05-2025"},
		{"Cilacap", 150, "Pabrik", "14-04-2025"},
		{"Tegal", 200, "Pembakaran", "12-02-2025"},
		{"Wonosobo", 50, "Kendaraan", "13-03-2025"},
		{"Magelang", 100, "Kendaraan", "11-10-2025"},
		{"Cikarang", 200, "Pabrik", "09-01-2025"},
	}
	menu()
}

func status(aqi int) string {
	if aqi <= 50 {
		return "Baik"
	} else if aqi <= 100 {
		return "Sedang"
	} else if aqi <= 150 {
		return "Sensitif"
	} else if aqi <= 200 {
		return "Tidak Sehat"
	} else if aqi <= 300 {
		return "Sangat Tidak Sehat"
	}
	return "Berbahaya"
}

func tambahData() {
	var kota, sumber, tanggal string
	var aqi int

	fmt.Print("Kota: ")
	fmt.Scanln(&kota)
	fmt.Print("AQI: ")
	_, err := fmt.Scanln(&aqi)
	if err != nil {
		fmt.Println(" ❌ Masukan AQI tidak valid. Silakan masukkan angka.")
		return
	}
	fmt.Print("Sumber: ")
	fmt.Scanln(&sumber)
	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scanln(&tanggal)
	if aqi > 0 && aqi <= 500 {
		data = append(data, DataPolusi{Kota: kota, AQI: aqi, Sumber: sumber, Tanggal: tanggal})
		fmt.Println(" ✅ Data berhasil ditambahkan.")
	} else {
		fmt.Println(" ❌ Masukan AQI di luar rentang valid (1-500).")
	}
}

func ubahData() {
	var kota, tanggal string
	fmt.Print("Masukkan kota yang ingin diubah: ")
	fmt.Scanln(&kota)
	fmt.Print("Masukkan tanggal pengukuran (DD-MM-YYYY): ")
	fmt.Scanln(&tanggal)

	for i := range data {
		if strings.EqualFold(data[i].Kota, kota) && data[i].Tanggal == tanggal {
			fmt.Print("Masukkan AQI baru: ")
			fmt.Scanln(&data[i].AQI)
			fmt.Print("Masukkan sumber baru: ")
			fmt.Scanln(&data[i].Sumber)
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println(" ❌ Data tidak ditemukan.")
}

func hapusData() {
	var kota, tanggal string
	fmt.Print("Masukkan kota yang ingin dihapus: ")
	fmt.Scanln(&kota)
	fmt.Print("Masukkan tanggal pengukuran (DD-MM-YYYY): ")
	fmt.Scanln(&tanggal)

	for i := range data {
		if strings.EqualFold(data[i].Kota, kota) && data[i].Tanggal == tanggal {
			data = append(data[:i], data[i+1:]...)
			fmt.Println(" ✅ Data berhasil dihapus.")
			return
		}
	}
	fmt.Println(" ❌ Data tidak ditemukan.")
}

func tampilData() {
	fmt.Println("+------------+------------+-----+-------------------+")
	fmt.Println("|  Tanggal   |   Kota     | AQI |      Sumber       |")
	fmt.Println("+------------+------------+-----+-------------------+")
	for _, d := range data {
		fmt.Printf("| %-10s | %-10s | %3d | %-17s |\n",
			d.Tanggal, d.Kota, d.AQI, d.Sumber)
	}
	fmt.Println("+------------+------------+-----+-------------------+")
}

func urutkanAQIInsertionSort() {
	n := len(data)
	for i := 1; i < n; i++ {
		x := data[i]
		j := i - 1
		for j >= 0 && data[j].AQI > x.AQI {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = x
	}
	fmt.Println(" ✅ Data diurutkan dari AQI terendah (Insertion Sort).")
}

func urutkanAQISelectionSort() {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if data[j].AQI < data[idxMin].AQI {
				idxMin = j
			}
		}
		if idxMin != i {
			data[i], data[idxMin] = data[idxMin], data[i]
		}
	}
	fmt.Println(" ✅ Data diurutkan dari AQI terendah (Selection Sort).")
}

func cariSequential() {
	var kota string
	fmt.Print("Masukkan nama kota yang ingin dicari: ")
	fmt.Scanln(&kota)

	found := false
	for _, d := range data {
		if strings.EqualFold(d.Kota, kota) {
			fmt.Printf("%v | %v | %v | %v | %v\n", d.Tanggal, d.Kota, d.AQI, status(d.AQI), d.Sumber)
			found = true
		}
	}
	if !found {
		fmt.Println(" ❌ Data tidak ditemukan.")
	}
}

func cariBinary() {
	urutkanDataByKota()

	var kota string
	fmt.Print("Masukkan nama kota yang ingin dicari: ")
	fmt.Scanln(&kota)

	kiri, kanan := 0, len(data)-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if strings.EqualFold(data[tengah].Kota, kota) {
			fmt.Printf("%v | %v | %v | %v | %v\n", data[tengah].Tanggal, data[tengah].Kota, data[tengah].AQI, status(data[tengah].AQI), data[tengah].Sumber)
			return
		} else if strings.ToLower(data[tengah].Kota) < strings.ToLower(kota) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println(" ❌ Data tidak ditemukan.")
}


func urutkanDataByKota() {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(data[j].Kota) < strings.ToLower(data[idxMin].Kota) {
				idxMin = j
			}
		}
		if idxMin != i {
			data[i], data[idxMin] = data[idxMin], data[i]
		}
	}
}

func tampilkanDenganStatus() {
	var st string
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan status polusi yang ingin ditampilkan (Baik/Sedang/Sensitif/Tidak Sehat/Sangat Tidak Sehat/Berbahaya): ")
	input, _ := read.ReadString('\n')
	st = strings.TrimSpace(input)

	ditemukan := false
	fmt.Println("+------------+------------+-----+-------------------------------+----------------+")
	fmt.Println("|  Tanggal   |   Kota     | AQI |           Status              |     Sumber     |")
	fmt.Println("+------------+------------+-----+-------------------------------+----------------+")

	for _, d := range data {
		if strings.EqualFold(status(d.AQI), st) {
			fmt.Printf("| %-10v | %-10v | %3v | %-29v | %-14v |\n",
				d.Tanggal, d.Kota, d.AQI, status(d.AQI), d.Sumber)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println(" ❌ Tidak ada data dengan status tersebut.")
	}
	fmt.Println("+------------+------------+-----+-------------------------------+----------------+")
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n+========== MENU ===========+")
		fmt.Println("| 1. Tambah Data            |")
		fmt.Println("| 2. Ubah Data              |")
		fmt.Println("| 3. Hapus Data             |")
		fmt.Println("| 4. Tampilkan Semua        |")
		fmt.Println("| 5. Urutkan AQI (Selec)    |")
		fmt.Println("| 6. Urutkan AQI (Insert)   |")
		fmt.Println("| 7. Cari Sequential        |")
		fmt.Println("| 8. Cari Binary            |")
		fmt.Println("| 9. Tampilkan Dengan Status|")
		fmt.Println("| 0. Keluar                 |")
		fmt.Println("+===========================+")
		fmt.Print("Pilih menu (0-9): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		pilih := input

		switch pilih {
		case "1":
			tambahData()
		case "2":
			ubahData()
		case "3":
			hapusData()
		case "4":
			tampilData()
		case "5":
			urutkanAQISelectionSort()
		case "6":
			urutkanAQIInsertionSort()
		case "7":
			cariSequential()
		case "8":
			cariBinary()
		case "9":
			tampilkanDenganStatus()
		case "0":
			fmt.Println(" ✅ Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println(" ❌ Pilihan tidak valid.")
		}
	}
}