package sortutils

import (
	"strconv"
	"strings"
)

// Comparator возвращает функцию-компаратор для сортировки строк.
// Функция возвращает компаратор в виде func(i, j int) bool, который можно использовать
// в sort.Slice() или аналогичных функциях.
func Comparator(lines []string, col int, numeric, reverse, month, human bool) func(i, j int) bool {
	return func(i, j int) bool {
		a := GetColumnValue(lines[i], col)
		b := GetColumnValue(lines[j], col)

		var less bool
		switch {
		case month:
			less = MonthOrder[a] < MonthOrder[b]
		case human:
			valA, errA := ParseHumanReadable(a)
			valB, errB := ParseHumanReadable(b)
			if errA == nil && errB == nil {
				less = valA < valB
			} else {
				less = a < b
			}
		case numeric:
			numA, errA := strconv.ParseFloat(a, 64)
			numB, errB := strconv.ParseFloat(b, 64)
			if errA == nil && errB == nil {
				less = numA < numB
			} else {
				less = a < b
			}
		default:
			less = a < b
		}

		if reverse {
			return !less
		}
		return less
	}
}

// MonthOrder используется для -M
var MonthOrder = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

// ParseHumanReadable парсит размеры вида 10K, 5M
func ParseHumanReadable(s string) (float64, error) {
	if len(s) == 0 {
		return 0, nil
	}
	multipliers := map[byte]float64{
		'K': 1024,
		'M': 1024 * 1024,
		'G': 1024 * 1024 * 1024,
	}

	last := s[len(s)-1]
	if mult, ok := multipliers[last]; ok {
		num, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			return 0, err
		}
		return num * mult, nil
	}
	return strconv.ParseFloat(s, 64)
}

// GetColumnValue возвращает значение из нужной колонки
func GetColumnValue(line string, col int) string {
	if col < 0 {
		return line
	}
	cols := strings.Fields(line)
	if col >= 0 && col < len(cols) {
		return cols[col]
	}
	return ""
}

// RemoveDuplicates удаляет дубликаты
func RemoveDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}

// IsSorted проверяет отсортирован ли слайс
func IsSorted(lines []string, cmp func(i, j int) bool) bool {
	for i := 1; i < len(lines); i++ {
		if cmp(i, i-1) {
			return false
		}
	}
	return true
}
