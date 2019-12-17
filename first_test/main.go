package main

import (
	"bufio"
	"fmt"
	"github.com/badoux/goscraper"
	"log"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//add func for read our map
func getStrings(lines []string) map[string]string {

	m := make(map[string]string)
	for ind, line := range lines {
		args := strings.Split(line, ",")
		if len(args) != 2 {
			fmt.Printf("error: invalid row %d {%s}\n", ind, line)
			continue
		}
		uStr := strings.TrimSpace(args[0])
		header := strings.TrimSpace(args[1])
		m[uStr] = header
	}
	return m
}

// add func for our request, and check title
func makeRequest(uStr, header string) {
	s, err := goscraper.Scrape(uStr, 5)
	if err != nil {
		fmt.Println(err)
	}

	if 	s.Preview.Title != header {
		fmt.Println(s.Preview.Title)
		fmt.Println(header)
	}

}

func main() {
	lines, err := readLines("foo.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	m := getStrings(lines)

	for u, header := range m {
		makeRequest(u, header)

	}

}
