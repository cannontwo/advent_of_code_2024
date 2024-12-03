package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

type token int

const (
	ILLEGAL token = iota
	EOF

	MUL
	DO
	DONT

	OPEN_PAREN
	CLOSE_PAREN
	COMMA

	NUMBER
)

var eof = rune(0)

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func scanNumber(reader *bufio.Reader) (token, string) {
	var buf bytes.Buffer

	for range 3 {
		r, _, err := reader.ReadRune()
		if err != nil {
			return EOF, string(eof)
		}
		if isDigit(r) {
			buf.WriteRune(r)
		} else {
			reader.UnreadRune()
			break
		}
	}

	return NUMBER, buf.String()
}

func scanMul(reader *bufio.Reader) (token, string) {
	var buf bytes.Buffer

	r, _, err := reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 'm' {
		return ILLEGAL, buf.String()
	}

	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 'u' {
		return ILLEGAL, buf.String()
	}

	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 'l' {
		return ILLEGAL, buf.String()
	}

	return MUL, buf.String()
}

func scanDo(reader *bufio.Reader) (token, string) {
	var buf bytes.Buffer

	r, _, err := reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 'd' {
		return ILLEGAL, buf.String()
	}

	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 'o' {
		return ILLEGAL, buf.String()
	}

	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	if r != 'n' {
		reader.UnreadRune()
		return DO, buf.String()
	} else {
		buf.WriteRune(r)
	}

	// Committed to processing DONT token
	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != '\'' {
		return ILLEGAL, buf.String()
	}

	r, _, err = reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}
	buf.WriteRune(r)
	if r != 't' {
		return ILLEGAL, buf.String()
	}

	return DONT, buf.String()
}

func scan(reader *bufio.Reader) (token, string) {
	r, _, err := reader.ReadRune()
	if err != nil {
		return EOF, string(eof)
	}

	if isDigit(r) {
		reader.UnreadRune()
		return scanNumber(reader)
	}

	switch r {
	case 'd':
		reader.UnreadRune()
		return scanDo(reader)
	case 'm':
		reader.UnreadRune()
		return scanMul(reader)
	case '(':
		return OPEN_PAREN, string(r)
	case ')':
		return CLOSE_PAREN, string(r)
	case ',':
		return COMMA, string(r)
	case eof:
		return EOF, string(eof)
	default:
		return ILLEGAL, string(r)
	}
}

func parseMul(tokens []token, literals []string) (int, int, bool) {
	var left, right int
	var err error

	if tokens[0] != MUL {
		return 0, 0, false
	}

	if tokens[1] != OPEN_PAREN {
		return 0, 0, false
	}

	if tokens[2] != NUMBER {
		return 0, 0, false
	} else {
		left, err = strconv.Atoi(literals[2])
		if err != nil {
			panic(err)
		}
	}

	if tokens[3] != COMMA {
		return 0, 0, false
	}

	if tokens[4] != NUMBER {
		return 0, 0, false
	} else {
		right, err = strconv.Atoi(literals[4])
		if err != nil {
			panic(err)
		}
	}

	if tokens[5] != CLOSE_PAREN {
		return 0, 0, false
	}

	return left, right, true
}

func parseDo(tokens []token) bool {
	if tokens[0] != DO {
		return false
	}

	if tokens[1] != OPEN_PAREN {
		return false
	}

	if tokens[2] != CLOSE_PAREN {
		return false
	}

	return true
}

func parseDont(tokens []token) bool {
	if tokens[0] != DONT {
		return false
	}

	if tokens[1] != OPEN_PAREN {
		return false
	}

	if tokens[2] != CLOSE_PAREN {
		return false
	}

	return true

}

func run_day_three() {
	if len(os.Args) < 3 {
		log.Fatalf("Please provide a filename to read as input for day 1")
	}

	filename := os.Args[2]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var tokens []token
	var literals []string
	var token token
	var lit string
	for token != EOF {
		token, lit = scan(reader)
		if token != ILLEGAL {
			tokens = append(tokens, token)
			literals = append(literals, lit)
		}
	}

	enabled := true
	var productSum, dontProductSum int
	var i int
	for i < len(tokens)-5 {
		if parseDo(tokens[i : i+4]) {
			enabled = true
			i += 3
			continue
		}

		if parseDont(tokens[i : i+4]) {
			enabled = false
			i += 3
			continue
		}

		left, right, ok := parseMul(tokens[i:i+6], literals[i:i+6])
		if ok {
			productSum += left * right
			if enabled {
				dontProductSum += left * right
			}

			i += 5
			continue
		}

		i += 1
	}

	fmt.Printf("Sum of well-formed products: %v\n", productSum)
	fmt.Printf("Sum of well-formed products, using do/don't: %v\n", dontProductSum)
}
