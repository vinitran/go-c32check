package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC32checkEncode(t *testing.T) {
	dataHex := "a46ff88886c2ef9762d970b4d2c63678835bd39d"
	version := 22
	encode := C32checkEncode(version, dataHex)
	assert.Equal(t, encode, "P2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7", "invalid")
}

func TestC32checkEncode1(t *testing.T) {
	dataHex := ""
	version := 0
	encode := C32checkEncode(version, dataHex)
	assert.Equal(t, encode, "0A0DR2R", "invalid")
}

func TestC32checkEncode2(t *testing.T) {
	dataHex := "1000000000000000000000000000000000000001"
	version := 22
	encode := C32checkEncode(version, dataHex)
	assert.Equal(t, encode, "P80000000000000000000000000000004R0CMNV", "invalid")
}

func TestC32checkDecode(t *testing.T) {
	dataHex := "P2J6ZY48GV1EZ5V2V5RB9MP66SW86PYKKNRV9EJ7"
	version, decode := C32checkDecode(dataHex)
	assert.Equal(t, decode, "a46ff88886c2ef9762d970b4d2c63678835bd39d", "invalid decode")
	assert.Equal(t, version, 22, "invalid version")
}

func TestC32checkDecode1(t *testing.T) {
	dataHex := "0A0DR2R"
	version, decode := C32checkDecode(dataHex)
	assert.Equal(t, decode, "", "invalid decode")
	assert.Equal(t, version, 0, "invalid version")
}

func TestC32checkDecode3(t *testing.T) {
	dataHex := "P80000000000000000000000000000004R0CMNV"
	version, decode := C32checkDecode(dataHex)
	assert.Equal(t, decode, "1000000000000000000000000000000000000001", "invalid decode")
	assert.Equal(t, version, 22, "invalid version")
}
