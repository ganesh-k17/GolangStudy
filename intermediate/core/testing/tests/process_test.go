package test

import (
	"example.com/testing/process"
	"testing"
)

func TestCheckEven(t *testing.T) { // Test method should start with caps.
	i := 9
	expected := "YES"
	res := process.CheckEven(i)
	if expected != res {
		t.Errorf("expected: %v, got:%v", expected, res)
	}
}

/*
Output:

$ go test
--- FAIL: TestCheckEven (0.00s)
    process_test.go:13: expected: YES, got:NO
FAIL
exit status 1
FAIL    example.com/testing     0.333s

*/
