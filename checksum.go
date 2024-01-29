package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func C32checksum(dataHex string) string {
	dataBytes, err := hex.DecodeString(dataHex)
	if err != nil {
		panic(err)
	}

	dataHash := sha256.Sum256(dataBytes)
	dataHashHash := sha256.Sum256(dataHash[:])

	checksum := hex.EncodeToString(dataHashHash[:4])
	return checksum
}

func C32checkEncode(version int, data string) string {
	if version < 0 || version >= 32 {
		panic("invalid version (must be between 0 and 31)")
	}

	match, _ := regexp.MatchString("^[0-9a-fA-F]*$", data)
	if !match {
		panic("invalid data (not a hex string)")
	}

	dataHex := strings.ToLower(data)
	if len(data)%2 != 0 {
		dataHex = fmt.Sprintf("0%s", dataHex)
	}

	versionHex := fmt.Sprintf("%x", version)
	if len(versionHex) == 1 {
		versionHex = fmt.Sprintf("0%s", versionHex)
	}

	checksumHex := C32checksum(fmt.Sprintf("%s%s", versionHex, dataHex))
	c32checkData := fmt.Sprintf("%s%s", dataHex, checksumHex)
	c32str := C32encode(c32checkData, 0)
	result := fmt.Sprintf("%s%s", string(c32[version]), c32str)

	return result
}

func C32checkDecode(c32data string) (int, string) {
	c32data = C32normalize(c32data)
	dataHex := C32decode(c32data[1:], 0)

	versionChar := c32data[0]
	version := strings.Index(c32, string(versionChar))

	checksum := dataHex[len(dataHex)-8:]

	versionHex := fmt.Sprintf("%02x", version)

	// Verify checksum
	if C32checksum(versionHex+dataHex[:len(dataHex)-8]) != checksum {
		panic("Invalid c32check string: checksum mismatch")
	}

	return version, dataHex[:len(dataHex)-8]
}
