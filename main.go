package main

import (
	"fmt"
	"log"
	"strings"
)

/*
----------
Tokenizer
----------
*/

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

/*
----------
Parser
----------
*/

type node struct {
	callee     *node
	expression *node
	arguments  *[]node
	context    *[]node
	kind       string
	value      string
	name       string
	body       []node
	params     []node
}

type ast node // alias type

var (
	pc int
	pt []token
)

func parser(tokens []token) ast {
	pc = 0
	pt = tokens

	ast_node := ast{
		kind: "Program",
		body: []node{},
	}

	// each loop appens a call CallExpression or a NumberLiteral
	for pc < len(pt) {
		ast_node.body = append(ast_node.body, walk())
	}

	return ast_node
}

// recursively fills the node (This is so smart lol, I would have never thought of this)
func walk() node {
	token := pt[pc]

	if token.kind == "number" {
		pc++

		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}

	if token.kind == "paren" && token.value == "(" {
		pc++
		token = pt[pc]

		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}

		pc++ // increment to skip the name token
		token = pt[pc]

		// Loop through tokens which will be used for params
		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}

		pc++

		return n
	}

	// If havne't recognized the token type by now, throw an error
	log.Fatal(token.kind)
	return node{}
}

/*
-----------
Traverser
-----------
*/

type visitor map[string]func(n *node, p node)

func traverser(a ast, v visitor) {
	traverseNode(node(a), node{}, v)
}

func traverseNode(n node, p node, v visitor) {
	for k, method := range v {
		if n.kind == k {
			method(&n, p)
		}
	}

	switch n.kind {
	case "Program":
		traverseArray(n.body, n, v)
	case "CallExpression":
		traverseArray(n.params, n, v)
	case "NumberLiteral":
		break
	default:
		log.Fatal(n.kind)
	}
}

func traverseArray(nodes []node, parent node, v visitor) {
	for _, child := range nodes {
		traverseNode(child, parent, v)
	}
}

/*
-----------
Transformer
-----------
*/

func transformer(a ast) ast {
	nast := ast{
		kind: "Program",
		body: []node{},
	}

	a.context = &nast.body

	traverser(a, map[string]func(n *node, p node){
		"NumberLiteral": func(n *node, p node) {
			*p.context = append(*p.context, node{
				kind:  "NumberLiteral",
				value: n.value,
			})
		},
		"CallExpression": func(n *node, p node) {
			e := node{
				kind: "CallExpression",
				callee: &node{
					kind: "Identifier",
					name: n.name,
				},
				arguments: new([]node), // This line caused me 2 hours of debugging a nil pointer....
			}

			n.context = e.arguments

			if p.kind != "CallExpression" {
				es := node{
					kind:       "ExpressionStatement",
					expression: &e,
				}

				*p.context = append(*p.context, es)
			} else {
				*p.context = append(*p.context, e)
			}
		},
	})

	return nast
}

/*
-----------
Code Generator
-----------
*/

func codeGenerator(n node) string {
	switch n.kind {
	case "Program":
		var result []string
		for _, ast_node := range n.body {
			result = append(result, codeGenerator(ast_node))
		}
		return strings.Join(result, "\n")
	case "ExpressionStatement":
		return codeGenerator(*n.expression) + ";"
	case "CallExpression":
		var ra []string
		c := codeGenerator(*n.callee)

		for _, ast_node := range *n.arguments {
			ra = append(ra, codeGenerator(ast_node))
		}

		r := strings.Join(ra, ", ")
		return c + "(" + r + ")"

	case "Identifier":
		return n.name
	case "NumberLiteral":
		return n.value
	default:
		log.Fatal("err")
		return ""
	}
}

/*
-----------
Compiler
-----------
*/

func compiler(input string) string {
	tokens := tokenizer(input)
	ast_node := parser(tokens)
	nast := transformer(ast_node)
	output := codeGenerator(node(nast))

	return output
}

func main() {
	program := "(add 32 (subtract 69 420))"
	output := compiler(program)

	fmt.Println(output)
}
