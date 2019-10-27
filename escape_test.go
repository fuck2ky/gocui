package gocui

import (
	"fmt"
	"testing"
)

func TestEscape(t *testing.T) {
	testCases := []struct {
		input string
		fg    Attribute
		bg    Attribute
	}{
		{
			input: "\033[48;5;200;38;5;100mHi!!",
			fg:    101,
			bg:    201,
		},
		{
			input: "\033[38;5;100;48;5;200mHi!!",
			fg:    101,
			bg:    201,
		},
		{
			input: "\033[38;5;100mHi!!",
			fg:    101,
			bg:    ColorDefault,
		},
		{
			input: "\033[48;5;100mHi!!",
			fg:    ColorDefault,
			bg:    101,
		},
		{
			input: "Hi!!",
			fg:    ColorDefault,
			bg:    ColorDefault,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			ei := newEscapeInterpreter(Output256)

			for _, r := range tc.input {
				_, err := ei.parseOne(r)
				if err != nil {
					t.Fatal(err)
				}
			}

			if ei.curFgColor != tc.fg {
				t.Fatalf("foreground color is not %d: %v", tc.fg, ei.curFgColor)
			}
			if ei.curBgColor != tc.bg {
				t.Fatalf("background color is not %d: %v", tc.bg, ei.curBgColor)
			}
		})
	}

}
