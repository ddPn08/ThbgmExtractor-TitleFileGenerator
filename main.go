package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var dir string

func readFile(pathname string) []byte {
	data, err := ioutil.ReadFile(filepath.Join(dir, pathname))
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func readShiftJis(b []byte) string {
	reader := transform.NewReader(bytes.NewBuffer(b), japanese.ShiftJIS.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}

func findName(musiccmt string, filename string) string {
	lines := strings.Split(musiccmt, "\r\n")
	found := false
	for _, line := range lines {
		if found {
			matches := regexp.MustCompile(`No.( ?\d+) +(.+)`).FindSubmatch([]byte(line))
			return string(matches[2])
		}
		found = strings.HasPrefix(fmt.Sprintf("@bgm/%s", filename), line) && len(line) > 0
	}
	return filename
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <dir>")
		os.Exit(1)
	}
	dir = os.Args[1]

	bgmfmt := readFile("thbgm.fmt")
	musiccmt := readFile("musiccmt.txt")

	musiccmtTxt := readShiftJis(musiccmt)

	result := ""

	for i := 0; i < (len(bgmfmt) / 52); i += 1 {
		entry := bgmfmt[i*52 : (i+1)*52]

		name := string(entry[0:15])

		startOffset := reverseBytes(entry[16:20])
		loopStartOffset := reverseBytes(entry[24:28])
		withLoopLength := reverseBytes(entry[28:32])

		loopLength := intToByteHex(byteHexToInt(withLoopLength) - byteHexToInt(loopStartOffset))

		result += fmt.Sprintf("%s,", byteTo8ByteHex(startOffset))
		result += fmt.Sprintf("%s,", byteTo8ByteHex(loopStartOffset))
		result += fmt.Sprintf("%s,", byteTo8ByteHex(loopLength))
		result += fmt.Sprintf("%s\n", findName(musiccmtTxt, name))
	}

	fmt.Print(result)
}
