package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFields(s string) map[int]struct{} {
	fields := make(map[int]struct{})
	parts := strings.Split(s, ",")
	for _, p := range parts {
		if strings.Contains(p, "-") {
			bounds := strings.SplitN(p, "-", 2)
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			for i := start; i <= end; i++ {
				fields[i] = struct{}{}
			}
		} else {
			num, _ := strconv.Atoi(p)
			fields[num] = struct{}{}
		}
	}
	return fields
}

func main() {
	fieldsFlag := flag.String("f", "", "specifying field numbers (example: 1,3-5)")
	delimiter := flag.String("d", "\t", "separator (tab by default)")
	separated := flag.Bool("s", false, "only strings with separator")

	flag.Parse()

	if *fieldsFlag == "" {
		fmt.Fprintln(os.Stderr, "ERROR: must specify -f")
		os.Exit(1)
	}

	fields := parseFields(*fieldsFlag)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		parts := strings.Split(line, *delimiter)

		var out []string
		for i := 1; i <= len(parts); i++ {
			if _, ok := fields[i]; ok {
				out = append(out, parts[i-1])
			}
		}

		if len(out) > 0 {
			fmt.Println(strings.Join(out, *delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		os.Exit(1)
	}
}
