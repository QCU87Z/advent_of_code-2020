package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInts(path string) []int {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open")
	}
	scanner := bufio.NewScanner(file)
	var text []int
	for scanner.Scan() {
		lineStr := scanner.Text()
		value, _ := strconv.Atoi(lineStr)
		text = append(text, value)
	}
	file.Close()
	return text
}

func main() {
	ints := getInts("intput")
	for _, j := range ints {
		for _, k := range ints {
			if j+k == 2020 {
				fmt.Printf("%d + %d = 2020\n", j, k)
				fmt.Printf("Which means %d * %d = %d\n", j, k, j*k)
			}
		}
	}
}
