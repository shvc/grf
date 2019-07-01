package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	mrand "math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

var version = "unknow"

func toBytes(s string) (int64, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	i := strings.IndexFunc(s, unicode.IsLetter)
	if i == -1 {
		return strconv.ParseInt(s, 10, 64)
	}

	bytesString, multiple := s[:i], s[i:]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil || bytes <= 0 {
		return 0, errors.New("invalid size")
	}

	switch multiple {
	case "T", "TB":
		return int64(bytes * TERABYTE), nil
	case "G", "GB":
		return int64(bytes * GIGABYTE), nil
	case "M", "MB":
		return int64(bytes * MEGABYTE), nil
	case "K", "KB":
		return int64(bytes * KILOBYTE), nil
	case "B":
		return int64(bytes), nil
	default:
		return 0, errors.New("invalid size")
	}
}

// gen random bytes
func randomBytes(len uint) []byte {
	buf := make([]byte, len)
	_, err := rand.Read(buf)
	if err != nil {
		mrand.Seed(time.Now().UnixNano())
		for i := len - len; i < len; i++ {
			buf[i] = byte(mrand.Intn(128))
		}
	}
	return buf
}

func main() {
	num := flag.Int("n", 1, "number of files")
	size := flag.String("s", "1M", "size(K,M,G,T) of file")
	out := flag.String("o", ".", "output dir")
	prefix := flag.String("p", "vager", "filename prefix")
	ver := flag.Bool("v", false, "show version")
	flag.Parse()

	if *ver == true {
		fmt.Printf("version: %s\n", version)
		return
	}

	fileSize, err := toBytes(*size)
	if err != nil {
		fmt.Printf("invalid file size: %s, erro:%s", *size, err)
		return
	}
	fmt.Printf("gen %d file with size[%s] to %s\n", *num, *size, *out)
	ifd := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < *num; i++ {
		filename := filepath.Join(*out, fmt.Sprintf("%s-%s-%d", *prefix, *size, i))
		fd, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Create file %s failed %s\n", filename, err)
			break
		}
		if _, err = io.CopyN(fd, ifd, fileSize); err != nil {
			fmt.Printf("Gen file %s failed %s\n", filename, err)
			fd.Close()
			break
		}
		fd.Close()
	}
}
