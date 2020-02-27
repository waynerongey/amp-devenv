package utilities

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//CreateUUID .
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

//Encrypt .
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

//GetIPFromAddress .
func GetIPFromAddress(address string) string {
	if strings.Contains(address, `:`) {
		return strings.Split(address, `:`)[0]
	}
	return address
}

//StringToUInt transforms string to uint64
func StringToUInt(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil || i == 0 {
		i = 1
	}
	return uint(i)
}
