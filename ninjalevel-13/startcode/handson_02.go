package main

import (
	"fmt"
	"go_sri/ninjalevel-13/quote"
	"go_sri/ninjalevel-13/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}