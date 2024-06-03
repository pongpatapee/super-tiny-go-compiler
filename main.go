package main

import "fmt"

type token struct {
	kind  string
	value string
}

func isNumber(char string) bool {
	if char == "" {
		return false
	}

	n := []rune(char)[0]

	if n >= '0' && n <= '9' {
		return true
	}

	return false
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}

	n := []rune(char)[0]

	if n >= 'a' && n <= 'z' {
		return true
	}

	return false
}

func tokenizer(input string) []token {
	// add new line to program
	input += "\n"

	// cursor for going through input
	current := 0

	tokens := []token{}

	for current < len([]rune(input)) {
		char := string([]rune(input)[current])

		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})

			current++
			continue
		}

		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})

			current++
			continue
		}

		if char == " " {
			current++
			continue
		}

		if isNumber(char) {

			value := ""
			for isNumber(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}

			tokens = append(tokens, token{
				kind:  "number",
				value: value,
			})
			continue
		}

		if isLetter(char) {
			value := ""
			for isLetter(char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})
			continue
		}

		break
	}

	return tokens
}

func main() {
	tokens := tokenizer("(add 10 (subtract 10 6))")

	fmt.Println(tokens)
}
