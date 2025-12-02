package utils

import (
	"bufio"
	"io"
	"strings"
)

// InputAsLines reads an io.Reader and breaks it up into a slice of strings
// where each element is a line in the input, with the trailing newline removed.
// If preserveWhitespace is false, strings.TrimSpace is called on the line before
// adding it to the slice.
func InputAsLines(in io.Reader, preserveWhitespace bool) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if !preserveWhitespace {
			line = strings.TrimSpace(line)
		}

		lines = append(lines, line)
	}

	return lines, scanner.Err()
}
