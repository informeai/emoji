package emoji

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`(:+[a-zA-Z0-9\+\-\(\)]+:)`)

//emojify is match of key in emojiList and return
//new string emojified.
func emojify(s string) string {
	em := emojiList()
	allPos := re.FindAllSubmatchIndex([]byte(s), -1)
	for i := 0; i < len(allPos); i++ {
		loc := re.FindIndex([]byte(s))
		start := loc[0]
		end := loc[1]
		pos := s[start:end]
		if val, ok := em[string(pos)]; ok {
			s = s[:start] + val + s[end:]
		}
	}
	return s
}

//Println print string emojified with new line.
func Println(s string) error {
	_, err := fmt.Println(emojify(s))
	return err
}

//Sprint return string emojified in single line.
func Sprint(s string) string {
	return fmt.Sprint(emojify(s))
}

//Sprintf return string emojified with format content.
func Sprintf(format string, a ...interface{}) string {
	return emojify(fmt.Sprintf(format, a...))
}
