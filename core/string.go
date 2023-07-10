package core

import "fmt"

func strMd5(text string) string {

	return ""
}

func strSha1(text string) string {

	return ""
}

func strSha256(text string) string {

	return ""
}

func strSha512(text string) string {

	return ""
}

type Result struct {
	md5    string
	sha1   string
	sha256 string
	sha512 string
}

func StringHashReport(text string, algorithms []string) {
	fmt.Println(text)
	var reportMsg Result

	for algorithm := range algorithms {
		fmt.Println(algorithm)
		fmt.Println(reportMsg)
	}
}
