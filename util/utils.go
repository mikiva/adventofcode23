package util

import (
	"fmt"
	"os"
	"strings"
)

func GetInputLines() ([]string, error) {
	if len(os.Args) < 2 {
		return []string{}, fmt.Errorf("provide filename")
	}
	filename := fmt.Sprintf("%s.txt", os.Args[1])
	content, _ := os.ReadFile(filename)
	str := string(content)
	lines := strings.Split(str, "\n")
	return lines, nil
}

func SplitInTwo(str, separator string) (string, string) {

	split := strings.Split(str, separator)
	newSplit := make([]string, 0)
	for _, s := range split {
		if len(strings.TrimSpace(s)) > 0 {
			fmt.Println("S", s)
			newSplit = append(newSplit, s)
		}
	}

	return strings.TrimSpace(newSplit[0]), strings.TrimSpace(newSplit[1])
}
