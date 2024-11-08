package main

import (
	"fmt"
	"strings"
)

var DEFDAULT_TEST_CMD = []string{"test:unit", "run"}
var DEFAULT_NPM_INSTALL_CMD = []string{"install"}
var DEFAULT_BUILD_CMD = []string{"build"}

func getTestCommand() string {
	return fmt.Sprint(DEFDAULT_TEST_CMD)
}

func testCommandToSlice(c string) []string {
	if c == "" {
		return DEFDAULT_TEST_CMD
	}
	return strings.Split(c, " ")
}

func buildCommandToSlice(c string) []string {
	if c == "" {
		return DEFAULT_BUILD_CMD
	}
	return strings.Split(c, " ")
}
