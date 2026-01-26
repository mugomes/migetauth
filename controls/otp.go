// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package controls

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func base32Decode(secret string) ([]byte, error) {
	secret = strings.ToUpper(secret)
	secret = strings.ReplaceAll(secret, " ", "")

	return base32.StdEncoding.
		WithPadding(base32.NoPadding).
		DecodeString(secret)
}

func intToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	var b string
	for n > 0 {
		b = strconv.Itoa('0'+(n%2)) + b
		n /= 2
	}
	return b
}

func binaryToInt(s string) int {
	n := 0
	for _, c := range s {
		n <<= 1
		if c == '1' {
			n |= 1
		}
	}
	return n
}

func leftPad(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}

func GenerateTOTP(secret string) (string, error) {
	secret = strings.ToUpper(strings.ReplaceAll(secret, " ", ""))

	key, err := base32.StdEncoding.
		WithPadding(base32.NoPadding).
		DecodeString(secret)
	if err != nil {
		return "", err
	}

	counter := time.Now().Unix() / 30

	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(counter))

	h := hmac.New(sha1.New, key)
	h.Write(buf[:])
	hash := h.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	code :=
		(int(hash[offset])&0x7F)<<24 |
			(int(hash[offset+1])&0xFF)<<16 |
			(int(hash[offset+2])&0xFF)<<8 |
			(int(hash[offset+3]) & 0xFF)

	return fmt.Sprintf("%06d", code%1_000_000), nil
}
