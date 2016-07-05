// counting sundays
package main

import (
	"fmt"
)

func firstSundaysInYear(year int, start int) int {

}

func daysInYear(year int) int {
	base := 365
	if isLeap(year) {
		base++
	}
	return base
}

func isLeap(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {

}