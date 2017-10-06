package chars

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/tecnologer/Strings/Fonts"
)

//Parse - Returns string parsing the letters formed with plus sign
func Parse(word string) string {
	return Parsef(word, '*')
}

//Parsef - Returns string parsing the letters formed with specific sign
func Parsef(word string, sign rune) string {
	word = strings.ToUpper(word)
	word = removeSpecialChars(word)

	var finalWord string
	for row := 0; row < 7; row++ {
		for _, char := range word {
			if letter, ok := Fonts.DefaultFont[char]; ok {
				finalWord += letter[row]
			}
		}
		finalWord += "\r\n"
	}

	if sign != '*' && finalWord != "" {
		var regex, _ = regexp.Compile(`\*`)
		finalWord = regex.ReplaceAllString(finalWord, string(sign))
	}

	return finalWord
}

//ParsefNExport creates a file with the parsed to letters formed with specific sign
func ParsefNExport(text string, sign rune) {
	dir, err := os.Getwd()
	check(err)
	// For more granular writes, open a file for writing.
	f, err := os.Create(dir + "\\output.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(Parsef(text, sign))
	w.Flush()
}

func removeSpecialChars(word string) string {
	var specialCharacters = map[string]string{
		"A": "áà",
		"E": "éè",
		"I": "íì",
		"O": "óò",
		"U": "úù",
	}

	for key, characters := range specialCharacters {
		var pattern = "(?i)[" + characters + "]"
		var re, _ = regexp.Compile(pattern)
		word = re.ReplaceAllString(word, key)
	}

	// var re, _ = regexp.Compile(`\W*`)
	// word = re.ReplaceAllString(word, "")
	return word
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
