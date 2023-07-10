package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"hash/crc32"
	"strconv"
)

func strCRC32(text string) string {
	crc32Table := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum([]byte(text), crc32Table)
	return strconv.FormatUint(uint64(hash), 10)
}

func strMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strSha1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strSha256(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func strSha512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func StringHashReport(text string, algorithms []string) {
	var reportMsg Result
	reportMsg.text = text

	for _, algorithm := range algorithms {
		if algorithm == "crc32" {
			reportMsg.crc32 = strCRC32(text)
		}
		if algorithm == "md5" {
			reportMsg.md5 = strMd5(text)
		}
		if algorithm == "sha1" {
			reportMsg.sha1 = strSha1(text)
		}
		if algorithm == "sha256" {
			reportMsg.sha256 = strSha256(text)
		}
		if algorithm == "sha512" {
			reportMsg.sha512 = strSha512(text)
		}
	}

	reportMsg.toString()

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s\n", green("Done!"))
	fmt.Println("")
}
