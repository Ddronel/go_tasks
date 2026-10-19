package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	values := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	if err := printTypes(os.Stdout, values); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	controlStr := joinValues(values)
	fmt.Println(controlStr)

	sliceRune := []rune(controlStr)
	fmt.Println(hashRunes(sliceRune, "go-2024"))
}

func printTypes(w io.Writer, values []any) error {
	for _, value := range values {
		if _, err := fmt.Fprintf(w, "%T\n", value); err != nil {
			return err
		}
	}
	return nil
}

func joinValues(values []any) string {
	var builder strings.Builder
	for _, value := range values {
		builder.WriteString(fmt.Sprint(value))
	}
	return builder.String()
}

func hashRunes(runes []rune, salt string) string {
	middle := len(runes) / 2
	saltedStr := string(runes[:middle]) + salt + string(runes[middle:])

	hash := sha256.Sum256([]byte(saltedStr))
	return fmt.Sprintf("%x", hash)
}
