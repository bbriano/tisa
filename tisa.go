// Tisa replace tab characters (ab)used for alignment with spaces.
// It only replace tab characters after the first non-tab character.
// Or in other words, Tisa preserves tab characters used for indentation.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Usage: tisa [ nspace ]")
		os.Exit(1)
	}

	nSpace := 8
	if len(os.Args) == 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v", err)
			os.Exit(1)
		}
		nSpace = i
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		i := strings.IndexFunc(line, func(r rune) bool { return r != '\t' })
		if i == -1 {
			fmt.Println(line)
			continue
		}

		out := line[:i]
		for _, r := range line[i:] {
			switch r {
			case '\t':
				w := len(out) + strings.Count(out, "\t")*(nSpace-1)
				out += strings.Repeat(" ", nSpace-w%nSpace)
			default:
				out += string(r)
			}
		}
		fmt.Println(out)
	}
}
