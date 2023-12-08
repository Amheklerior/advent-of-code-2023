package lib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not open the file: %s", err)
	}
	content := string(bytes)
	return content
}

func Scanner(fileContent string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(fileContent))
}
