package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-utility/utils/greputils"
	"os"
)

func readLines(files []string) ([]string, error) {
	var lines []string

	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines, scanner.Err()
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return lines, nil
}

func main() {
	after := flag.Int("A", 0, "print N lines After match")
	before := flag.Int("B", 0, "print N lines Before match")
	circle := flag.Int("C", 0, "print N lines of Context (before and after)")

	count := flag.Bool("c", false, "can find the number of lines that matches the given string")
	insensitive := flag.Bool("i", false, "option enables to search for a string case insensitively in the given file")
	numeric := flag.Bool("n", false, "show the line number of file with the line matched. ")
	inverting := flag.Bool("v", false, "can display the lines that are not matched with the specified search string ")
	fixed := flag.Bool("F", false, "treats the template as a fixed string, not a regular expression.")

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "go-grep [OPTIONS] PATTERN [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	lines, err := readLines(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR reading input:", err)
		os.Exit(1)
	}

	if *circle > 0 {
		*after = *circle
		*before = *circle
	}

	opts := greputils.Options{
		After:       *after,
		Before:      *before,
		Count:       *count,
		Insensitive: *insensitive,
		Number:      *numeric,
		Invert:      *inverting,
		Fixed:       *fixed,
	}

	greputils.Run(pattern, lines, opts)

}
