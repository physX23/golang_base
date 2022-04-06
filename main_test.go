package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err != nil {
		t.Errorf("Test failed: error is")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("Test failed - result not match\n %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1
`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err == nil {
		t.Errorf("Test failed: error is %v", err)
	}
}
