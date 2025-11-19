package main

import "fmt"

var diskon1 float64 = 50.0
var diskon2 float64 = 30.0

func tarifParkir(jam float64, membership bool, libur bool) (tarif float64) {
	tarifDasar := 5000.0
	tarifLanjut := 2000.0
	tambahLibur := 3000.0
	switch {
	case jam <=0:
		tarif = 0
	case jam >0 && jam <= 2:
		tarif+=tarifDasar
	case jam > 2:
		tarif+=tarifDasar
		tarif+=tarifLanjut*(jam-2)
	}

	if membership {
		if jam <= 5 {
			tarif*=(1 - (diskon1/100))
		} else {
			tarif*=(1 - (diskon2/100))
		}
	}
	
	if libur {
		tarif += tambahLibur
	}

	return tarif
}

type Parkir struct {
	jam float64
	membership bool
	libur bool
}

func main() {
	scenario1 := Parkir{
		jam: 4.0,
		membership: false,
		libur: false,
	}

	scenario2 := Parkir{
		jam: 2.0,
		membership: true,
		libur: true,
	}

	fmt.Println("Orang pertama: Rp", tarifParkir(scenario1.jam, scenario1.membership, scenario1.libur))
	fmt.Println("Orang kedua: Rp", tarifParkir(scenario2.jam, scenario2.membership, scenario2.libur))
}