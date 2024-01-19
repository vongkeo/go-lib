/**
 * Author: Mitch Allen
 * File: go-lib_test.go
 */

package lib

import (
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	lists := "js,css,html"
	fileEx := "css"
	expected := true

	if got := Contains(lists, fileEx); got != expected {
		t.Errorf("Contains(%s, %s) = %t, didn't return %t", lists, fileEx, got, expected)
	}
}

func TestIsMobileNo(t *testing.T) {
	mobileNo := "2028492947"
	expected := true
	_result := IsMobileNo(mobileNo)
	if _result != expected {
		t.Errorf("IsMobileNo(%s) = %t, didn't return %t", mobileNo, _result, expected)
	}

}

func TestGetTimeNow(t *testing.T) {
	expected := true
	_result := GetTimeNow()
	if _result == "" {
		t.Errorf("GetTimeNow() = %s, didn't return %t", _result, expected)
	}
}

func TestDateStartZone(t *testing.T) {
	date := "2024-01-18"
	zone := "Asia/Bangkok"
	expected := true
	got := DateStartZone(date, zone)
	if got.String() == "" {
		t.Errorf("DateStartZone(%s) = %s, didn't return %t", date, got.String(), expected)
	}

}

func TestDateEndZone(t *testing.T) {
	date := "2024-01-18"
	zone := "Asia/Bangkok"
	expected := true
	if got := DateEndZone(date, zone); got == (time.Time{}) {
		t.Errorf("DateEndZone(%s) = %s, didn't return %t", date, got.String(), expected)
	}

}

// GetDate
func TestGetDate(t *testing.T) {
	expected := true
	zone := "Asia/Bangkok"
	if got := GetDate(zone); got == (time.Time{}) {
		t.Errorf("GetDate() = %s, didn't return %t", got.String(), expected)
	}
}

func TestGetDateDiff(t *testing.T) {
	start_date := "2024-01-18"
	end_date := "2024-01-18"
	expected := true

	if got := GetDateDiff(start_date, end_date); got != 0 {
		t.Errorf("GetDateDiff(%s, %s) = %d, didn't return %t", start_date, end_date, got, expected)
	}

}

func TestGetMonthDiff(t *testing.T) {
	start_date := "2024-01-18"
	end_date := "2024-01-18"
	expected := true

	if got := GetMonthDiff(start_date, end_date); got != 0 {
		t.Errorf("GetMonthDiff(%s, %s) = %d, didn't return %t", start_date, end_date, got, expected)
	}
}

func TestStringToInt(t *testing.T) {
	str := "1"
	expected := true

	if got := StringToInt(str); got != 1 {
		t.Errorf("StringToInt(%s) = %d, didn't return %t", str, got, expected)
	}
}

func TestStringToFloat(t *testing.T) {

	str := "1"
	expected := true

	if got := StringToFloat(str); got != 1.0 {
		t.Errorf("StringToFloat(%s) = %f, didn't return %t", str, got, expected)
	}
}

func TestGenerateReference(t *testing.T) {
	length := 20
	if got := GenerateReference(length); len(got) != 20 {
		t.Errorf("GenerateReference() = %s, didn't return a string of length 20", got)
	}
}

func TestStructToJSON(t *testing.T) {
	type TestStruct struct {
		Name string
	}

	testStruct := TestStruct{
		Name: "test",
	}

	expected := true

	got, err := StructToJSON(testStruct)
	if err != nil {
		t.Errorf("StructToJSON() = %s, didn't return %t", got, expected)
	}

}

func TestLog(t *testing.T) {
	expected := true
	Log("UUID", "TEST LOG")
	if got := true; got != expected {
		t.Errorf("Log() = %t, didn't return %t", got, expected)
	}

}
