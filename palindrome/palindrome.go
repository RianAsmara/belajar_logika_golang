package polindrome

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Make function to check reverse string.
*/

func reverseString(str string) (result string) {

	// looping to each character of string
	for _, v := range str {
		// add character as a string. and reverse
		result = string(v) + result
	}
	return
}

/*
Validate start number is less than end number
and make sure minimum number is 1 and maximum nunmber is 10^9
*/
/*
-	Input :
 	- start : Is the start number of range numbers.
 	- end 	: Is the end  number of range numbers.
 -	Output :
 	- validate : return true if valid, false if invalid
 	- err : return error if invalid
*/

func validateInput(start, end float64) (validate bool, err error) {
	// init minNumber
	minNumber := float64(1)

	// init maxNumber
	maxNumber := math.Pow(10, 9)

	// init error message if invalid
	err = errors.New("the value is invalid. Make sure you input the right value")
	// do checking
	if (start < end) && (start >= minNumber && start <= maxNumber) && (end >= minNumber && end <= maxNumber) {
		// return true when input is valid
		validate = true
		// and set err if value is null
		err = nil
	}
	return
}

/*

Make function to count how many palindrome number in range between two numbers
-	Input :
	- start : range of start number
	- end 	: range of end number
	- ex : 1,10 99,100
- Output :
	- total palindrome number between two number
	- return an error if there any error while finding palindrome number
*/
func countPalindromeNumber(start, end float64) (float64, error) {

	// init validate
	validate, err := validateInput(start, end)

	// init total palindrome number
	totalPalindromeNumber := float64(0)

	// check if value is valid
	if validate {

		// find and count palindrome number using looping
		for number := start; number <= end; number++ {

			// convert number to string
			convNumber := strconv.FormatFloat(number, 'f', 0, 64)

			// get reverse string converted number
			riverseStrNumber := reverseString(convNumber)

			// check the number is palindrome if converted number is equal to rever string
			if strings.Compare(convNumber, riverseStrNumber) == 0 {
				// icrese total of polindrom numbers
				totalPalindromeNumber++
			}
		}
		// set error is null to set the program is not error
		err = nil
	}
	return totalPalindromeNumber, err
}

/*
	Make function to count how many palindrome numbers
-	input 	:	Input must be string and split by space

-	output 	: 	Total palindrome number
*/

func PolindromeNumber(input string) (float64, error) {
	fmt.Println("[Count Palindrome Number]")
	// check input is correct
	if len(input) > 0 {
		// split input by empty space
		splitString := strings.Split(input, " ")
		// check size if not equal by 2
		if len(splitString) != 2 {
			return 0, errors.New("the input must be 2 number that split by empty space")
		}
		// init start and end number
		startNumb, _ := strconv.ParseFloat(splitString[0], 6)
		endNumb, _ := strconv.ParseFloat(splitString[1], 6)
		// print input value
		fmt.Println("Input : ", input)
		// get total palindrome number
		output, err := countPalindromeNumber(startNumb, endNumb)
		// print the output
		fmt.Println("Output : ", output)
		return output, err
	}
	return 0, errors.New("the input must be a string")

}
