package playground

import (
	"strconv"
	"strings"
)

// Write a program that automatically converts English text to Morse code and vice versa.
var letter = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '?'}
var code = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..", ".----", "..---", "...--", "....-", ".....", "-....", "--...", "---..", "----.", "-----", "..--.."}

func convertToMorseCode(input string) string {
	var output string = ""
	morseMap := make(map[byte]string)
	input = strings.ToLower(input)
	for i := 0; i < len(letter); i++ {
		morseMap[letter[i]] = code[i]
	}

	for i := 0; i < len(input); i++ {
		switch input[i : i+1] {
		case " ":
			output += "/ "
		default:
			output += morseMap[input[i]] + " "
		}
	}
	output = strings.TrimRight(output, " ")
	// fmt.Println("Output: ", output)
	return output
}

func convertFromMorseCode(input string) string {
	var output string = ""
	morseMap := make(map[string]byte)
	for i := 0; i < len(code); i++ {
		morseMap[code[i]] = letter[i]
	}

	for i := 0; i < len(input); i++ {
		switch input[i : i+1] {
		case "/":
			output += " "
			i++
		default:
			// find next space for end of code
			remainingStr := input[i : len(input)-1]
			nextSpace := strings.Index(remainingStr, " ")
			character := ""
			if nextSpace == -1 {
				character = input[i:]
				i = len(input)
			} else {
				character = input[i:(i + nextSpace)]
				i += nextSpace
			}
			output += string(morseMap[character])
			// fmt.Println("Word: ", output)
		}
	}

	return output
}

// Write a function that computes the list of the first 100 Fibonacci numbers. The first two
// Fibonacci numbers are 1 and 1. The n+1-st Fibonacci number can be computed by adding the n-th
// and the n-1-th Fibonacci number. The first few are therefore 1, 1, 1+1=2, 1+2=3, 2+3=5, 3+5=8.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765,
// 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811
func getFibonacciNumbers(limit int, numbers []int) []int {
	if limit == 0 {
		return []int{0}
	}

	var current = numbers[len(numbers)-1]
	var newFib int = 1
	if current > 0 {
		newFib = numbers[len(numbers)-1] + numbers[len(numbers)-2]
		if newFib > limit {
			return numbers
		}
	}
	return getFibonacciNumbers(limit, append(numbers, newFib))
}
func getFibonacciNumbersUpToLimit(limit int) []int {
	var output []int = getFibonacciNumbers(limit, []int{0})
	//fmt.Println("Fib: ", output)
	return output
}

// Write a function that takes a number and returns a list of its digits.
// So for 2342 it should return [2,3,4,2].
func convertDigitToList(digit int) []int {
	var digitStr string = strconv.Itoa(digit)
	var tmpA = make([]string, len(digitStr))
	var output []int
	for i := 0; i < len(digitStr); i++ {
		tmpA[i] = digitStr[i : i+1]
		if tmpI, err := strconv.Atoi(tmpA[i]); err == nil {
			output = append(output, tmpI)
		}
	}
	return output
}

// Write a function that combines two lists by alternatingly taking elements,
// e.g. [a,b,c], [1,2,3] → [a,1,b,2,c,3]
func combineLists(list1 []string, list2 []string) []string {
	if len(list1) <= 0 {
		return list2
	}
	if len(list2) <= 0 {
		return list1
	}
	var listSize = len(list1) + len(list2)
	var result = make([]string, listSize)
	var counter1, counter2, resultCounter int = 0, 0, 0
	for i := 0; i < listSize; i++ {
		if len(list1) > counter1 {
			result[resultCounter] = list1[counter1]
			counter1++
			resultCounter++
		}
		if len(list2) > counter2 {
			result[resultCounter] = list2[counter2]
			counter2++
			resultCounter++
		}
	}

	return result
}

// is the given string a palindrome
func isPalindrome(text string) bool {

	if len(text) <= 0 {
		return false
	}

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")
	var end int = len(text)

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[end-i-1] {
			return false
		}
	}
	return true
}

// is the element in the list
func isStringInList(name string, listOfNames []string) bool {
	for _, n := range listOfNames {
		if strings.Compare(n, name) == 0 {
			return true
		}
	}
	return false
}
