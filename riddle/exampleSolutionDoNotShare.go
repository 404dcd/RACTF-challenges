package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var primesUnder125 = []byte{37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
var mul5under125 = []byte{35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120, 125}

var std = []byte{33, 34, 36, 38, 39, 42, 44, 46, 48, 49, 51, 52, 54, 56, 57, 58, 62, 63, 64, 66, 68, 69, 72, 74, 76, 77, 78, 81, 82, 84, 86, 87, 88, 91, 92, 93, 94, 96, 98, 99, 102, 104, 106, 108, 111, 112, 114, 116, 117, 118, 119, 121, 122, 123, 124}

var charset [26][]byte

func main() {
	charset[1-1] = std
	charset[2-1] = primesUnder125
	charset[3-1] = std
	charset[4-1] = std
	charset[5-1] = std
	charset[6-1] = std
	charset[7-1] = []byte{51}
	charset[8-1] = []byte{110}
	charset[9-1] = []byte{71}
	charset[10-1] = []byte{108}
	charset[11-1] = []byte{49}
	charset[12-1] = []byte{115}
	charset[13-1] = []byte{104}
	charset[14-1] = []byte{95}
	charset[15-1] = []byte{99}
	charset[16-1] = []byte{52}
	charset[17-1] = []byte{110}
	charset[18-1] = []byte{95}
	charset[19-1] = []byte{11, 22, 33, 44, 55, 66, 77, 88, 99, 111}
	charset[20-1] = primesUnder125
	charset[21-1] = mul5under125
	charset[22-1] = []byte{34, 36, 38, 42, 44, 46, 48, 52, 54, 56, 58, 62, 64, 66, 68, 72, 74, 76, 78, 82, 84, 86, 88, 92, 94, 96, 98, 102, 104, 106, 108, 112, 114, 116, 118, 122, 124}
	charset[23-1] = []byte{34, 36, 38, 42, 44, 46, 48, 52, 54, 56, 58, 62, 64, 66, 68, 72, 74, 76, 78, 82, 84, 86, 88, 92, 94, 96, 98, 102, 104, 106, 108, 112, 114, 116, 118, 122, 124}
	charset[24-1] = []byte{34, 36, 38, 42, 44, 46, 48, 52, 54, 56, 58, 62, 64, 66, 68, 72, 74, 76, 78, 82, 84, 86, 88, 92, 94, 96, 98, 102, 104, 106, 108, 112, 114, 116, 118, 122, 124}
	charset[25-1] = []byte{100}
	charset[26-1] = []byte{125}

	var flagind [26]int
	var flag [26]byte

	for i := 0; charset[0][i] != 'r'; i++ {
		flagind[0] = i + 1
	}
	for i := 0; charset[1][i] != 'a'; i++ {
		flagind[1] = i + 1
	}
	for i := 0; charset[2][i] != 'c'; i++ {
		flagind[2] = i + 1
	}
	for i := 0; charset[3][i] != 't'; i++ {
		flagind[3] = i + 1
	}
	for i := 0; charset[4][i] != 'f'; i++ {
		flagind[4] = i + 1
	}
	for i := 0; charset[5][i] != '{'; i++ {
		flagind[5] = i + 1
	}

	inc := 1

	for {
		i := len(flagind) - 1
		flagind[i] += inc
		for flagind[i] >= len(charset[i]) {
			flagind[i] = 0
			i--
			flagind[i] += inc
		}

		for i := 0; i < len(flagind); i++ {
			flag[i] = charset[i][flagind[i]]
		}

		total := 0
		for _, val := range flag {
			total += int(val)
			if total >= 2435 {
				break
			}
		}

		if total == 2435 {
			ff := ""
			f := ""
			for _, val := range flag {
				f += string(val)
				ff += strconv.Itoa(int(val))
			}
			if strings.Count(f, "_") > 2 && len(ff) == 65 {
				h := sha256.New()
				var stuff []byte
				for _, val := range flag {
					stuff = append(stuff, val)
				}
				h.Write(stuff)
				if fmt.Sprintf("%x", h.Sum(nil))[0:2] == "90" {
					fmt.Println(f)
				}
			}
		}

	}
}
