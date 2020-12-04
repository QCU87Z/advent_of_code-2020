package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// File format
// 	min-max letter: password

func getPasswords(path string) []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open")
	}
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		lineStr := scanner.Text()
		text = append(text, lineStr)
	}
	file.Close()
	return text
}

func main() {
	pwd := getPasswords("input")
	for _, pass := range pwd {
		a := regexp.MustCompile(`(\d*)`)
		// fmt.Printf("%s\t", pass)
		minAndMax := a.FindAllString(pass, 2)
		fmt.Println(minAndMax)

		b := regexp.MustCompile(`([a-zA-Z])`)
		letter := b.FindAllString(pass, 2)
		fmt.Println(letter)
	}
}
