package main

import (
	"fmt"
	"strings"
)

func main() {
	//Link do desafio: https://leetcode.com/problems/repeated-substring-pattern/description/?envType=daily-question&envId=2023-11-29
	TestarFuncao("babbabbabbabbab", true)
	TestarFuncao("abab", true)
	TestarFuncao("aba", false)
	TestarFuncao("abcabcabcabc", true)
}

func TestarFuncao(s string, r bool) {
	if repeatedSubstringPattern(s) == r {
		fmt.Println("Teste concluido com sucesso para a entrada " + s)
	} else {
		fmt.Println("Teste falhou para a entrada " + s)
	}
}

// Solução
func repeatedSubstringPattern(s string) bool {
	tamanhoString := len(s)
	for i := 1; i < tamanhoString; i++ {
		subStr := s[0:i]
		QtdaRepeticoes := strings.Count(s, subStr)
		if (len(subStr) * QtdaRepeticoes) == tamanhoString {
			return true
		}
	}
	return false
}
