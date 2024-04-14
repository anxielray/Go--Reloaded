package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// STARTing modification
func modify(str string) string {
	var modifications, nModification, modifiedText, formattedText, formattedTxt, good, better string
	// check if there are enough arguments passed onto the commandline
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments passed...")
	}
	// read the contents of the inputfile
	fileContent, err := os.ReadFile(str)
	if err != nil {
		fmt.Println(err)
	}
	pt := regexp.MustCompile(`\s\,`)
	better = pt.ReplaceAllString(string(fileContent), ", ")
	// slice the content of the file and append each to an array, this is done by checking for the white spaces and loop over it
	fileArray := strings.Fields(better)
	for index, word := range fileArray {
		// modify for the upper and remove the (up)
		if word == "(up)" {
			fileArray[index-1] = strings.ToUpper(fileArray[index-1])
			fileArray[index] = ""
			// modify the lower and remove the (low)
		} else if word == "(low)" {
			fileArray[index-1] = strings.ToLower(fileArray[index-1])
			fileArray[index] = ""
			// modify the capitalization  and remove the (cap)
		} else if word == "(cap)" {
			fileArray[index-1] = capitalizeFirstLetter(fileArray[index-1])
			fileArray[index] = ""
			// modify for the vowels and h... words
		} else if word == "a" {
			wordToChange := fileArray[index+1]
			slicedRuneOfWordToChange := []rune(wordToChange)
			if slicedRuneOfWordToChange[0] == 'a' || slicedRuneOfWordToChange[0] == 'e' ||
				slicedRuneOfWordToChange[0] == 'i' || slicedRuneOfWordToChange[0] == 'o' ||
				slicedRuneOfWordToChange[0] == 'u' || slicedRuneOfWordToChange[0] == 'h' {
				fileArray[index] = "an"
			}
			// modify for hexadecimal to decimal number
		} else if word == "(hex)" {
			number := strconv.FormatInt(hexToDecimal(fileArray[index-1]), 10)
			fileArray[index-1] = number
			fileArray[index] = ""
			// modify for binary number to decimal number
		} else if word == "(bin)" {
			nbr := strconv.FormatInt(binToDecimal(fileArray[index-1]), 10)
			fileArray[index-1] = nbr
			fileArray[index] = ""
			// check for (cap, n) modify and remove (cap, n)
		} else if word == "(cap," {
			changeToString := string(fileArray[index+1])
			changeToString = strings.TrimRight(changeToString, ")")
			num, err := strconv.Atoi(changeToString)
			if err != nil {
				fmt.Println(err)
			} else {
				for l := 1; l <= num; l++ {
					fileArray[index-l] = capitalizeFirstLetter(fileArray[index-l])
					fileArray[index] = ""
					fileArray[index+1] = ""
				}
			}
			// check for (low, n) modify it and remove the (low, n)
		} else if word == "(low," {
			changeToString := string(fileArray[index+1])
			changeToString = strings.TrimRight(changeToString, ")")
			num, err := strconv.Atoi(changeToString)
			if err != nil {
				fmt.Println(err)
			} else {
				for l := 1; l <= num; l++ {
					fileArray[index-l] = strings.ToLower(fileArray[index-l])
					fileArray[index] = ""
					fileArray[index+1] = ""
				}
			}
			// check for (up, n) modify and remove the (up, n)
		} else if word == "(up," {
			changeToString := string(fileArray[index+1])
			changeToString = strings.TrimRight(changeToString, ")")
			num, err := strconv.Atoi(changeToString)
			if err != nil {
				fmt.Println(err)
			} else {
				for l := 1; l <= num; l++ {
					fileArray[index-l] = strings.ToUpper(fileArray[index-l])
					fileArray[index] = ""
					fileArray[index+1] = ""
				}
			}
		}
		// convert the []string and convert it to a single string
		modifications = strings.Join(fileArray, " ")
		// remove the extra spaces left out by removing the words in brackets
		nModification = strings.ReplaceAll(modifications, "  ", " ")
		// check for the punctuation marks and modify accordingly
		re := regexp.MustCompile(`\s*([.,!?;:])`)
		modifiedText = re.ReplaceAllString(nModification, "$1")
		// check for the multiple punctuations
		pattern := regexp.MustCompile(`'\s*(.*?)\s*'`)
		formattedText = pattern.ReplaceAllString(modifiedText, "'$1'")
		patt := regexp.MustCompile(`(\s\.\.\.\s)`)
		formattedTxt = patt.ReplaceAllString(formattedText, "... ")
		ret := regexp.MustCompile(`(\!\?\s*)`)
		good = ret.ReplaceAllString(formattedTxt, "!?")

	}
	return good
}

func capitalizeFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(string(s[1:]))
}

func hexToDecimal(s string) int64 {
	var result int64
	for _, digit := range s {
		var value int64
		switch {
		case digit >= '0' && digit <= '9':
			value = int64(digit - '0')
		case digit >= 'a' && digit <= 'z':
			value = int64(digit - 'a')
		case digit >= 'A' && digit <= 'Z':
			value = int64(digit - 'A')
		default:
			return 0
		}
		result = result*16 + value
	}
	return result
}

func binToDecimal(stn string) int64 {
	var finalAnswer int64
	for _, num := range stn {
		if num != '0' && num != '1' {
			return 0
		}
		finalAnswer = finalAnswer*2 + int64(num-'0')
	}
	return finalAnswer
}

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	// opening the sample.txt file
	openedFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer openedFile.Close()
	// read the content of the opened file
	_, err = openedFile.Stat()
	if err != nil {
		fmt.Println(err)
	}
	// modify the data input from the previous file sample.txt using the modifications unified under one function modify()
	modifications := []byte(modify(inputFile))
	// create a new file result.txt to save our modifications
	newFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	// write the modified data into the new file
	_, err = newFile.Write(modifications)
	if err != nil {
		fmt.Println(err)
	}
}
