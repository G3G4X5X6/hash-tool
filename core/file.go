package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func CalculateFileCRC32(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	table := crc32.MakeTable(crc32.IEEE)

	hash := crc32.New(table)
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	crc32Value := hash.Sum32()
	crc32String := strconv.FormatUint(uint64(crc32Value), 16)
	return crc32String, nil
}

func CalculateFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	md5Hash := fmt.Sprintf("%x", hash.Sum(nil))
	return md5Hash, nil
}

func CalculateFileSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sha1Hash := fmt.Sprintf("%x", hash.Sum(nil))
	return sha1Hash, nil
}

func CalculateFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sha256Hash := hash.Sum(nil)
	sha256String := hex.EncodeToString(sha256Hash)
	return sha256String, nil
}

func CalculateFileSHA512(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	hash := sha512.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sha512Hash := hash.Sum(nil)
	sha512String := hex.EncodeToString(sha512Hash)
	return sha512String, nil
}

func fileHash(filePath string, algorithms []string, threadsChan *chan int) {

	var fileResult Result
	fileResult.text = filePath

	for _, algorithm := range algorithms {
		if algorithm == "crc32" {
			crc32Value, err := CalculateFileCRC32(filePath)
			if err != nil {
				fmt.Println("Error calculating file CRC32:", err)
				return
			}
			fileResult.crc32 = crc32Value
		}
		if algorithm == "md5" {
			md5Hash, err := CalculateFileMD5(filePath)
			if err != nil {
				fmt.Println("Error calculating file MD5:", err)
				return
			}
			fileResult.md5 = md5Hash
		}
		if algorithm == "sha1" {
			sha1Hash, err := CalculateFileSHA1(filePath)
			if err != nil {
				fmt.Println("Error calculating file SHA1:", err)
				return
			}
			fileResult.sha1 = sha1Hash
		}
		if algorithm == "sha256" {
			sha256String, err := CalculateFileSHA256(filePath)
			if err != nil {
				fmt.Println("Error calculating file SHA256:", err)
				return
			}
			fileResult.sha256 = sha256String
		}
		if algorithm == "sha512" {
			sha512String, err := CalculateFileSHA512(filePath)
			if err != nil {
				fmt.Println("Error calculating file SHA512:", err)
				return
			}
			fileResult.sha512 = sha512String
		}
	}

	// 输出
	fileResult.toString()

	// 线程结束，释放资源信号
	<-*threadsChan
}

func FilesHashReport(paths []string, algorithms []string, threads int) string {

	startTime := time.Now()

	go func() {
		bar := progressbar.Default(-1, "Hashing......")
		for {
			_ = bar.Add(1)
		}
	}()

	var pathChan = make(chan string, 20)

	go func() {
		for _, path := range paths {

			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					pathChan <- path
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
		close(pathChan)
	}()

	// 线程数控制
	var threadsChan = make(chan int, threads)
	for path := range pathChan {
		threadsChan <- 1
		go fileHash(path, algorithms, &threadsChan)
	}

	for {
		if len(threadsChan) <= 0 {
			break
		}
	}

	fmt.Println()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Program runtime: %s\n", elapsedTime)
	color.Cyan("——————————————————————————————————————————————————————————\n\n")
	return ""
}
