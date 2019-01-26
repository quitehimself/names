package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestProgram(t *testing.T) {
	outputRE := regexp.MustCompile("^[a-z]* [a-z]* [a-z]*\n$")
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		panic(fmt.Sprintf("unable to intercept stdout for main: %s", err))
	}
	oldStdout := os.Stdout
	os.Stdout = stdoutWriter
	main()
	os.Stdout.Close()
	os.Stdout = oldStdout
	mainOut, err := ioutil.ReadAll(stdoutReader)
	if err != nil {
		panic(fmt.Sprintf("unable to read stdout for main: %s", err))
	}
	if !outputRE.Match(mainOut) {
		t.Error("main output is not three space-separated words")
	}
}
