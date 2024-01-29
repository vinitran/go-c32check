package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
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

var ADDR_BITCOIN_TO_STACKS = map[int]int{
	0:   22, // P
	5:   20, // M
	111: 26, // T
	196: 21, // N
}

func B58toC32(b58check string, version int) string {
	addrInfo, addrVersionByte, err := base58.CheckDecode(b58check)
	if err != nil {
		panic(err)
	}

	hash160String := hex.EncodeToString(addrInfo)
	stacksVersion := version
	if stacksVersion < 0 {
		stacksVersion = int(addrVersionByte)
		if _, ok := ADDR_BITCOIN_TO_STACKS[int(addrVersionByte)]; ok {
			stacksVersion = ADDR_BITCOIN_TO_STACKS[int(addrVersionByte)]
		}
	}

	return C32address(stacksVersion, hash160String)
}

var ADDR_STACKS_TO_BITCOIN = map[int]int{
	22: 0,   // P
	20: 5,   // M
	26: 111, // T
	21: 196, // N
}

func C32toB58(c32string string, version int) string {
	stacksVersion, hash160String := C32addressDecode(c32string)
	bitcoinVerion := version

	if version < 0 {
		bitcoinVerion = stacksVersion
		if _, ok := ADDR_STACKS_TO_BITCOIN[stacksVersion]; ok {
			stacksVersion = ADDR_BITCOIN_TO_STACKS[stacksVersion]
		}
	}

	prefix := fmt.Sprintf("%x", bitcoinVerion)
	if len(prefix) == 1 {
		prefix = fmt.Sprintf("0%s", prefix)
	}

	return Encode(hash160String, prefix)
}
