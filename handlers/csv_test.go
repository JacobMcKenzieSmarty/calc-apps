package handler

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

const input = `1,+,2
2,-,1
NaN,+,2
1,+,NaN
1,nop,2
3,*,4
20,/,10`

const expectedOutput = `
1,+,2,3
2,-,1,1
3,*,4,12
20,/,10,2
`

func TestCSVHappyPath(t *testing.T) {
	var logOutput bytes.Buffer
	var logger = log.New(&logOutput, "", log.LstdFlags|log.Lmicroseconds)
	var output bytes.Buffer

	err := NewCSVHandler(strings.NewReader(input), &output, logger).Handle()
	assertError(t, nil, err)
	assertEqual(t, expectedOutput, output.String())
	if t.Failed() {
		t.Log("Log output:\n" + logOutput.String())
	}

	if err != nil {
		fmt.Println(err)
	}
}

func TestCSVReadError(t *testing.T) {
	var logOutput bytes.Buffer
	var logger = log.New(&logOutput, "", log.LstdFlags|log.Lmicroseconds)
	var input ErringReader
	boink := errors.New("BOINK")
	input.err = boink
	err := NewCSVHandler(input, &logOutput, logger).Handle()
	assertError(t, err, boink)
}

func TestCSVWriteError(t *testing.T) {
	var logOutput bytes.Buffer
	var logger = log.New(&logOutput, "", log.LstdFlags|log.Lmicroseconds)
	boink := errors.New("BOINK")
	var output ErringWriter
	output.err = boink

	err := NewCSVHandler(strings.NewReader(input), output, logger).Handle()
	assertError(t, err, boink)

}

type ErringReader struct {
	err error
}

func (e ErringReader) Read(p []byte) (n int, err error) {
	return 0, e.err
}

type ErringWriter struct {
	err error
}

func (e ErringWriter) Write(p []byte) (n int, err error) {
	return 0, e.err
}
