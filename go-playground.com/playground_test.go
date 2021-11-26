package playground

import (
	"reflect"
	"testing"
)

func TestConvertToMorseCode(t *testing.T) {
	if convertToMorseCode("Hello") != ".... . .-.. .-.. ---" {
		t.Fatal()
	}
	if convertToMorseCode("Hello Me") != ".... . .-.. .-.. --- / -- ." {
		t.Fatal()
	}
	if convertToMorseCode("What is your name?") != ".-- .... .- - / .. ... / -.-- --- ..- .-. / -. .- -- . ..--.." {
		t.Fatal()
	}
}

func TestConvertFromMorseCode(t *testing.T) {
	if convertFromMorseCode(".... . .-.. .-.. ---") != "hello" {
		t.Fatal()
	}
	if convertFromMorseCode(".... . .-.. .-.. --- / -- .") != "hello me" {
		t.Fatal()
	}
	if convertFromMorseCode(".-- .... .- - / .. ... / -.-- --- ..- .-. / -. .- -- . ..--..") != "what is your name?" {
		t.Fatal()
	}
}

func TestGetFirstFibonacciNumbers(t *testing.T) {
	if !reflect.DeepEqual(getFibonacciNumbersUpToLimit(0), []int{0}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(getFibonacciNumbersUpToLimit(2), []int{0, 1, 1, 2}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(getFibonacciNumbersUpToLimit(7), []int{0, 1, 1, 2, 3, 5}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(getFibonacciNumbersUpToLimit(35), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(getFibonacciNumbersUpToLimit(3000), []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584}) {
		t.Fatal()
	}
}

func TestConvertDigitToList(t *testing.T) {
	if !reflect.DeepEqual(convertDigitToList(0), []int{0}) {
		t.Fatal()
	}
	var result []int = convertDigitToList(123)
	if !reflect.DeepEqual(result, []int{1, 2, 3}) {
		t.Fatalf("Error: %d", result)
	}
}

func TestCombineLists(t *testing.T) {
	if !reflect.DeepEqual(combineLists(nil, []string{"a"}), []string{"a"}) {
		t.Fatal()
	}
	if !reflect.DeepEqual(combineLists([]string{"b"}, nil), []string{"b"}) {
		t.Fatal()
	}

	var list1 = []string{"a"}
	var list2 = []string{"1"}
	if !reflect.DeepEqual(combineLists(list1, list2), []string{"a", "1"}) {
		t.Fatal()
	}

	list1 = []string{"a", "b"}
	list2 = []string{"1"}
	if !reflect.DeepEqual(combineLists(list1, list2), []string{"a", "1", "b"}) {
		t.Fatal()
	}

	list1 = []string{"a"}
	list2 = []string{"1", "2"}
	if !reflect.DeepEqual(combineLists(list1, list2), []string{"a", "1", "2"}) {
		t.Fatal()
	}

	list1 = []string{"a", "b"}
	list2 = []string{"1", "2", "3", "4"}
	if !reflect.DeepEqual(combineLists(list1, list2), []string{"a", "1", "b", "2", "3", "4"}) {
		t.Fatal()
	}

	list1 = []string{"a", "b", "c", "d"}
	list2 = []string{"1", "2"}
	if !reflect.DeepEqual(combineLists(list1, list2), []string{"a", "1", "b", "2", "c", "d"}) {
		t.Fatal()
	}
}

func TestIsPalindrome(t *testing.T) {
	if isPalindrome("") {
		t.Fatal()
	}
	if !isPalindrome("Civic") {
		t.Fatal()
	}
	if !isPalindrome("Noon") {
		t.Fatal()
	}
	if !isPalindrome("Racecar") {
		t.Fatal()
	}
	if !isPalindrome("Taco cat") {
		t.Fatal()
	}
	if !isPalindrome("Rotavator") {
		t.Fatal()
	}
	if isPalindrome("Rottavator") {
		t.Fatal()
	}
}

func TestIsStringInList(t *testing.T) {
	names := []string{"Steve", "Kelly", "Jon"}
	if !isStringInList("Steve", names) {
		t.Fatal()
	}

	if isStringInList("Steeeeve", names) {
		t.Fatal()
	}
}
