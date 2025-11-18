package main

import "fmt"

func tarifParkir(jam int, membership bool, libur bool) (tarif int) {
	tarifDasar := 5000
	tarifLanjut := 2000
	diskon1 := 50
	diskon2 := 30
	tambahLibur := 3000
	switch {
	case jam <= 2:
		tarif = tarifDasar
	case jam > 2 && jam <= 5:
		if membership {
			tarif = tarifDasar + tarifLanjut*(1-diskon1/100)*(jam-2)
		} else {
			tarif = tarifDasar + tarifLanjut*(jam-2)
		}
	case jam > 5:
		if membership {
			tarif = tarifDasar + tarifLanjut*(1-diskon2/100)*(jam-2)
		} else {
			tarif = tarifDasar + tarifLanjut*(jam-2)
		}
	}

	if libur {
		tarif += tambahLibur
	}

	return tarif
}

func main() {
	fmt.Println("Orang pertama: Rp", tarifParkir(4, false, false))
	fmt.Println("Orang kedua: Rp", tarifParkir(2, true, true))
}