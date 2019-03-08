package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	mrand "math/rand"
	"os"
	"path/filepath"
	"time"
)

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
	size := flag.Int64("s", 1024, "size of file")
	out := flag.String("o", ".", "output dir")
	flag.Parse()

	fmt.Printf("gen %d file[%d] to %s\n", *num, *size, *out)
	ifd := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < *num; i++ {
		filename := filepath.Join(*out, fmt.Sprintf("file-%d-%d", *size, i))
		fd, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Create file %s failed %s\n", filename, err)
			break
		}
		if _, err = io.CopyN(fd, ifd, *size); err != nil {
			fmt.Printf("Gen file %s failed %s\n", filename, err)
			fd.Close()
			break
		}
		fd.Close()
	}
}
