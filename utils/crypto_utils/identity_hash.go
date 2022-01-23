package crypto_utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

var (
	saltSize = 16
	IdentityHash identityHash = &identity{}
)

type identityHash interface {
	GenerateSalt(salt string) ([]byte, error)
	GenerateIdentityKey(key string, salt []byte) string
	DoMatch(identityKey string, currentKey string, salt []byte) bool
}

type identity struct{}

func (i *identity) GenerateSalt(salt string) ([]byte, error) {
	// convert the slice of the byte
	val := []byte(salt)
	//make the additional salt of the given salt size
	addVal := make([]byte, saltSize)

	//append both the salts to create the new salt value
	finalSaltVal := append(val, addVal...)

	//read the rand salt
	if _, err := rand.Read(finalSaltVal); err != nil {
		return nil, err
	}

	return finalSaltVal, nil
}

func (i *identity) GenerateIdentityKey(key string, salt []byte) string {
	//convert the key to the slice of byte
	keyBytes := []byte(key)
	// create the sha512 hasher
	sha512Hasher := sha512.New()
	//append the salt and key
	appendedKey := append(keyBytes, salt...)
	//write into the sha512 hasher
	sha512Hasher.Write(appendedKey)
	//get the sha-512 hashed key
	hashedKey := sha512Hasher.Sum(nil)
	//convert the hashedKey to base64
	base64EncodedHashKey := base64.URLEncoding.EncodeToString(hashedKey)

	return base64EncodedHashKey
}

func (i *identity) DoMatch(identityKey string, currentKey string, salt []byte) bool {
	currentHash := i.GenerateIdentityKey(currentKey, salt)

	return identityKey == currentHash
}
