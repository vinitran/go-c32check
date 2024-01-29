package main

import (
	"fmt"
	"regexp"
)

func C32address(version int, hash160hex string) string {
	match, _ := regexp.MatchString("^[0-9a-fA-F]{40}$", hash160hex)
	if !match {
		panic("invalid argument: not a hash160 hex string")
	}

	c32string := C32checkEncode(version, hash160hex)
	return fmt.Sprintf("S%s", c32string)
}

func C32addressDecode(c32addr string) (int, string) {
	if len(c32addr) <= 5 {
		panic("invalid c32 address: invalid length")
	}

	if c32addr[0] != 'S' {
		panic("invalid c32 address: must start with \"S\"")
	}
	return C32checkDecode(c32addr[1:])
}
