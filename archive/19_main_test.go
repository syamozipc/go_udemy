// main.goのテスト用ファイル
// {対象ファイル名}_test.go と命名する

package main

import (
	"testing"
)

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
} {
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

// func名はTestで始まらなければならない
func TestDivision(t *testing.T){
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one")
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}