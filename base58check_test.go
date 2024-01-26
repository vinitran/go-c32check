package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeWithDefaultPrefix(t *testing.T) {
	hexString := "f5f2d624cfb5c3f66d06123d0829d1c9cebf770e"
	hexStringExpected := "1PRTTaJesdNovgne6Ehcdu1fpEdX7913CK"
	dataEncode := Encode(hexString, "00")
	assert.Equal(t, dataEncode, hexStringExpected, "invalid")
}

func TestEncodeWithPrefix(t *testing.T) {
	hexString := "1E99423A4ED27608A15A2616A2B0E9E52CED330AC530EDCC32C8FFC6A526AEDD"
	prefix := "80"
	hexStringExpected := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"
	dataEncode := Encode(hexString, prefix)
	assert.Equal(t, dataEncode, hexStringExpected, "invalid")
}

func TestDecode1(t *testing.T) {
	hexString := "KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ"
	prefix, data, err := Decode(hexString)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, prefix, "80", "invalid prefix")
	assert.Equal(t, data, "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd", "invalid data")
}

func TestDecode2(t *testing.T) {
	hexString := "1PRTTaJesdNovgne6Ehcdu1fpEdX7913CK"
	prefix, data, err := Decode(hexString)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, prefix, "00", "invalid prefix")
	assert.Equal(t, data, "f5f2d624cfb5c3f66d06123d0829d1c9cebf770e", "invalid data")
}
