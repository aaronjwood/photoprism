package fs

import (
	"crypto/sha512"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
)

// Hash returns the hash of a file as string.
func Hash(fileName string) string {
	var result []byte

	file, err := os.Open(fileName)

	if err != nil {
		return ""
	}

	defer file.Close()

	hash := sha512.New512_256()

	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(result))
}

// Checksum returns the CRC32 checksum of a file as string.
func Checksum(fileName string) string {
	var result []byte

	file, err := os.Open(fileName)

	if err != nil {
		return ""
	}

	defer file.Close()

	hash := crc32.New(crc32.MakeTable(crc32.Castagnoli))

	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(result))
}

// IsHash tests if a string looks like a hash.
func IsHash(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 48 || r > 57) && (r < 97 || r > 102) && (r < 65 || r > 70) {
			return false
		}
	}

	switch len(s) {
	case 8, 16, 32, 40, 56, 64, 80, 128, 256:
		return true
	}

	return false
}
