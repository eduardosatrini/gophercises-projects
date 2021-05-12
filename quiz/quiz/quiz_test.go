package quiz

import (
	"reflect"
	"testing"
)

func TestConfigFlag(t *testing.T) {

	fileName := ConfigFlag()
	result := *fileName
	expected := "problems.csv"

	if result != expected {
		t.Errorf("result: %s, expected: %s", result, expected)
	}

}

func TestOpenCSVFile(t *testing.T) {

	result := OpenCSVFile("problems.csv")
	var expected [][]string

	if reflect.TypeOf(result) != reflect.TypeOf(expected) {
		t.Errorf("result: %T, expected: %T", result, expected)
	}

}
