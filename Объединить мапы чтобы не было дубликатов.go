package main // Есть две мапы, нужно их объединить, чтобы не было дубликатов

import "fmt"

func main() {
	mapa1 := map[string]int{
		"Москва":          12000000,
		"Санкт-Петербург": 5000000,
	}
	mapa2 := map[string]int{
		"Новосибирск":     15000000,
		"Санкт-Петербург": 5000000,
	}
	mapa3 := make(map[string]int) // объем можно определить как len(mapa1)+len(mapa2)
	for key, i := range mapa1 {
		mapa3[key] = i
	}
	for key, i := range mapa2 {
		mapa3[key] = i
	}
	fmt.Println(mapa3)
}

// Второй вариант изменить только одну мапу и ее вывести, не создавая новую
// for key, i := range mapa1 {
//		mapa2[key] = i
//}
// fmt.Println(mapa2)
