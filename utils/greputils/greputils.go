package greputils

import (
	"fmt"
	"regexp"
	"strings"
)

type Options struct {
	After       int
	Before      int
	Count       bool
	Insensitive bool
	Number      bool
	Invert      bool
	Fixed       bool
}

func compilePattern(pattern string, opts Options) (*regexp.Regexp, string, error) {
	if opts.Fixed {
		if opts.Insensitive {
			pattern = strings.ToLower(pattern)
		}
		return nil, pattern, nil
	}

	if opts.Insensitive {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	return re, pattern, err
}

func matchLine(line string, re *regexp.Regexp, fixedPattern string, opts Options) bool {
	var ok bool
	if opts.Fixed {
		target := line
		if opts.Insensitive {
			target = strings.ToLower(line)
		}
		ok = strings.Contains(target, fixedPattern)
	} else {
		ok = re.MatchString(line)
	}

	if opts.Invert {
		ok = !ok
	}
	return ok
}

func printMatches(lines []string, matches []bool, opts Options) {
	if opts.Count {
		count := 0
		for _, m := range matches {
			if m {
				count++
			}
		}
		fmt.Println(count)
		return
	}

	printed := make(map[int]bool)
	for i, m := range matches {
		if m {
			start := i - opts.Before
			if start < 0 {
				start = 0
			}
			end := i + opts.After
			if end >= len(lines) {
				end = len(lines) - 1
			}
			for j := start; j <= end; j++ {
				if !printed[j] {
					if opts.Number {
						fmt.Printf("%d:%s\n", j+1, lines[j])
					} else {
						fmt.Println(lines[j])
					}
					printed[j] = true
				}
			}
		}
	}
}

// Run -
func Run(pattern string, lines []string, opts Options) {
	re, fixedPattern, err := compilePattern(pattern, opts)
	if err != nil {
		fmt.Println("invalid regexp:", err)
		return
	}

	matches := make([]bool, len(lines))
	for i, line := range lines {
		matches[i] = matchLine(line, re, fixedPattern, opts)
	}

	printMatches(lines, matches, opts)
}
