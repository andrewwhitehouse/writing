package day1

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	actual, err := Parse("11\n13\n17\n19\n301")
	expected := []uint16{11, 13, 17, 19, 301}
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Parse was incorrect, got: %v, want: %v.", actual, expected)
	}
}
