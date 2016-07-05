// number letter counts
package main

import (
	"fmt"
	"math"
)

func numLetters(x int) int {

	if x == 1000 {
		return 11
	}

	sum := 0

	hundreds := int(math.Floor(float64(x) / 100))
	tens := x % 100
	tens = int(math.Floor(float64(tens) / 10))
	ones := x % 10

	// non-zero hunds: "hundred and"
	basehund := 10

	switch {
		case hundreds == 0:
		
		case hundreds == 1 || hundreds == 2 || hundreds == 6:
			sum += len("one") + basehund
		
		case hundreds == 3 || hundreds == 7 || hundreds == 8:
			sum += len("three") + basehund 
		
		case hundreds == 4 || hundreds == 5 || hundreds == 9:
			sum += len("four") + basehund	
	}

	if tens == 0 && ones == 0 {
		sum -= 3
	}

	switch {
		case tens == 0:
		
		case tens == 2 || tens == 3 || tens == 8 || tens == 9:
			sum += len("twenty")
		
		case tens == 4 || tens == 5 || tens == 6:
			sum += len("forty")
		
		case tens == 7:
			sum += len("seventy")
		
		case tens == 1:
			switch {
				case ones == 0:
					sum += len("ten")
				
				case ones == 1:
					sum += len("eleven")
				
				case ones == 2:
					sum += len("twelve")
				
				case ones == 3:
					sum += len("thirteen")
				
				case ones == 4:
					sum += len("fourteen")
				
				case ones == 5:
					sum += len("fifteen")
				
				case ones == 6:
					sum += len("sixteen")
				
				case ones == 7:
					sum += len("seventeen")
				
				case ones == 8:
					sum += len("eighteen")
				
				case ones == 9:
					sum += len("nineteen")
			}
			return sum
		}

	switch {
		case ones == 0:
		
		case ones == 1:
			sum += len("one")
		
		case ones == 2:
			sum += len("two")
		
		case ones == 3:
			sum += len("three")
		
		case ones == 4:
			sum += len("four")
		
		case ones == 5:
			sum += len("five")
		
		case ones == 6:
			sum += len("six")
		
		case ones == 7:
			sum += len("seven")
		
		case ones == 8:
			sum += len("eight")
		
		case ones == 9:
			sum += len("nine")
		
	}
	return sum
}

func main() {
	sum := 0

	for i := 1; i < 1001; i++ {
		sum += numLetters(i)
	}

	fmt.Println("Number of letters:", sum)
	
}