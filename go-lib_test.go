/**
 * Author: Mitch Allen
 * File: go-lib_test.go
 */

package lib

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b

	if got := Add(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubtract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b

	if got := Subtract(a, b); got != expected {
		t.Errorf("Subtract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestContains(t *testing.T) {
	lists := "js,css,html"
	fileEx := "css"
	expected := true

	if got := Contains(lists, fileEx); got != expected {
		t.Errorf("Contains(%s, %s) = %t, didn't return %t", lists, fileEx, got, expected)
	}
}
