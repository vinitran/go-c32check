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
