package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertTime = StringToTime("2024-02-02T12:10:00")

	if convertTime.Year() != 2024 {
		t.Error("convert time should be 2024")
	}

	if convertTime.Month() != 2 {
		t.Error("convert time should be 2")
	}

	if convertTime.Day() != 2 {
		t.Error("convert time should be 2")
	}

	if convertTime.Hour() != 12 {
		t.Error("convert time should be 12")
	}

	if convertTime.Minute() != 10 {
		t.Error("convert time should be 10")
	}

	if convertTime.Second() != 0 {
		t.Error("convert time should be 0")
	}

}
