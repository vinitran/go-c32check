package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func Encode(data string, prefix string) string {
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	prefixBytes, err := hex.DecodeString(prefix)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	hash.Write(append(prefixBytes, dataBytes...))
	checksum := sha256.New()
	checksum.Write(hash.Sum(nil))

	checksumBytes := checksum.Sum(nil)[:4]
	resultBytes := append(append(prefixBytes, dataBytes...), checksumBytes...)

	encoded := base58.Encode(resultBytes)
	return encoded
}

func Decode(data string) (string, string, error) {
	decoded := base58.Decode(data)
	prefixBytes := decoded[:1]
	dataBytes := decoded[1 : len(decoded)-4]
	checksumBytes := decoded[len(decoded)-4:]

	hash := sha256.New()
	hash.Write(append(prefixBytes, dataBytes...))
	checksum := sha256.New()
	checksum.Write(hash.Sum(nil))

	for i := range checksumBytes {
		if checksumBytes[i] != checksum.Sum(nil)[i] {
			return "", "", fmt.Errorf("invalid checksum")
		}
	}

	//dataBytes = dataBytes[:len(dataBytes)-1]
	return hex.EncodeToString(prefixBytes), hex.EncodeToString(dataBytes), nil
}
