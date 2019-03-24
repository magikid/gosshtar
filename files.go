package main

import (
	"bufio"
	"os"
)

func parseFile(path string) []string {
	_, err := os.Stat(path)
	handleError(err)

	return readLines(path)
}

func readLines(path string) []string {
	file, err := os.Open(path)
	handleError(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	handleError(scanner.Err())

	return lines
}
