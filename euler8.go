package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calcProd(slice string) uint64 {
	var prod uint64 = 1
	for i := 0; i < len(slice); i++ {
		tmp, _ := strconv.Atoi(string(slice[i]))
		prod *= uint64(tmp)
	}
	return prod
}

func intervalProds(size int, slice string) uint64 {
	var best uint64 = 1
	for i := 0; i + size <= len(slice); i++ {
		tmp := calcProd(slice[i:i+size])
		if best < tmp {
			best = tmp
		}
	}
	return best
}

func main() {

	var big_string string = "73167176531330624919225119674426574742355349194934"
	big_string += "96983520312774506326239578318016984801869478851843"
	big_string += "85861560789112949495459501737958331952853208805511"
	big_string += "12540698747158523863050715693290963295227443043557"
	big_string += "66896648950445244523161731856403098711121722383113"
	big_string += "62229893423380308135336276614282806444486645238749"
	big_string += "30358907296290491560440772390713810515859307960866"
	big_string += "70172427121883998797908792274921901699720888093776"
	big_string += "65727333001053367881220235421809751254540594752243"
	big_string += "52584907711670556013604839586446706324415722155397"
	big_string += "53697817977846174064955149290862569321978468622482"
	big_string += "83972241375657056057490261407972968652414535100474"
	big_string += "82166370484403199890008895243450658541227588666881"
	big_string += "16427171479924442928230863465674813919123162824586"
	big_string += "17866458359124566529476545682848912883142607690042"
	big_string += "24219022671055626321111109370544217506941658960408"
	big_string += "07198403850962455444362981230987879927244284909188"
	big_string += "84580156166097919133875499200524063689912560717606"
	big_string += "05886116467109405077541002256983155200055935729725"
	big_string += "71636269561882670428252483600823257530420752963450"

	validStrings := strings.Split(big_string, "0")
	filteredStrings := make([]string, 1)
	size := 13
	var prod uint64 = 1

	fmt.Println("Set of split strings:")
	for _, v := range validStrings {
		if len(v) >= size {
			filteredStrings = append(filteredStrings, v)
		}
	}

	fmt.Println("Set of filtered split strings:")
	for _, v := range filteredStrings {
		fmt.Println(v)
	}

	fmt.Println("Greatest product in each string:")
	for _, v := range filteredStrings {
		tmp := intervalProds(size, v)
		fmt.Println(tmp)
		if tmp > prod {
			prod = tmp
		}
	}

	fmt.Println("Greatest", size, "consec digit product is", prod)
}