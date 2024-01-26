package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

const c32 = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
const hexChars = "0123456789abcdef"

func C32encode(inputHex string, minLength int) string {
	if _, err := hex.DecodeString(inputHex); err != nil {
		panic("Not a hex-encoded string")
	}

	if len(inputHex)%2 != 0 {
		inputHex = "0" + inputHex
	}

	inputHex = strings.ToLower(inputHex)

	var res []rune
	carry := 0

	for i := len(inputHex) - 1; i >= 0; i-- {
		if carry < 4 {
			currentCode := strings.Index(hexChars, string(inputHex[i])) >> carry
			var nextCode int
			if i != 0 {
				nextCode = strings.Index(hexChars, string(inputHex[i-1]))
			}
			nextBits := 1 + carry
			nextLowBits := nextCode % (1 << nextBits) << (5 - nextBits)
			curC32Digit := rune(c32[currentCode+nextLowBits])
			carry = nextBits
			res = append([]rune{curC32Digit}, res...)
		} else {
			carry = 0
		}
	}

	C32leadingZeros := 0
	for _, r := range res {
		if r != '0' {
			break
		} else {
			C32leadingZeros++
		}
	}

	res = res[C32leadingZeros:]

	zeroPrefix := strings.Repeat("0", len([]rune(inputHex))-len([]rune(strings.TrimLeft(inputHex, "0"))))
	numLeadingZeroBytesInHex := len(zeroPrefix)

	for i := 0; i < numLeadingZeroBytesInHex; i++ {
		res = append([]rune{rune(c32[0])}, res...)
	}

	if minLength > 0 && minLength < len(res) {
		res = res[len(res)-minLength:]
	}

	return string(res)
}

func C32normalize(c32input string) string {
	return strings.NewReplacer("O", "0", "I", "1", "L", "1").Replace(strings.ToUpper(c32input))
}

func C32decode(c32input string, minLength int) string {
	c32input = C32normalize(c32input)
	match, _ := regexp.MatchString(fmt.Sprintf("^[%s]*$", c32), c32input)
	if !match {
		panic("Not a c32-encoded string")
	}

	zeroPrefixRegex := regexp.MustCompile("^" + string(c32[0]) + "*")
	zeroPrefix := zeroPrefixRegex.FindString(c32input)
	numLeadingZeroBytes := len(zeroPrefix)

	var res []rune
	carry := 0
	carryBits := 0

	for i := len(c32input) - 1; i >= 0; i-- {
		if carryBits == 4 {
			res = append([]rune{rune(hexChars[carry])}, res...)
			carryBits = 0
			carry = 0
		}

		currentCode := strings.Index(c32, string(c32input[i])) << carryBits
		currentValue := currentCode + carry
		currentHexDigit := rune(hexChars[currentValue%16])
		carryBits += 1
		carry = currentValue >> 4

		if carry > 1<<carryBits {
			panic("Panic error in decoding.")
		}

		res = append([]rune{currentHexDigit}, res...)
	}

	// One last carry
	res = append([]rune{rune(hexChars[carry])}, res...)

	if len(res)%2 == 1 {
		res = append([]rune{'0'}, res...)
	}

	hexLeadingZeros := 0
	for _, r := range res {
		if r != '0' {
			break
		} else {
			hexLeadingZeros++
		}
	}

	res = res[hexLeadingZeros-(hexLeadingZeros%2):]

	hexStr := string(res)
	for i := 0; i < numLeadingZeroBytes; i++ {
		hexStr = "00" + hexStr
	}

	if minLength > 0 {
		count := minLength*2 - len(hexStr)
		for i := 0; i < count; i += 2 {
			hexStr = "00" + hexStr
		}
	}

	return hexStr
}
