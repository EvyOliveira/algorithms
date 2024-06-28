package main

import (
	"crypto/sha256"
	"fmt"
)

type Hash interface {
	GenerateHash(input string) string
}

type HashSHA256 struct{}

func (HashSHA256 *HashSHA256) GenerateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GenerateHash(input string, hash Hash) string {
	return hash.GenerateHash(input)
}

func main() {
	inputString := "hello, world!"

	hashSha256 := &HashSHA256{}
	hash := GenerateHash(inputString, hashSha256)
	fmt.Println("hash of", inputString, ":", hash)
}
