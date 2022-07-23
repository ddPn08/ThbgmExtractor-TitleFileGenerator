package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func reverseBytes(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i += 1 {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
	return bytes
}

func byteHexToInt(bytes []byte) int {
	h := hex.EncodeToString(bytes)
	i, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func intToByteHex(num int) []byte {
	h := fmt.Sprintf("%x", num)
	for len(h) < 8 {
		h = "0" + h
	}
	b, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func byteTo8ByteHex(bytes []byte) string {
	h := hex.EncodeToString(bytes)
	for len(h) < 8 {
		h = "0" + h
	}
	return h
}
