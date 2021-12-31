package string_test

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "abc"
	repeat := strings.Repeat(str, 5)
	expected := "abcabcabcabcabc"
	if repeat != expected {
		t.Errorf("Expected '%q', but '%q'", expected, repeat)
	}
}
