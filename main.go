package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-sort-utility/utils/sortutils"
	"os"
	"sort"
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
	column := flag.Int("k", 0, "column number to sort by (1-indexed)")
	reverse := flag.Bool("r", false, "reverse sort order")
	numeric := flag.Bool("n", false, "numeric sort")
	unique := flag.Bool("u", false, "unique lines only")
	month := flag.Bool("M", false, "sort by month name")
	check := flag.Bool("c", false, "check if data is sorted")
	human := flag.Bool("h", false, "human numeric sort (e.g., 10K, 5M)")

	flag.Parse()
	files := flag.Args()

	lines, err := readLines(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR reading input:", err)
		os.Exit(1)
	}

	if *unique {
		lines = sortutils.RemoveDuplicates(lines)
	}

	colIndex := -1
	if *column > 0 {
		colIndex = *column - 1
	}

	cmp := sortutils.Comparator(lines, colIndex, *numeric, *reverse, *month, *human)

	if *check {
		if sortutils.IsSorted(lines, cmp) {
			fmt.Println("Input is sorted")
			os.Exit(0)
		} else {
			fmt.Println("Input is not sorted")
			os.Exit(1)
		}
	}

	sort.Slice(lines, cmp)

	for _, line := range lines {
		fmt.Println(line)
	}
}
