package main

import (
	"fmt"
)

func main() {

	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2

	canPlaceFlowers(flowerbed, n)
}

func mergeAlternately(word1 string, word2 string) string {

	len1 := len(word1)
	len2 := len(word2)
	merge := []byte{}
	var len int
	if len1 > len2 {
		len = len1
	} else {
		len = len2
	}
	for i := 0; i < len; i++ {
		if i < len1 {
			if word1[i] != 0 {
				merge = append(merge, word1[i])
			}
		}
		if i < len2 {
			if word2[i] != 0 {
				merge = append(merge, word2[i])
			}
		}
	}
	return string(merge)
}

func gcdOfStrings(word1 string, word2 string) string {

	if len(word1) < len(word2) {
		gcdOfStrings(word2, word1)
	}
	if word1+word2 != word2+word1 {
		return ""
	}
	a, b := len(word1), len(word2)
	for b != 0 {
		a, b = b, a%b
		fmt.Printf("==> a %v b %v \n", a, b)
	}
	return string(word1[:a])

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	Kids_with_largest_candies := []bool{}
	var greatest int
	greatest = candies[0]
	for i := range candies {
		if greatest < candies[i] {
			greatest = candies[i]
		}
	}

	for j := range candies {
		if candies[j]+extraCandies < greatest {
			Kids_with_largest_candies = append(Kids_with_largest_candies, false)
		} else {
			Kids_with_largest_candies = append(Kids_with_largest_candies, true)
		}
	}
	return Kids_with_largest_candies
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	var plantCount int
	if len(flowerbed) == 1 {
		if n > 0 {
			if flowerbed[0] == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}

	}
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		plantCount += 1
	}
	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		flowerbed[len(flowerbed)-1] = 1
		plantCount += 1
	}
	for i := 1; i < len(flowerbed)-1; i++ {

		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			plantCount += 1
		}

	}

	if plantCount >= n {
		return true
	} else {
		return false
	}
}
