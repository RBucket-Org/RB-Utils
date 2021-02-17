package crypto_utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

//GetMd5 : this function returns the hex and md5 encoded value
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

//GetSha512 : this function returns the sha512 value encoded with and string value
func GetSha512(input string) string {
	hash := sha512.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
