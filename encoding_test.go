package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestC32Encode(t *testing.T) {
	inputHex := "a46ff88886c2ef9762d970b4d2c63678835bd39d"
	outputExpected := "MHQZH246RBQSERPSE2TD5HHPF21NQMWX"
	dataEncode := C32encode(inputHex, 0)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}

func TestC32Encode1(t *testing.T) {
	inputHex := ""
	outputExpected := ""
	dataEncode := C32encode(inputHex, 0)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}

func TestC32Encode2(t *testing.T) {
	inputHex := "0000000000000000000000000000000000000001"
	outputExpected := "00000000000000000001"
	dataEncode := C32encode(inputHex, 20)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}

func TestC32Decode(t *testing.T) {
	inputHex := "MHQZH246RBQSERPSE2TD5HHPF21NQMWX"
	outputExpected := "a46ff88886c2ef9762d970b4d2c63678835bd39d"
	dataEncode := C32decode(inputHex, 0)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}

func TestC32Decode1(t *testing.T) {
	inputHex := ""
	outputExpected := ""
	dataEncode := C32decode(inputHex, 0)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}

func TestC32Decode2(t *testing.T) {
	inputHex := "00000000000000000001"
	outputExpected := "0000000000000000000000000000000000000001"
	dataEncode := C32decode(inputHex, 20)
	assert.Equal(t, dataEncode, outputExpected, "invalid")
}
