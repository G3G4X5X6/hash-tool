package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"strconv"

	"github.com/fatih/color"
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

type Result struct {
	crc32  string
	md5    string
	sha1   string
	sha256 string
	sha512 string
}

func StringHashReport(text string, algorithms []string) {
	fmt.Println(text)
	var reportMsg Result

	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	color.Cyan("——————————————————————————————————————————————————————————")
	fmt.Printf("%s: %s\n", red("Text"), cyan(text))
	color.Cyan("——————————————————————————————————————————————————————————")
	for _, algorithm := range algorithms {
		if algorithm == "crc32" {
			reportMsg.crc32 = strCRC32(text)
			fmt.Println("crc32:\t", reportMsg.crc32)
		}
		if algorithm == "md5" {
			reportMsg.md5 = strMd5(text)
			fmt.Println("md5:\t", reportMsg.md5)
		}
		if algorithm == "sha1" {
			reportMsg.sha1 = strSha1(text)
			fmt.Println("sha1:\t", reportMsg.sha1)
		}
		if algorithm == "sha256" {
			reportMsg.sha256 = strSha256(text)
			fmt.Println("sha256:\t", reportMsg.sha256)
		}
		if algorithm == "sha512" {
			reportMsg.sha512 = strSha512(text)
			fmt.Println("sha512:\t", reportMsg.sha512)
		}
	}

	green := color.New(color.FgGreen).SprintFunc()
	color.Cyan("——————————————————————————————————————————————————————————")
	fmt.Printf("%s\n", green("Done!"))
	fmt.Println("")
}
