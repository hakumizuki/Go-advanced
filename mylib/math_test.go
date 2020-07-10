package mylib

import(
	"testing"
)

func TestAverage(t *testing.T) { // 関数名はTestXxx() $go test ./...で全てのファイル中のテストファイルをテストする
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v) // tのメソッドを使う
	}
}