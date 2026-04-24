package decoder

import (
	"strings"
	"strconv"
	"fmt"
)

const (
	bracketOpen rune = '['
	bracketClose rune = ']'
	delimiter string = " "
	encodingPatternLength int = 3
	ErrorMessage string = "Error\n"
)


func Decode(input string) string {
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
					return ErrorMessage
				}

				buffer = ""
				output += parsed
				bufferOpen = false

			} else {
				return ErrorMessage
			}
			continue
		}

		if bufferOpen {
			buffer += string(letter)
			continue
		}

		output += string(letter)
	}

	return output
}


func Encode(input string) string {
	var buffer string
	var output string
	var skip int = 0

	for i, _ := range input {
		if skip > 0 {
			skip--
			continue
		}

		buffer = ""

		// check a certain amount of characters for repetitions
		// there is a better way to do this but am keeping it in since it works well enough
		for j := range encodingPatternLength {
			// end of input, no more duplicates can be found
			if i+j >= len(input) {
				output += buffer
				break
			}

			buffer += string(input[i+j])

			index := i
			duplicates := 0

			// check repeats
			for {
				// end of input, no more duplicates can be found
				if index+len(buffer) >= len(input) {
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
	
			} else if j == encodingPatternLength - 1 {
				output += string(buffer[0])
			}
		}
	}

	return output
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




