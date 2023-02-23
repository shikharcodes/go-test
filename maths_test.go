package maths

import (
	"fmt"
	"testing"
)

type addTest struct {
	arg1, arg2, expected_result int
}

var addTests = []addTest{
	addTest{2, 4, 6},
	addTest{4, 3, 7},
	addTest{12, 11, 23},
	addTest{9, 6, 15},
	addTest{8, 4, 12},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected_result {
			t.Errorf("Output %q is not equal to expected result %q", output, test.expected_result)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 8)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 8)) // output: 12
}
