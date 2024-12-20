package main

import (
	"fmt"
)

func main() {
	given := []int{1, 2, 3, 4}
	correct := []int{24, 12, 8, 6}
	mine := productExceptSelf(given)
	fmt.Printf("==> given (%v)\n", given)
	fmt.Printf("==> mine (%v)\n", mine)
	fmt.Printf("==> correct (%v)\n", correct)
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

func reverseVowels(s string) string {

	sorted := []byte{}

	vowels := []byte{}
	for i := range s {
		if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" || string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			vowels = append(vowels, s[i])
		}
		sorted = append(sorted, s[i])

	}
	for i := range sorted {
		if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" || string(s[i]) == "U" || string(sorted[i]) == "a" || string(sorted[i]) == "e" || string(sorted[i]) == "i" || string(sorted[i]) == "o" || string(sorted[i]) == "u" {
			sorted[i] = vowels[len(vowels)-1]
			vowels = vowels[:len(vowels)-1]
		}
	}

	return string(sorted)
}

func reverseWords(s string) string {
	var reversed []byte
	word := []byte{}
	word = nil
	for i := range s {
		if string(s[i]) != " " {
			word = append(word, s[i])
		}
		if string(s[i]) == " " && word != nil {
			reversed = []byte(string(word) + " " + string(reversed))
			word = nil
		}
		if i == len(s)-1 && word != nil {
			reversed = []byte(string(word) + " " + string(reversed))
			word = nil
		}

	}
	return string(reversed[:len(reversed)-1])

}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		result[i] = right[i] * left[i]
	}
	return result

}
