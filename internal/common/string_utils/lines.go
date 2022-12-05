package stringUtils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetLine(input string) chan string {
	lines := make(chan (string))

	go func() {
		for _, line := range strings.Split(input, "\n") {
			lines <- line
		}
		close(lines)
	}()

	return lines
}

func GetLineAsString(filename string) (chan string, chan error) {
	lines := make(chan (string))
	errors := make(chan (error))

	go func() {
		file, err := os.Open(filename)
		defer file.Close()

		if err != nil {
			close(lines)
			close(errors)
			log.Fatal(err)
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
			errors <- scanner.Err()
		}

		close(lines)
		close(errors)
	}()

	return lines, errors
}
