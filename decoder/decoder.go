package decoder

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
)

const (
	bracketOpen rune = '['
	bracketClose rune = ']'
	delimiter string = " "
	repeatSearchLength int = 128
	ErrorMessage string = "Error\n"
)


func Decode(input string) (error, string) {
	output := ""
	buffer := ""
	bufferOpen := false

	for _, letter := range input {
		if letter == bracketOpen {
			bufferOpen = true
			continue
		}

		if letter == bracketClose {
			if bufferOpen {
				parsed := parseEncoding(buffer)
				if parsed == ErrorMessage {
					return errors.New(ErrorMessage), ""
				}

				buffer = ""
				output += parsed
				bufferOpen = false

			} else {
				return errors.New(ErrorMessage), ""
			}
			continue
		}

		if bufferOpen {
			buffer += string(letter)
			continue
		}

		output += string(letter)
	}

	return nil, output
}


func Encode(input string) (string, float32){
	if len(input) == 1 {
		return input, 0.0
	}

	var buffer string
	var output string
	var skip int = 0
	var searchLength int = repeatSearchLength

	if len(input) < searchLength {
		searchLength = len(input) / 2
	}

	for i, _ := range input {
		if skip > 0 {
			skip--
			continue
		}

		buffer = ""

		// check for repetitions up to some amount of characters
		// the time complexity of this function becomes a problem when searching for longer patterns
		// however, if the intended usage is only for short inputs this works fine

		// see more:
		// https://en.wikipedia.org/wiki/String-searching_algorithm#Naive_string_search

		for j := range searchLength {
			// end of input
			if i+j >= len(input) {
				output += string(buffer[0])
				break
			}

			buffer += string(input[i+j])

			index := i
			duplicates := 0

			// check repeats
			for {
				// end of input, no more duplicates can be found
				if index+len(buffer) > len(input) {
					break
				}

				if string(input[index:index+len(buffer)]) == buffer {
					duplicates++
				} else {
					break
				}

				index += len(buffer)
			}

			if duplicates > 1 {
				output += fmt.Sprintf("[%d %s]", duplicates, buffer)
				skip = duplicates * len(buffer) - 1

				break
	
			} else if j == searchLength - 1 {
				output += string(buffer[0])
			}
		}
	}

	shortened := (1.0 - float32(len(output)) / float32(len(input))) * 100.0

	return output, shortened
}


func parseEncoding(buffer string) string {
	delimiterIndex := strings.Index(buffer, delimiter)
	if delimiterIndex == -1 {
		return ErrorMessage
	}

	strNum := buffer[:delimiterIndex]
	strArt := buffer[delimiterIndex + 1:]

	repeats, err := strconv.Atoi(strNum)

	if err != nil {
		return ErrorMessage
	}

	return strings.Repeat(strArt, repeats)
}




