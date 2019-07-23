package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
)

const (
	myByte = 1 << (10 * iota)
	myKilobyte
	myMegabyte
	myGigabyte
	myTerabyte
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
		return int64(bytes * myTerabyte), nil
	case "G", "GB":
		return int64(bytes * myGigabyte), nil
	case "M", "MB":
		return int64(bytes * myMegabyte), nil
	case "K", "KB":
		return int64(bytes * myKilobyte), nil
	case "B":
		return int64(bytes), nil
	default:
		return 0, errors.New("invalid size")
	}
}

func main() {
	num := flag.Uint64("n", 1, "number of files")
	size := flag.String("s", "1M", "size(K,M,G,T) of file")
	out := flag.String("o", ".", "output dir")
	prefix := flag.String("p", "vager", "filename prefix")
	threads := flag.Int("t", runtime.NumCPU(), "threads")
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
	filePrefix := filepath.Join(*out, fmt.Sprintf("%s-%s", *prefix, *size))
	fmt.Printf("gen %d file with size[%s] to %s\n", *num, *size, *out)
	var index uint64
	wg := sync.WaitGroup{}
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func(fileprefix string, n uint64) {
			ifd := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				fn := atomic.AddUint64(&index, 1)
				if fn > n {
					break
				}
				filename := fmt.Sprintf("%s-%d", fileprefix, fn)
				ofd, err := os.Create(filename)
				if err != nil {
					fmt.Printf("create file %s failed %s\n", filename, err)
					break
				}
				if _, err = io.CopyN(ofd, ifd, fileSize); err != nil {
					fmt.Printf("gen random file %s failed %s\n", filename, err)
					ofd.Close()
					break
				}
				ofd.Close()
			}
			wg.Done()
		}(filePrefix, *num)
	}
	wg.Wait()
}
