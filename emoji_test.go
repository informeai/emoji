package emoji

import (
	"testing"
)

//go test -v -run ^TestEmojiList
func TestEmojiList(t *testing.T) {
	emojis := emojiList()
	if emojis == nil {
		t.Errorf("TestEmojiList(): emojiList is nil")
	}
}

//go test -v -run ^TestEmojify
func TestEmojify(t *testing.T) {
	s := emojify("hello :heart:  is :+1: \nthis :relaxed:")
	if len(s) == 0 {
		t.Errorf("TestEmojify(): got -> %v, want: len > 0 ", len(s))
	}
}

//go test -v -run ^TestPrintln
func TestPrintln(t *testing.T) {
	err := Println("hello :fire:")
	if err != nil {
		t.Errorf("TestPrintln(): got -> %v, want: nil", err)
	}
}

//go test -v -run ^TestSprint
func TestSprint(t *testing.T) {
	s := Sprint("this is :star:")
	if len(s) == 0 {
		t.Errorf("TestSprint(): got -> %v, want: len > 0", len(s))
	}
}

//go test -v -run ^TestSprintf
func TestSprintf(t *testing.T) {
	s := Sprintf("%v", "why my :dog:")
	if len(s) == 0 {
		t.Errorf("TestSprintf(): got -> %v, want: len > 0", len(s))
	}
}
