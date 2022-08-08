package solution_67_test

import (
	"bytes"
	"strconv"
	"testing"
)

func addBinary(a string, b string) string {

	return Bitcalculator(a, b)
}

func Bitcalculator(a string, b string) string {

	var buffer bytes.Buffer
	alen := len(a) - 1
	blen := len(b) - 1
	carry := 0

	for alen >= 0 || blen >= 0 || carry != 0 {

		var byte1 int
		var byte2 int

		if alen >= 0 {
			byte1 = convertToBinary(a[alen])
		} else {
			byte1 = 0
		}

		if blen >= 0 {
			byte2 = convertToBinary(b[blen])
		} else {
			byte2 = 0
		}

		alen--
		blen--

		result := byte1 + byte2

		if result == 2 && carry == 0 {
			carry = 1
			result = 0
		} else if result == 2 && carry == 1 {
			carry = 1
			result = 1
		} else if result == 1 && carry == 1 {
			carry = 1
			result = 0
		} else if result == 0 && carry == 1 {
			result = 1
			carry = 0
		}
		buffer.WriteString(strconv.Itoa(result))
	}

	return Reverse(buffer.String())
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convertToBinary(s byte) int {
	if s == '1' {
		return 1
	}
	return 0
}

func TestXxx(t *testing.T) {

	result := addBinary("11", "1")

	if result != "100" {
		t.Errorf("Result [%s] Not Equals 100", result)
	}
}

func TestXxx2(t *testing.T) {

	result := addBinary("1010", "1011")

	if result != "10101" {
		t.Errorf("Result [%s] Not Equals 100", result)
	}
}

func TestXxx3(t *testing.T) {

	result := addBinary("1111", "1111")

	if result != "11110" {
		t.Errorf("Result [%s] Not Equals 100", result)
	}
}
