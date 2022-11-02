package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the card number to be validated: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ccNum := make([]int, 16, 16)
	i := 0
	for _, rune := range input {
		n, err := strconv.Atoi(string(rune))
		if err == nil {
			ccNum[i] = n
			i++
		}
	}

	ccNum = ccNum[:i]

	valid, cardType := ValidateCard(ccNum)
	if valid && cardType != "Unknown" {
		fmt.Printf("Card number is a valid %s card\n", cardType)
	} else {
		fmt.Println("Card number is NOT valid!")
	}
}

func ValidateCard(cNum []int) (valid bool, cardType string) {
	valid = luhnCheck(cNum)
	cardType = getCardType(cNum)
	return valid, cardType
}

func getCardType(cNum []int) string {
	var sArray = make([]string, len(cNum))
	for i, n := range cNum {
		sArray[i] = strconv.Itoa(n)
	}
	s := strings.Join(sArray, "")
	match, _ := regexp.MatchString("^4[0-9]{12}(?:[0-9]{3})?$", s)
	if match {
		return "Visa"
	}
	match, _ = regexp.MatchString("^5[1-5][0-9]{14}$", s)
	if match {
		return "MasterCard"
	}
	match, _ = regexp.MatchString("^(?:220[0-4])\\d+$", s)
	if match {
		return "MIR"
	}
	return "Unknown"
}

func luhnCheck(cNum []int) bool {
	checksum := 0
	for i, n := range cNum {
		if (i+len(cNum)%2)%2 == 0 {
			checksum += sumDigits(2 * n)
		} else {
			checksum += n
		}
	}

	return checksum%10 == 0
}

func sumDigits(n int) int {
	hundreds := n / 100
	tens := (n - 100*hundreds) / 10
	ones := n - 100*hundreds - 10*tens
	return ones + tens + hundreds
}
