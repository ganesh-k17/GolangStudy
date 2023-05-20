package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("password encrypted as :", encodeWithMD5(password))
}

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}
