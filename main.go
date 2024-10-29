package main

import "fmt"

func main() {
	word1 := "ab"
	word2 := "pqrs"

	fmt.Printf("===> Check %v \n", MergeAlternately(word1, word2))
}

func MergeAlternately(word1 string, word2 string) string {

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
