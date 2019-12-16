package second_test

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func getStrings(lines []string) map[string]string {

	m := make(map[string]string)
	for ind, line := range lines {
		args := strings.Split(line, ",")
		if len(args) != 2 {
			fmt.Printf("error: invalid row %d {%s}\n", ind, line)
			continue
		}
		uStr := strings.TrimSpace(args[0])
		code := args[1]
		m[uStr] = code
	}

	return m
}

func MakeRequest(uStr string, code int) {
	resp, err := http.Get(uStr)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != code {
		fmt.Printf("Ошибка. http-статус: %s\n", resp.StatusCode)
		return
	}

}

func secondTest() {
	lines, err := readLines("foo.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	m := getStrings(lines)

	for u, code := range m {
		MakeRequest(u, code)

	}

}
