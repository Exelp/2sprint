package main // Вывести кто в мапе получили оценку больше 4 и посчитать количество тех кто с оценкой 5,и вывести общее арифмет

import "fmt"

func main() {
	students := map[string]int{
		"Тони Старк":   5,
		"Брюс Уэйн":    4,
		"Кларк Кент":   3,
		"Питер Паркер": 5,
		"Джеймс Бонд":  4,
	}
	sum := 0
	count := 0
	for name, ocenka := range students {
		sum += ocenka
		if ocenka > 4 {
			fmt.Println(name, ocenka)
		}
		if ocenka == 5 {
			count++
		}
	}
	fmt.Println("Количество тех кто получил 5:", count)
	fmt.Println(float64(sum) / float64(len(students)))

}
