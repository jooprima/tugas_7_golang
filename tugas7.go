package main

import "fmt"
import "reflect"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var nomor = 12
	var huruf = "abc"
	var bacatipe1 = reflect.ValueOf(nomor)
	var bacatipe2 = reflect.ValueOf(huruf)

	fmt.Println("Tipe variabel :", bacatipe1.Type())
	fmt.Println("Nilai variabel :", bacatipe1.Int())
	fmt.Println("Tipe variabel :", bacatipe2.Type())
	fmt.Println("Nilai variabel :", bacatipe2.String())

}
