package main

import "fmt"

func main() {
	cetak("Hello, World!")
	cetak(1)
	cetak(true)
	cetak2("Hello, World!")
	cetak2(1)
	cetakSemua("Hello, World!", "dfd")
}

func cetak[T string | int | bool](arg T) {
	fmt.Println(arg)
}

func cetak2(arg any) {
	fmt.Println(arg)
}

func isEqual[t comparable](a, b t) bool {
	return a == b
}

func cetakSemua[T string | int | bool](args ...T) {
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	fmt.Println(args)
}
