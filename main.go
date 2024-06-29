package main

import (
	"fmt"
)

func main() {
	resReverseAlph := reverseAlphabet("EIGEN12345")
	fmt.Println("1.", resReverseAlph)

	sentence := "Saya sangat senang mengerjakan soal algoritma"
	resLongestChar := findLongestCharacter(sentence)
	fmt.Println("2.", resLongestChar)

	INPUT := []string{"xc", "dz", "bbb", "dz"}
	QUERY := []string{"bbb", "ac", "dz"}
	resDuplicate := CountDuplicateWord(INPUT, QUERY)
	fmt.Println("3.", resDuplicate)

	Matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
	resDiagonalMatrix := DiagonalMatrixSubstraction(Matrix)
	fmt.Println("4.", resDiagonalMatrix)
}

// 1. Terdapat string "NEGIE1", silahkan reverse alphabet nya dengan angka tetap diakhir kata Hasil = "EIGEN1"
func reverseAlphabet(input string) (res string) {
	length := len(input)
	var numb []string
	for i := length - 1; i >= 0; i-- {
		if input[i] >= '0' && input[i] <= '9' {
			numb = append(numb, string(input[i]))
			continue
		}
		res += string(input[i])
	}
	for i := len(numb) - 1; i >= 0; i-- {
		res += numb[i]
	}

	return
}

//  2. Diberikan contoh sebuah kalimat, silahkan cari kata terpanjang dari kalimat tersebut,
//     Jika ada kata dengan panjang yang sama silahkan ambil salah satu
func findLongestCharacter(input string) (res string) {
	longestWord := ""
	currentWord := ""

	// length := len(input)-1
	for i, v := range input {
		if string(input[i]) == " " {
			if len(currentWord) > len(longestWord) {
				longestWord = currentWord
			}
			currentWord = ""
			continue
		}
		currentWord += string(v)
	}

	res = fmt.Sprintf("%s: %d character", longestWord, len(longestWord))

	return
}

//  3. Terdapat dua buah array yaitu array INPUT dan array QUERY,
//     silahkan tentukan berapa kali kata dalam QUERY terdapat pada array INPUT
func CountDuplicateWord(INPUT, QUERY []string) (res []int) {
	queryMap := make(map[string]int)

	// put word in query array to queryMap
	for _, v := range QUERY {
		queryMap[string(v)] = 0
	}

	// count how many word in query array found in input array
	for _, v := range INPUT {
		if _, isFound := queryMap[string(v)]; isFound {
			queryMap[string(v)]++
		}
	}

	for _, v := range QUERY {
		res = append(res, queryMap[string(v)])
	}

	return
}

// 4. Silahkan cari hasil dari pengurangan dari jumlah diagonal sebuah matrik NxN
func DiagonalMatrixSubstraction(matrix [][]int) int {
	firstDiagonal := 0
	secondDiagonal := 0
	length := len(matrix)

	for i, val := range matrix {
		firstDiagonal += val[i]
		secondDiagonal += val[(length-1)-i]
	}
	fmt.Println("frst diagonal:", firstDiagonal)
	fmt.Println("second diagonal:", secondDiagonal)

	return firstDiagonal - secondDiagonal
}
