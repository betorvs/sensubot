package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
)

//StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Md5Hash encodes a string using the MD5 algorithm
func Md5Hash(hashData string) string {
	hash := md5.Sum([]byte(hashData))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

//ValidateInt validates if a string is an valid integer
func ValidateInt(value string, name string) bool {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("[ERROR] ValidateInt " + name + " is not an int value: \"" + value + "\" | " + err.Error())
		return false
	}
	if (intVal == 0) && (name == "productId") {
		log.Printf("[INFO] ValidateInt ProductId 0 not allowed")
		return false
	}
	return true
}
