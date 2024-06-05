package main

import (
	"bufio"
	"os"
	"slices"
)

func main() {
	seen := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if slices.Index(seen, text) == -1 {
			println(text)
			seen = append(seen, text)
		}
	}
}
