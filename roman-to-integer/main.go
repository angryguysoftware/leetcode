package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	var sum = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' && len(s) != i+1 {
			if s[i+1] == 'V' {
				sum += 4
				i++
				continue
			} else if s[i+1] == 'X' {
				sum += 9
				i++
				continue
			} else {
				sum++
			}
		} else if s[i] == 'X' && len(s) != i+1 {
			if s[i+1] == 'L' {
				sum += 40
				i++
				continue
			} else if s[i+1] == 'C' {
				sum += 90
				i++
				continue
			} else {
				sum += 10
			}
		} else if s[i] == 'C' && len(s) != i+1 {
			if s[i+1] == 'D' {
				sum += 400
				i++
				continue
			} else if s[i+1] == 'M' {
				sum += 900
				i++
				continue
			} else {
				sum += 100
			}
		} else if s[i] == 'I' {
			sum += 1
		} else if s[i] == 'V' {
			sum += 5
		} else if s[i] == 'X' {
			sum += 10
		} else if s[i] == 'L' {
			sum += 50
		} else if s[i] == 'C' {
			sum += 100
		} else if s[i] == 'D' {
			sum += 500
		} else if s[i] == 'M' {
			sum += 1000
		}
	}
	return sum
}
