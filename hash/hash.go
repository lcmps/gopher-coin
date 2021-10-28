package hash

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func EncodeToSha512(data []byte) string {
	h := sha512.Sum512(data)
	return hex.EncodeToString(h[:])
}

func ValidateSha512(data []byte, hash string) bool {
	return EncodeToSha512(data) == hash
}

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		fmt.Println("Error encoding:", err.Error())
	}
	return buf.Bytes()
}
