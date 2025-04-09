package main // Программа которая принимает текст в виде строки, разбивает на слова, подсчитать сколько раз слова встречаются
//и вывести результат в виде списка слов и их частота в тексте, отсортировать по частоте

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "Мама мыла раму Мама мыла пол Папа не мыл раму Папа отдыхал"
	words := strings.Split(strings.ToLower(text), " ") // Вместо сплит есть функция Fields()
	counts := map[string]int{}                         // мапа повторений слов
	for _, word := range words {
		if _, ok := counts[word]; !ok {
			counts[word] = 0
		}
		counts[word]++
	}
	uniWords := make([]string, 0, len(counts))
	for word := range counts {
		uniWords = append(uniWords, word)
	}
	sort.Slice(uniWords, func(i, j int) bool { // Функция сортировки стоковая есть в интернете
		return counts[uniWords[i]] > counts[uniWords[j]]
	})
	for _, key := range uniWords {
		fmt.Printf("%s: %d\n", key, counts[key])
	}
	fmt.Println(counts)
}
