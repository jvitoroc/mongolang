package main

import (
	"fmt"
	"mongolang/expr"
)

func main() {
	t := expr.NewTokenizer()
	for _, c := range t.Tokenize("(answers.status='really complex string with \"quoted substring\"')and modiefiedCount>123213.123123") {
		fmt.Println(c)
	}
}
