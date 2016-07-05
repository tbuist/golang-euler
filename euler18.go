// maximum path sum I
package main

import (
	"fmt"
	"strings"
	"strconv"
)

// make a *struct instead? 8 byte ptr vs 48 bytes
func recurBest(nums *[][]int, track *[]int, curSum int, bestSum *int, rowIdx int, colIdx int) {

	if rowIdx > 13 {
		if curSum > *bestSum {
			fmt.Println("row:",rowIdx,"col:",colIdx,"sum:",curSum,"best:",*bestSum,"order:",*track)
			*bestSum = curSum
		}
		return
	}

	curSum += (*nums)[rowIdx+1][colIdx]
	*track = append(*track, (*nums)[rowIdx+1][colIdx])

	recurBest(nums, track, curSum, bestSum, rowIdx+1, colIdx)

	curSum -= (*nums)[rowIdx+1][colIdx]
	curSum += (*nums)[rowIdx+1][colIdx+1]
	*track = append((*track)[:len(*track)-1], (*nums)[rowIdx+1][colIdx+1])

	recurBest(nums, track, curSum, bestSum, rowIdx+1, colIdx+1)

	curSum -= (*nums)[rowIdx+1][colIdx+1]
	*track = (*track)[:len(*track)-1]

	return
}

func main() {

	numstr := "75 95 64 17 47 82 18 35 87 10 20 04 82 47 65 19 01 23 75 03 34 88 02 77 73 07 63 67 99 65 04 28 06 16 70 92 41 41 26 56 83 40 80 70 33 41 48 72 33 47 32 37 16 94 29 53 71 44 65 25 43 91 52 97 51 14 70 11 33 28 77 73 17 78 39 68 17 57 91 71 52 38 17 14 91 43 58 50 27 29 48 63 66 04 68 89 53 67 30 73 16 69 87 40 31 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"
	numstrarr := strings.Split(numstr, " ")
	var tmpnumarr []int
	for _, v := range numstrarr {
		tmp, _ := strconv.Atoi(v)
		tmpnumarr = append(tmpnumarr, tmp)
	}

	// important one
	numsarr := make([][]int, 15)
	for i, _ := range numsarr {
		numsarr[i] = make([]int, 0)
	}

	left := 0
	right := 1

	for i := 0; i < 15; i++ {
		numsarr[i] = append(numsarr[i], tmpnumarr[left:right]...)
		tmp := left
		left = right
		right = right + (right - tmp) + 1
	}

	numsarr[0][0] = 75

	best := 0
	bestSum := &best
	numsptr := &numsarr
	trackNums := make([]int, 0)
	track := &trackNums

	recurBest(numsptr, track, 75, bestSum, 0, 0)


	fmt.Println("Maximum path sum:", *bestSum)
}