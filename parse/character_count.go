// Package parse provides functions to count letters, characters, numbers and symbols in a string.
package parse

 import (
    "fmt"
    "bufio"
    "os"
    "strings"
	"unicode/utf8"
)

func LetterCount(s string) (map[rune]int, int) {

		charCount := utf8.RuneCountInString(s)
		fmt.Printf("Character count: %d\n", charCount)

		byteCount := len(s)
		fmt.Printf("Byte count: %d\n", byteCount)

		counts := make(map[rune]int)
		for _, ch := range s {
			counts[ch]++
		}

		w := bufio.NewWriter(os.Stdout)

		for ch := 'a'; ch <= 'z'; ch++ {
        fmt.Fprintf(w, "%c: %d\n", ch, counts[ch])
    }
		return counts, len(s)
}

func PrintCharacterCount(counts map[rune]int, w *bufio.Writer) {
		for ch := 'a'; ch <= 'z'; ch++ {
				fmt.Fprintf(w, "%c: %d\n", ch, counts[ch])
		}
}

func ReadInput(reader *bufio.Reader) (string, error) {
		s, err := reader.ReadString('\n')
		if err != nil {
				return "", err
		}
		s = strings.TrimSpace(s)
		return s, nil
}

func RuneCount(s string) int {
		return utf8.RuneCountInString(s)
}

