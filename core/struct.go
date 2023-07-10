package core

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Result struct {
	text   string
	crc32  string
	md5    string
	sha1   string
	sha256 string
	sha512 string
}

func (reportMsg Result) toString() {
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	color.Cyan("——————————————————————————————————————————————————————————")
	fmt.Printf("%s: %s\n", red("Text"), cyan(reportMsg.text))
	color.Cyan("——————————————————————————————————————————————————————————")

	fmt.Println("crc32:\t", strings.ToUpper(reportMsg.crc32))
	fmt.Println("md5:\t", strings.ToUpper(reportMsg.md5))
	fmt.Println("sha1:\t", strings.ToUpper(reportMsg.sha1))
	fmt.Println("sha256:\t", strings.ToUpper(reportMsg.sha256))
	fmt.Println("sha512:\t", strings.ToUpper(reportMsg.sha512))

	color.Cyan("——————————————————————————————————————————————————————————")
}
