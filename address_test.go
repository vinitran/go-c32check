package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC32address(t *testing.T) {
	dataHex := "a46ff88886c2ef9762d970b4d2c63678835bd39d"
	version := 22
	c32addr := C32address(version, dataHex)
	assert.Equal(t, c32addr, "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7", "invalid prefix")
}

func TestC32address1(t *testing.T) {
	dataHex := "1000000000000000000000000000000000000001"
	version := 22
	c32addr := C32address(version, dataHex)
	assert.Equal(t, c32addr, "SP80000000000000000000000000000004R0CMNV", "invalid prefix")
}

func TestC32addressDecode(t *testing.T) {
	dataHex := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"
	version, c32addr := C32addressDecode(dataHex)
	assert.Equal(t, version, 22, "invalid prefix")
	assert.Equal(t, c32addr, "a46ff88886c2ef9762d970b4d2c63678835bd39d", "invalid prefix")
}

func TestC32addressDecode1(t *testing.T) {
	dataHex := "SP80000000000000000000000000000004R0CMNV"
	version, c32addr := C32addressDecode(dataHex)
	assert.Equal(t, version, 22, "invalid prefix")
	assert.Equal(t, c32addr, "1000000000000000000000000000000000000001", "invalid prefix")
}

func TestB58toC32(t *testing.T) {
	b58check := "A7RjcihhakxJfAqgwTVsLTyc8kbhDJPMVY"
	c32check := "SP2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"
	c32 := B58toC32(b58check, -1)
	assert.Equal(t, c32, c32check, "invalid prefix")
	b58 := C32toB58(c32, -1)
	assert.Equal(t, b58, b58check, "invalid prefix")
}

func TestB58toC32_1(t *testing.T) {
	b58check := "9stsUTaRHnyTRFWnbwiyCWwfpkkKCFYBD4"
	c32check := "SP80000000000000000000000000000004R0CMNV"
	c32 := B58toC32(b58check, -1)
	assert.Equal(t, c32, c32check, "invalid prefix")
	b58 := C32toB58(c32, -1)
	assert.Equal(t, b58, b58check, "invalid prefix")
}

func TestB58toC32_2(t *testing.T) {
	b58check := "9stsUTaRHnyTRFWnbwiyCWwfpkkK9ZxEPC"
	c32check := "SP800000000000000000000000000000033H8YKK"
	c32 := B58toC32(b58check, -1)
	assert.Equal(t, c32, c32check, "invalid prefix")
	b58 := C32toB58(c32, -1)
	assert.Equal(t, b58, b58check, "invalid prefix")
}
