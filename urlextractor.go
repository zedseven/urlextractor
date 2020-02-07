package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const chunkSize = 50000

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please call this program with a single argument: the path to a file.")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	r := bufio.NewReader(f)
	b := make([]byte, chunkSize)
	t := make([]byte, chunkSize * 2)

	urlDetector := regexp.MustCompile(
		"(?:http|ftp|https)://[\\w_-]+(?:(?:\\.[\\w_-]+)+)(?:[\\w.,@?^=%&:/~+#-]*[\\w@?^=%&/~+#-])?")
	urls := make([]string, 0)

	for {
		n, err := r.Read(b)
		if n > 0 {
			t = append(t[chunkSize:], b...)
			//fmt.Println(string(t))
			urls = append(urls, urlDetector.FindAllString(string(t), -1)...)
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
				return
			}
			break
		}
	}

	urls = sliceUniqMap(urls)
	for _, v := range urls {
		fmt.Println(v)
	}
}

//https://www.reddit.com/r/golang/comments/5ia523/idiomatic_way_to_remove_duplicates_in_a_slice/db6qa2e/
func sliceUniqMap(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}