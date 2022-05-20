package main

import (
	"asciiartreverse"
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return file
}

func sliceFile(file *os.File) []string {
	var slice []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return slice
}

func checkAscii(mapslice map[int][]string, fileart []string) string {
	checkascii := make([]string, 8)
	runewidth := []rune(fileart[0])
	output := ""
	// loop through each index of the first slice (first line) of the slices of the inpput file
	for i := 0; i < len(runewidth); i++ {
		// loop through each slice, which is each new line of the file
		for j, line := range fileart {
			// add in the art by vertical line
			slice := []rune(line)
			checkascii[j] += string(slice[i])
		}
		for k, art := range mapslice {
			count := 0
			match := true
			for num, slice := range art {
				if num != 0 && slice == checkascii[count] {
					count++
				} else if num != 0 {
					match = false
					break
				}
			}
			if match {
				output += string(rune(k))
				checkascii = make([]string, 8)
				break
			}
		}

	}
	return output
}

func main() {
	// ASCII ART file open
	file := openFile("standard.txt")

	// ASCII ART file to slice of string by line
	lttrlines := sliceFile(file)

	if len(os.Args) != 2  || len(os.Args[1]) < 10 || os.Args[1][0:10] != "--reverse=" {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println("EX: go run . --reverse=<fileName>")
	} else {
		// INPUT FILE open
		reverse := openFile(os.Args[1][10:])

		// INPUT file to slice of string by line
		fileart := sliceFile(reverse)

		// create map of ascii art file
		mapslice := asciiartreverse.Createmap(lttrlines)

		// compare input file with map of ascii art and collect index(ascii number) to ceate string
		output := checkAscii(mapslice, fileart)

		// print to terminal
		fmt.Println(output)

	}

	file.Close()
}
