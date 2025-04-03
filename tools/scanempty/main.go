package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isTriviallyEmpty(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		lines = append(lines, line)
	}

	if len(lines) == 1 && strings.HasPrefix(lines[0], "package ") {
		return true
	}
	if len(lines) == 2 &&
		strings.HasPrefix(lines[0], "package ") &&
		lines[1] == "func init() {}" {
		return true
	}

	return false
}

func main() {
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" && isTriviallyEmpty(path) {
			fmt.Println("🧹 Tom Go-fil:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Feil:", err)
	}
}
