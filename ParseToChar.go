package chars

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/tecnologer/ASCIIArt/Fonts"
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

func ParsefBackground(word string, sign rune, background rune) string {
	word = strings.ToUpper(word)
	word = removeSpecialChars(word)

	var finalWord string
	var lenght int
	for row := 0; row < 7; row++ {
		for _, char := range word {
			if letter, ok := Fonts.BoldFont[char]; ok {
				letterPart := letter[row]
				if row == 0 {
					lenght += len(letterPart)
				}

				finalWord += replace(letterPart, sign, background)
			}
		}
		finalWord += "\r\n"
	}

	return addHeaderFooter(finalWord, background, lenght)
}

func replace(letterPart string, char rune, bkg rune) string {
	var spce, _ = regexp.Compile(`\s`)
	var strks, _ = regexp.Compile(`\*`)

	if char == ' ' || bkg == '*' {
		letterPart = spce.ReplaceAllString(letterPart, "x")
		letterPart = strks.ReplaceAllString(letterPart, string(char))
		var special, _ = regexp.Compile(`x`)
		letterPart = special.ReplaceAllString(letterPart, string(bkg))
	} else {
		if char != '*' && letterPart != "" {
			var regex, _ = regexp.Compile(`\*`)
			letterPart = regex.ReplaceAllString(letterPart, string(char))
		}

		//reemplace white spaces to specific char
		if bkg != ' ' && bkg != char {
			var regex, _ = regexp.Compile(`\s`)
			letterPart = regex.ReplaceAllString(letterPart, string(bkg))
		}
	}

	return letterPart
}

func addHeaderFooter(finalWord string, bkg rune, lenght int) string {
	var HF string
	for i := lenght; i > 0; i-- {
		HF += string(bkg)
	}

	return HF + "\r\n" + finalWord + HF
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
